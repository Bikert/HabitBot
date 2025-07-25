import { FC, useEffect } from 'react'
import { TelegramWebApp } from '../index'
import { create } from 'zustand/index'

interface BackButtonProps {
  onClick?: VoidFunction
}

const backButton = TelegramWebApp.BackButton

type GlobalButtonTrackerState = {
  openedButtonsCount: number
  increment: () => void
  decrement: () => void
}

const backButtonTrackedStore = create<GlobalButtonTrackerState>()((set) => ({
  openedButtonsCount: 0,
  increment: () => {
    set((state) => ({ openedButtonsCount: state.openedButtonsCount + 1 }))
  },
  decrement: () => {
    set((state) => ({ openedButtonsCount: state.openedButtonsCount - 1 }))
  },
}))

export const BackButton: FC<BackButtonProps> = ({ onClick = () => {} }) => {
  useEffect(() => {
    backButtonTrackedStore.getState().increment()
    backButton.show()
    return () => {
      backButtonTrackedStore.getState().decrement()
      if (backButtonTrackedStore.getState().openedButtonsCount === 0) {
        backButton.hide()
      }
    }
  }, [])

  useEffect(() => {
    TelegramWebApp.onEvent('backButtonClicked', onClick)
    return () => {
      TelegramWebApp.offEvent('backButtonClicked', onClick)
    }
  }, [onClick])

  return null
}
