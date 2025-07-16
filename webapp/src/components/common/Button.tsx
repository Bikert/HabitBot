import { composeRenderProps, Button as RACButton, ButtonProps as RACButtonProps } from 'react-aria-components'
import { tv } from 'tailwind-variants'
import { focusRing } from './utils'

export interface ButtonVariants {
  variant?: 'filled' | 'elevated' | 'tonal' | 'outlined' | 'text'
  color?: 'primary' | 'secondary' | 'tertiary' | 'destructive'
  size?: 'sm' | 'md' | 'xs' | 'lg'
}

export type ButtonProps = RACButtonProps & ButtonVariants

export const buttonStyles = tv({
  extend: focusRing,
  base: 'relative cursor-default rounded-full text-center font-medium transition',
  variants: {
    variant: {
      filled: 'bg-primary text-on-primary hover:bg-primary-hover pressed:bg-primary-press',
      elevated: 'bg-surface-container-low text-primary shadow shadow-shadow',
      tonal:
        'bg-secondary-container text-on-secondary-container hover:bg-secondary-container-hover disabled:text-on-surface pressed:bg-secondary-container-press',
      outlined:
        'border border-outline-variant bg-outline-variant text-on-surface-variant selected:bg-inverse-surface selected:text-inverse-on-surface',
      text: 'text-primary hover:bg-primary/10 disabled:bg-on-background/10 disabled:text-on-surface/40 pressed:bg-primary/10',
    },
    color: {
      primary: '',
      secondary: 'override-color-secondary',
      tertiary: 'override-color-tertiary',
      destructive: 'override-color-destructive',
    },
    size: {
      xs: 'px-2 py-1 text-xs',
      sm: 'px-4 py-2 text-sm',
      md: 'text-md px-6 py-2',
      lg: 'px-12 py-3 text-lg',
    },
    isDisabled: {
      // TODO
      true: 'border-black/5 bg-gray-100 text-gray-300 dark:border-white/5 dark:bg-zinc-800 dark:text-zinc-600 forced-colors:text-[GrayText]',
    },
    isPressed: {
      true: 'rounded-xl',
    },
  },
  defaultVariants: {
    variant: 'filled',
    color: 'primary',
    size: 'sm',
  },
})

export function Button(props: ButtonProps) {
  return (
    <RACButton
      {...props}
      className={composeRenderProps(props.className, (className, renderProps) =>
        buttonStyles({ ...renderProps, variant: props.variant, size: props.size, color: props.color, className }),
      )}
    />
  )
}
