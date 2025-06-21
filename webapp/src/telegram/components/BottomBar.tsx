import { ReactNode, useEffect } from 'react'
import { TelegramWebApp } from '../index'

const defaultBottomBarColor =
  TelegramWebApp.themeParams.bottom_bar_bg_color ?? TelegramWebApp.themeParams.secondary_bg_color

export const BottomBar = ({
  bgColor = defaultBottomBarColor,
  children = null,
}: {
  bgColor?: string
  children?: ReactNode
}) => {
  useEffect(() => {
    TelegramWebApp.setBottomBarColor(bgColor)
  }, [bgColor])

  useEffect(() => {
    return () => {
      TelegramWebApp.setBottomBarColor(defaultBottomBarColor)
    }
  }, [])

  return <>{children}</>
}
