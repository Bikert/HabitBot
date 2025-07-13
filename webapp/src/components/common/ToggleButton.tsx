import {
  composeRenderProps,
  ToggleButton as RACToggleButton,
  ToggleButtonProps as RACToggleButtonProps,
} from 'react-aria-components'
import { tv } from 'tailwind-variants'
import { buttonStyles, type ButtonVariants } from './Button'

export type ToggleButtonProps = RACToggleButtonProps & ButtonVariants

const styles = tv({
  extend: buttonStyles,
  variants: {
    isSelected: {
      false: '',
      true: 'rounded-xl',
    },
  },
  compoundVariants: [
    {
      variant: 'primary',
      isSelected: false,
      class:
        'focus-within::bg-surface-container-focus bg-surface-container text-on-surface-variant hover:bg-surface-container-hover pressed:bg-surface-container-press',
    },
  ],
})

export function ToggleButton(props: ToggleButtonProps) {
  return (
    <RACToggleButton
      {...props}
      className={composeRenderProps(props.className, (className, renderProps) =>
        styles({ ...renderProps, className, variant: props.variant, size: props.size }),
      )}
    />
  )
}
