import { composeRenderProps, Button as RACButton, ButtonProps as RACButtonProps } from 'react-aria-components'
import { tv } from 'tailwind-variants'
import { focusRing } from './utils'

export interface ButtonVariants {
  variant?: 'primary' | 'secondary' | 'destructive' | 'icon'
  size?: 'sm' | 'md' | 'xs' | 'lg'
}

export type ButtonProps = RACButtonProps & ButtonVariants

export const buttonStyles = tv({
  extend: focusRing,
  // TODO
  base: 'cursor-default rounded-full border border-black/10 text-center font-medium shadow-[inset_0_1px_0_0_rgba(255,255,255,0.1)] transition dark:border-white/10 dark:shadow-none',
  variants: {
    variant: {
      primary: 'bg-primary text-on-primary hover:bg-primary-hover pressed:bg-primary-press',
      secondary: 'bg-secondary text-on-secondary hover:bg-secondary-hover pressed:bg-secondary-press',
      destructive: 'bg-error text-on-error hover:bg-error-hover pressed:bg-error-press',
      // TODO
      icon: 'flex items-center justify-center border-0 p-1 text-gray-600 hover:bg-black/[5%] disabled:bg-transparent dark:text-zinc-400 dark:hover:bg-white/10 pressed:bg-black/10 dark:pressed:bg-white/20',
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
    variant: 'primary',
    size: 'sm',
  },
})

export function Button(props: ButtonProps) {
  return (
    <RACButton
      {...props}
      className={composeRenderProps(props.className, (className, renderProps) =>
        buttonStyles({ ...renderProps, variant: props.variant, size: props.size, className }),
      )}
    />
  )
}
