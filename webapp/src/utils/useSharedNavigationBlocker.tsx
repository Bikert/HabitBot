import { useBlockersStore } from '../stores/useBlockersStore'
import { BlockerFunction, useBlocker } from 'react-router'
import { useCallback, useEffect } from 'react'

export function useSharedNavigationBlocker() {
  const blockers = useBlockersStore((state) => state.blockers)
  const blocker = useBlocker(
    useCallback<BlockerFunction>(
      ({ historyAction }) => historyAction === 'POP' && blockers.size !== 0,
      [blockers.size],
    ),
  )

  useEffect(() => {
    if (blocker.state === 'blocked') {
      for (const blockerCallback of blockers) {
        blockerCallback()
      }
      blocker.reset()
    }
  }, [blocker, blockers])
}
