import { TelegramWebApp } from '../telegram'
import { useDebugStore } from '../stores/useDebugStore'

export function Debug() {
  const user = TelegramWebApp.initDataUnsafe?.user
  const debug = useDebugStore((state) => state.enabled)
  if (!debug) return null
  return (
    <div className="max-w-screen overflow-x-hidden whitespace-nowrap">
      <div>Version: {TelegramWebApp.version}</div>
      {user?.id}
      <div className="overflow-x-hidden text-ellipsis">{TelegramWebApp.initData}</div>
      <div>{window.location.href}</div>
    </div>
  )
}
