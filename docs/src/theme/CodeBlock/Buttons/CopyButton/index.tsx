import React, {useCallback, type ReactNode} from 'react';
import CopyButton from '@theme-original/CodeBlock/Buttons/CopyButton';
import {useCodeBlockContext} from '@docusaurus/theme-common/internal';
import type {Props} from '@theme/CodeBlock/Buttons/CopyButton';

declare global {
  interface Window {
    posthog?: {
      capture: (event: string, properties?: Record<string, unknown>) => void;
    };
  }
}

export default function CopyButtonWrapper(props: Props): ReactNode {
  const {metadata} = useCodeBlockContext();

  const handleClick = useCallback(() => {
    window.posthog?.capture('code_copied', {
      helper: window.location.hash.replace('#', '') || null,
      page: window.location.pathname,
      code_preview: metadata.code?.slice(0, 120),
    });
  }, [metadata.code]);

  return (
    <span onClick={handleClick}>
      <CopyButton {...props} />
    </span>
  );
}
