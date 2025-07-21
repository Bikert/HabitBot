import { create } from 'zustand/index'

export interface BlockerState {
  blockers: Set<() => void>
  addBlocker: (blocker: () => void) => void
  removeBlocker: (blocker: () => void) => void
}

export const useBlockersStore = create<BlockerState>()((set) => ({
  blockers: new Set<() => void>(),
  addBlocker: (blocker: () => void) =>
    set((state) => ({
      blockers: new Set(state.blockers).add(blocker),
    })),
  removeBlocker: (blocker: () => void) =>
    set((state) => {
      const blockers = new Set(state.blockers)
      blockers.delete(blocker)
      return { blockers }
    }),
}))
