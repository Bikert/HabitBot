import {
  ColorSwatchPicker as AriaColorSwatchPicker,
  ColorSwatchPickerItem as AriaColorSwatchPickerItem,
  ColorSwatchPickerItemProps,
  ColorSwatchPickerProps,
} from 'react-aria-components'
import { ColorSwatch } from './ColorSwatch'
import { composeTailwindRenderProps, focusRing } from './utils'
import { tv } from 'tailwind-variants'

export function ColorSwatchPicker({ children, ...props }: Omit<ColorSwatchPickerProps, 'layout'>) {
  return (
    <AriaColorSwatchPicker {...props} className={composeTailwindRenderProps(props.className, 'flex gap-1')}>
      {children}
    </AriaColorSwatchPicker>
  )
}

const itemStyles = tv({
  extend: focusRing,
  base: 'relative rounded-full border-3 border-outline/20',
  variants: {
    isSelected: {
      true: 'border-primary',
    },
  },
})

export function ColorSwatchPickerItem(props: ColorSwatchPickerItemProps) {
  return (
    <AriaColorSwatchPickerItem {...props} className={itemStyles}>
      <ColorSwatch />
    </AriaColorSwatchPickerItem>
  )
}
