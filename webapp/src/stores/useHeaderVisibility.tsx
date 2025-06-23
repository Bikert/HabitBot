import { create } from 'zustand'
import { immer } from 'zustand/middleware/immer'
import { createJSONStorage, persist } from 'zustand/middleware'
import { DeviceStorage } from '../telegram/telegramStorage'

export interface HeaderVisibilityState {
  visible: boolean
  show: () => void
  hide: () => void
  toggle: () => void
}

export const useHeaderVisibility = create<HeaderVisibilityState>()(
  persist(
    immer((set) => ({
      visible: false,
      show: () =>
        set((state) => {
          state.visible = true
        }),
      hide: () =>
        set((state) => {
          state.visible = true
        }),
      toggle: () =>
        set((state) => {
          state.visible = !state.visible
        }),
    })),
    {
      name: 'debug-storage',
      storage: createJSONStorage(() => DeviceStorage),
    },
  ),
)
