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
function trackSearch(): void {
  let inputEl: HTMLInputElement | null = null;

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
  };

  const observer = new MutationObserver(() => {
    const input = document.querySelector<HTMLInputElement>('.DocSearch-Input');
    if (input) attachInputListener(input);
  });

  observer.observe(document.body, {childList: true, subtree: true});
}

export function onRouteDidUpdate(): void {
  // Re-run on each navigation in case DOM changed
}

// Runs once on initial load
trackSponsorClicks();
trackSearch();
