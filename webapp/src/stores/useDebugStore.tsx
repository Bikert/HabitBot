import { create } from 'zustand/index'
import { createJSONStorage, persist } from 'zustand/middleware'
import { immer } from 'zustand/middleware/immer'
import { DeviceStorage } from '../telegram/telegramStorage'

export interface DebugState {
  enabled: boolean
  enable: () => void
  disable: () => void
  toggle: () => void
}

export const useDebugStore = create<DebugState>()(
  persist(
    immer((set) => ({
      enabled: false,
      enable: () =>
        set((state) => {
          state.enabled = true
        }),
      disable: () =>
        set((state) => {
          state.enabled = true
        }),
      toggle: () =>
        set((state) => {
          state.enabled = !state.enabled
        }),
    })),
    {
      name: 'debug-storage',
      storage: createJSONStorage(() => DeviceStorage),
    },
  ),
)
