import { useBlockersStore } from '../stores/useBlockersStore'
import { useEffect } from 'react'

export function useRegisterBlockerCallback({
  blockerCallback,
  isBlocked,
}: {
  blockerCallback: () => void
  isBlocked: boolean
}) {
  const addBlocker = useBlockersStore((state) => state.addBlocker)
  const removeBlocker = useBlockersStore((state) => state.removeBlocker)
  useEffect(() => {
    if (isBlocked) {
      addBlocker(blockerCallback)
      return () => {
        removeBlocker(blockerCallback)
      }
    }
  }, [addBlocker, blockerCallback, isBlocked, removeBlocker])
}
