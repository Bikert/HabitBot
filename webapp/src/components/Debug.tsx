import { TelegramWebApp } from '../telegram'
import { useDebugStore } from '../stores/useDebugStore'

export function Debug() {
  const user = TelegramWebApp.initDataUnsafe?.user
  const debug = useDebugStore((state) => state.enabled)
  return (
    <>
      {debug && (
        <>
          <div>Version: {TelegramWebApp.version}</div>
          {user?.id}
          <div>{TelegramWebApp.initData}</div>
          <div>{window.location.href}</div>
        </>
      )}
    </>
  )
}
