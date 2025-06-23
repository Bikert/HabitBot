import { FC, useEffect } from 'react'
import { BottomButtonProps, useBottomButton } from './bottomButton'
import { SecondaryButton as SecondaryButtonType } from '../types'
import { TelegramWebApp } from '../index'

const secondaryButton = TelegramWebApp.SecondaryButton

export const SecondaryButton: FC<BottomButtonProps & { position?: SecondaryButtonType['position'] }> = ({
  disabled,
  color,
  textColor,
  text,
  onClick,
  progress,
  hasShineEffect,
  position = 'bottom',
}) => {
  useBottomButton({
    type: 'secondary',
    disabled,
    progress,
    textColor,
    text,
    onClick,
    color,
    hasShineEffect,
  })

  useEffect(() => {
    secondaryButton.setParams({
      position,
    })
  }, [position])

  return null
}
