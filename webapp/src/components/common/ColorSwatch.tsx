import { ColorSwatch as AriaColorSwatch, ColorSwatchProps } from 'react-aria-components'
import { composeTailwindRenderProps } from './utils'

export function ColorSwatch(props: ColorSwatchProps) {
  return (
    <AriaColorSwatch
      {...props}
      className={composeTailwindRenderProps(props.className, 'h-6 w-6 rounded-full border border-outline/10')}
      style={({ color }) => ({
        background: `linear-gradient(${color}, ${color}),
          repeating-conic-gradient(#CCC 0% 25%, white 0% 50%) 50% / 16px 16px`,
      })}
    />
  )
}
