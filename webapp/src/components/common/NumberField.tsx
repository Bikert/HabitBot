import { NumberField as AriaNumberField, NumberFieldProps as AriaNumberFieldProps } from 'react-aria-components'
import { composeTailwindRenderProps } from './utils'
import { Description, FieldError, Input, Label } from './Field'
import { type AdditionalFieldProps, textInputStyles } from './TextField'

export type NumberFieldProps = AriaNumberFieldProps & AdditionalFieldProps

export function NumberField({ label, description, errorMessage, ...props }: NumberFieldProps) {
  return (
    <AriaNumberField {...props} className={composeTailwindRenderProps(props.className, 'flex flex-col gap-1')}>
      {({ isInvalid }) => (
        <>
          {label && <Label>{label}</Label>}
          <Input className={textInputStyles} />
          {!isInvalid && description && <Description>{description}</Description>}
          <FieldError>{errorMessage}</FieldError>
        </>
      )}
    </AriaNumberField>
  )
}
