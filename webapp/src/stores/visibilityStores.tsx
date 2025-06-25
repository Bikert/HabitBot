import { create } from 'zustand/index'
import { createJSONStorage, persist } from 'zustand/middleware'
import { immer } from 'zustand/middleware/immer'
import { DeviceStorage } from '../telegram/telegramStorage'

export interface VisibilityState {
  visible: boolean
  show: () => void
  hide: () => void
  toggle: () => void
}
export interface ConfigState {
  debugInformation: VisibilityState
  header: VisibilityState
  demoButtons: VisibilityState
}

function createVisibilityStore(storeName: string) {
  return create<VisibilityState>()(
    persist(
      immer((set) => ({
        visible: false,
        show: () =>
          set((state) => {
            state.visible = true
          }),
        hide: () =>
          set((state) => {
            state.visible = false
          }),
        toggle: () =>
          set((state) => {
            state.visible = !state.visible
          }),
      })),
      {
        name: storeName,
        storage: createJSONStorage(() => DeviceStorage),
      },
    ),
  )
}

export const useDebugInformationVisibility = createVisibilityStore('debug-storage')
export const useHeaderVisibility = createVisibilityStore('header-storage')
export const useDemoButtonsVisibility = createVisibilityStore('demo-buttons-storage')
