import {
  TextField as AriaTextField,
  TextFieldProps as AriaTextFieldProps,
  ValidationResult,
} from 'react-aria-components'
import { tv } from 'tailwind-variants'
import { Description, fieldBorderStyles, FieldError, Input, Label } from './Field'
import { composeTailwindRenderProps } from './utils'

export const textInputStyles = tv({
  base: 'rounded-md border-2',
  variants: {
    isFocused: fieldBorderStyles.variants.isFocusWithin,
    isInvalid: fieldBorderStyles.variants.isInvalid,
    isDisabled: fieldBorderStyles.variants.isDisabled,
  },
})

export interface AdditionalFieldProps {
  label?: string
  description?: string
  errorMessage?: string | ((validation: ValidationResult) => string)
}

export type TextFieldProps = AriaTextFieldProps & AdditionalFieldProps

export function TextField({ label, description, errorMessage, ...props }: TextFieldProps) {
  return (
    <AriaTextField {...props} className={composeTailwindRenderProps(props.className, 'flex flex-col gap-1')}>
      {({ isInvalid }) => (
        <>
          {label && <Label>{label}</Label>}
          <Input className={textInputStyles} />
          {!isInvalid && description && <Description>{description}</Description>}
          <FieldError>{errorMessage}</FieldError>
        </>
      )}
    </AriaTextField>
  )
}
