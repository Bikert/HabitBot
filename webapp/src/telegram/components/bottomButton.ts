import { useEffect } from 'react'
import { TelegramWebApp } from '../index'

export type BottomButtonProps = {
  disabled?: boolean
  progress?: boolean
  color?: string
  textColor?: string
  onClick?: VoidFunction
  text: string
  hasShineEffect?: boolean
}

const { bottom_bar_bg_color, button_color, button_text_color } = TelegramWebApp.themeParams

type ButtonTypes = 'main' | 'secondary'

const defaultButtonColors: Record<ButtonTypes, Parameters<typeof TelegramWebApp.MainButton.setParams>[0]> = {
  main: {
    color: button_color,
    text_color: button_text_color,
  },
  secondary: {
    color: bottom_bar_bg_color,
    text_color: button_color,
  },
}

const isButtonShown: Record<ButtonTypes, boolean> = {
  main: false,
  secondary: false,
}

export const useBottomButton = ({
  type,
  progress = false,
  disabled = false,
  color,
  textColor,
  text,
  onClick,
  hasShineEffect = false,
}: {
  type: ButtonTypes
} & BottomButtonProps) => {
  const button = type === 'main' ? TelegramWebApp.MainButton : TelegramWebApp.SecondaryButton

  useEffect(() => {
    button.show()
    isButtonShown[type] = true
    return () => {
      isButtonShown[type] = false
      setTimeout(() => {
        if (!isButtonShown[type]) {
          button.hide()
        }
      }, 10)
    }
  }, [button, type])

  useEffect(() => {
    if (progress) {
      button.showProgress()
      button.disable()
    } else {
      button.hideProgress()
    }

    if (disabled || progress) {
      button.disable()
    } else {
      button.enable()
    }

    return () => {
      button.hideProgress()
      button.enable()
    }
  }, [button, disabled, progress])

  useEffect(() => {
    button.setParams({
      color: color ?? defaultButtonColors[type].color,
      text_color: textColor ?? defaultButtonColors[type].text_color,
      has_shine_effect: hasShineEffect,
    })
  }, [color, textColor, hasShineEffect, button, type])

  useEffect(() => {
    button.setText(text)
  }, [button, text])

  useEffect(() => {
    if (onClick) {
      button.onClick(onClick)
      return () => {
        button.offClick(onClick)
      }
    }
  }, [button, onClick])
}
