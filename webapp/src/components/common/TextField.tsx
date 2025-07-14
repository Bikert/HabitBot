import {
  TextField as AriaTextField,
  NumberField as AriaNumberField,
  TextFieldProps as AriaTextFieldProps,
  NumberFieldProps as AriaNumberFieldProps,
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

interface AdditionalFieldProps {
  label?: string
  description?: string
  errorMessage?: string | ((validation: ValidationResult) => string)
}

export type TextFieldProps = AriaTextFieldProps & AdditionalFieldProps

export type NumberFieldProps = AriaNumberFieldProps & AdditionalFieldProps

export function TextField({ label, description, errorMessage, ...props }: TextFieldProps) {
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

export function NumberField({ label, description, errorMessage, ...props }: NumberFieldProps) {
  return (
    <AriaNumberField {...props} className={composeTailwindRenderProps(props.className, 'flex flex-col gap-1')}>
      {({ isInvalid }) => (
        <>
          {label && <Label>{label}</Label>}
          <Input className={inputStyles} />
          {!isInvalid && description && <Description>{description}</Description>}
          <FieldError>{errorMessage}</FieldError>
        </>
      )}
    </AriaNumberField>
  )
}
