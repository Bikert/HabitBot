import { create } from 'zustand/index'
import { TelegramWebApp } from '../telegram'

export const useTelegramTheme = create<{ theme: 'dark' | 'light' }>()(() => ({
  theme: TelegramWebApp.colorScheme === 'light' ? 'light' : 'dark',
}))

TelegramWebApp.onEvent('themeChanged', () => {
  useTelegramTheme.setState({ theme: TelegramWebApp.colorScheme === 'light' ? 'light' : 'dark' })
})
