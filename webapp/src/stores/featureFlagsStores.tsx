import { create } from 'zustand/index'
import { createJSONStorage, persist } from 'zustand/middleware'
import { immer } from 'zustand/middleware/immer'
import { DeviceStorage } from '../telegram/telegramStorage'

export interface FeatureFlagState {
  active: boolean
  enable: () => void
  disable: () => void
  toggle: () => void
}

function createFeatureFlagStore(storeName: string, defaultValue = false) {
  return create<FeatureFlagState>()(
    persist(
      immer((set) => ({
        active: defaultValue,
        enable: () =>
          set((state) => {
            state.active = true
          }),
        disable: () =>
          set((state) => {
            state.active = false
          }),
        toggle: () =>
          set((state) => {
            state.active = !state.active
          }),
      })),
      {
        name: storeName,
        storage: createJSONStorage(() => DeviceStorage),
      },
    ),
  )
}

export const useShowDebugInformation = createFeatureFlagStore('debug-storage')
export const useShowHeader = createFeatureFlagStore('header-storage')
export const useShowDemoButtons = createFeatureFlagStore('demo-buttons-storage')
export const useFullscreenState = createFeatureFlagStore('fullscreen-storage')
export const useEmulateSlowConnection = createFeatureFlagStore('slow-connection-storage')
export const useVerticalSwipes = createFeatureFlagStore('vertical-swipes-storage', true)
