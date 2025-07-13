import {
  FieldErrorProps,
  Group,
  GroupProps,
  InputProps,
  LabelProps,
  FieldError as RACFieldError,
  Input as RACInput,
  Label as RACLabel,
  Text,
  TextProps,
  composeRenderProps,
} from 'react-aria-components'
import { twMerge } from 'tailwind-merge'
import { tv } from 'tailwind-variants'
import { composeTailwindRenderProps, focusRing } from './utils'

export function Label(props: LabelProps) {
  return (
    <RACLabel
      {...props}
      className={twMerge('w-fit cursor-default text-sm font-medium text-primary', props.className)}
    />
  )
}

export function Description(props: TextProps) {
  return <Text {...props} slot="description" className={twMerge('text-sm text-on-surface-variant', props.className)} />
}

export function FieldError(props: FieldErrorProps) {
  return (
    <RACFieldError
      {...props}
      className={composeTailwindRenderProps(props.className, 'text-sm text-error forced-colors:text-[Mark]')}
    />
  )
}

export const fieldBorderStyles = tv({
  variants: {
    isFocusWithin: {
      false: 'border-outline forced-colors:border-[ButtonBorder]',
      true: 'focus:border-primary focus:forced-colors:border-[Highlight]',
    },
    isInvalid: {
      true: 'border-error forced-colors:border-[Mark]',
    },
    isDisabled: {
      true: 'border-on-surface forced-colors:border-[GrayText]',
    },
  },
})

export const fieldGroupStyles = tv({
  extend: focusRing,
  base: 'group flex h-9 items-center overflow-hidden rounded-lg border-2 bg-white dark:bg-zinc-900 forced-colors:bg-[Field]',
  variants: fieldBorderStyles.variants,
})

export function FieldGroup(props: GroupProps) {
  return (
    <Group
      {...props}
      className={composeRenderProps(props.className, (className, renderProps) =>
        fieldGroupStyles({ ...renderProps, className }),
      )}
    />
  )
}

export function Input(props: InputProps) {
  return (
    <RACInput
      {...props}
      className={composeTailwindRenderProps(
        props.className,
        'min-w-0 flex-1 px-3 py-2.5 text-sm text-on-surface outline-0 disabled:text-on-surface/38',
      )}
    />
  )
}
