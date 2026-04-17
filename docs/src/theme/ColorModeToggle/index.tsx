import React, {type ReactNode} from 'react';
import OriginalColorModeToggle from '@theme-original/ColorModeToggle';
import type ColorModeToggleType from '@theme/ColorModeToggle';
import type {WrapperProps} from '@docusaurus/types';

type Props = WrapperProps<typeof ColorModeToggleType>;

export default function ColorModeToggleWrapper(props: Props): ReactNode {
  const handleChange: Props['onChange'] = (newMode) => {
    window.posthog?.capture('color_mode_toggled', {
      from: props.value ?? 'system',
      to: newMode ?? 'system',
    });
    props.onChange(newMode);
  };

  return <OriginalColorModeToggle {...props} onChange={handleChange} />;
}
