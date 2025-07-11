import {
  TextField as AriaTextField,
  TextFieldProps as AriaTextFieldProps,
  ValidationResult,
} from 'react-aria-components'
import { tv } from 'tailwind-variants'
import { Description, FieldError, Input, Label, fieldBorderStyles } from './Field'
import { composeTailwindRenderProps } from './utils'

const inputStyles = tv({
  base: 'rounded-md border-2',
  variants: {
    isFocused: fieldBorderStyles.variants.isFocusWithin,
    isInvalid: fieldBorderStyles.variants.isInvalid,
    isDisabled: fieldBorderStyles.variants.isDisabled,
  },
})

export interface TextFieldProps extends AriaTextFieldProps {
  label?: string
  description?: string
  errorMessage?: string | ((validation: ValidationResult) => string)
}

export function TextField({ label, description, errorMessage, ...props }: TextFieldProps) {
  console.log(errorMessage, 'error message')
  return (
    <AriaTextField {...props} className={composeTailwindRenderProps(props.className, 'flex flex-col gap-1')}>
      {({ isInvalid }) => (
        <>
          {label && <Label>{label}</Label>}
          <Input className={inputStyles} />
          {!isInvalid && description && <Description>{description}</Description>}
          <FieldError>{errorMessage}</FieldError>
        </>
      )}
    </AriaTextField>
  )
}
