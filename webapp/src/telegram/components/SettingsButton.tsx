import { FC, useEffect } from 'react'
import { TelegramWebApp } from '../index'
import { create } from 'zustand/index'

interface SettingsButtonProps {
  onClick?: VoidFunction
}

const settingsButton = TelegramWebApp.SettingsButton

type GlobalButtonTrackerState = {
  openedButtonsCount: number
  increment: () => void
  decrement: () => void
}

const settingsButtonTrackedStore = create<GlobalButtonTrackerState>()((set) => ({
  openedButtonsCount: 0,
  increment: () => {
    set((state) => ({ openedButtonsCount: state.openedButtonsCount + 1 }))
  },
  decrement: () => {
    set((state) => ({ openedButtonsCount: state.openedButtonsCount - 1 }))
  },
}))

export const SettingsButton: FC<SettingsButtonProps> = ({ onClick = () => {} }) => {
  useEffect(() => {
    settingsButtonTrackedStore.getState().increment()
    settingsButton.show()
    return () => {
      settingsButtonTrackedStore.getState().decrement()
      if (settingsButtonTrackedStore.getState().openedButtonsCount === 0) {
        settingsButton.hide()
      }
    }
  }, [])

  useEffect(() => {
    TelegramWebApp.onEvent('settingsButtonClicked', onClick)
    return () => {
      TelegramWebApp.offEvent('settingsButtonClicked', onClick)
    }
  }, [onClick])

  return null
}
