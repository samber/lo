declare global {
  interface Window {
    posthog?: {
      capture: (event: string, properties?: Record<string, unknown>) => void;
    };
  }
}

// --- Sponsor click tracking (navbar + sidebar) ---
function trackSponsorClicks(): void {
  document.addEventListener('click', (e) => {
    const anchor = (e.target as Element).closest('a[href*="sponsors/samber"]');
    if (!anchor) return;

    const isNavbar = anchor.closest('.navbar') !== null;
    window.posthog?.capture('sponsor_clicked', {
      location: isNavbar ? 'navbar' : 'sidebar',
      href: (anchor as HTMLAnchorElement).href,
    });
  });
}

// --- Search query tracking ---
// Captures both explicit submissions (Enter) and, crucially, zero-result
// queries -- Algolia DocSearch is an as-you-type search, so most users never
// press Enter. A query that returns zero results is the cheapest, most
// direct signal of a content gap: it's a developer's own words for
// something the site doesn't (yet) cover.
function trackSearch(): void {
  let inputEl: HTMLInputElement | null = null;
  let debounceTimer: ReturnType<typeof setTimeout> | undefined;
  let lastReportedQuery = '';

  const reportResultsState = (query: string) => {
    if (!query || query === lastReportedQuery) return;
    lastReportedQuery = query;

    const noResults = document.querySelector('.DocSearch-NoResults') !== null;
    if (noResults) {
      window.posthog?.capture('search_no_results', {query});
      return;
    }
    const hitCount = document.querySelectorAll('.DocSearch-Hit').length;
    if (hitCount > 0) {
      window.posthog?.capture('search_results', {query, count: hitCount});
    }
  };

  const attachInputListener = (input: HTMLInputElement) => {
    if (input === inputEl) return;
    inputEl = input;

    input.addEventListener('keydown', (e) => {
      if (e.key === 'Enter' && input.value.trim()) {
        window.posthog?.capture('search_submitted', {
          query: input.value.trim(),
        });
      }
    });

    input.addEventListener('input', () => {
      const query = input.value.trim();
      if (debounceTimer) clearTimeout(debounceTimer);
      if (!query) return;
      // Wait for DocSearch's own results DOM to settle before reading it.
      debounceTimer = setTimeout(() => reportResultsState(query), 400);
    });
  };

  const observer = new MutationObserver(() => {
    const input = document.querySelector<HTMLInputElement>('.DocSearch-Input');
    if (input) attachInputListener(input);
  });

  observer.observe(document.body, {childList: true, subtree: true});
}

// --- Helper card visibility tracking ---
// Answers "which of the 449 helpers do people actually look at?" -- a
// ranking that exists nowhere else (not in GitHub stars, not in pkg.go.dev
// importers). Fires once per helper per page view, after it has been at
// least half visible for 2 uninterrupted seconds (skims don't count).
function trackHelperVisibility(): void {
  const seen = new Set<string>();
  const dwellTimers = new Map<Element, ReturnType<typeof setTimeout>>();

  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        const el = entry.target as HTMLElement;
        const slug = el.dataset.helperSlug;
        if (!slug) return;

        if (entry.isIntersecting) {
          if (seen.has(slug) || dwellTimers.has(el)) return;
          const timer = setTimeout(() => {
            dwellTimers.delete(el);
            if (seen.has(slug)) return;
            seen.add(slug);
            const cards = Array.from(document.querySelectorAll<HTMLElement>('[data-helper-slug]'));
            window.posthog?.capture('helper_viewed', {
              helper: slug,
              category: el.dataset.helperCategory,
              subCategory: el.dataset.helperSubcategory,
              position_in_page: cards.indexOf(el),
            });
          }, 2000);
          dwellTimers.set(el, timer);
        } else {
          const timer = dwellTimers.get(el);
          if (timer) {
            clearTimeout(timer);
            dwellTimers.delete(el);
          }
        }
      });
    },
    {threshold: 0.5}
  );

  const attachToVisibleCards = () => {
    document.querySelectorAll<HTMLElement>('[data-helper-slug]').forEach((el) => {
      if (el.dataset.helperObserved) return;
      el.dataset.helperObserved = 'true';
      observer.observe(el);
    });
  };

  attachToVisibleCards();
  const domObserver = new MutationObserver(attachToVisibleCards);
  domObserver.observe(document.body, {childList: true, subtree: true});
}

export function onRouteDidUpdate(): void {
  // Re-run on each navigation in case DOM changed
}

// Runs once on initial load (browser only)
if (typeof document !== 'undefined') {
  trackSponsorClicks();
  trackSearch();
  trackHelperVisibility();
}
