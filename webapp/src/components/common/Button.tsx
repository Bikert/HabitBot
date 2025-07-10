import { composeRenderProps, Button as RACButton, ButtonProps as RACButtonProps } from 'react-aria-components'
import { tv } from 'tailwind-variants'
import { focusRing } from './utils'

export interface ButtonProps extends RACButtonProps {
  variant?: 'primary' | 'secondary' | 'destructive' | 'icon'
}

const button = tv({
  extend: focusRing,
  base: 'cursor-default rounded-lg border border-black/10 px-5 py-2 text-center text-sm shadow-[inset_0_1px_0_0_rgba(255,255,255,0.1)] transition dark:border-white/10 dark:shadow-none',
  variants: {
    variant: {
      primary: 'bg-primary text-on-primary hover:bg-primary-hover pressed:bg-primary-press',
      secondary: 'bg-secondary text-on-secondary hover:bg-secondary-hover pressed:bg-secondary-press',
      destructive: 'bg-error text-on-error hover:bg-error-hover pressed:bg-error-press',
      icon: 'flex items-center justify-center border-0 p-1 text-gray-600 hover:bg-black/[5%] disabled:bg-transparent dark:text-zinc-400 dark:hover:bg-white/10 pressed:bg-black/10 dark:pressed:bg-white/20',
    },
    isDisabled: {
      true: 'border-black/5 bg-gray-100 text-gray-300 dark:border-white/5 dark:bg-zinc-800 dark:text-zinc-600 forced-colors:text-[GrayText]',
    },
  },
  defaultVariants: {
    variant: 'primary',
  },
})

export function Button(props: ButtonProps) {
  return (
    <RACButton
      {...props}
      className={composeRenderProps(props.className, (className, renderProps) =>
        button({ ...renderProps, variant: props.variant, className }),
      )}
    />
  )
}
