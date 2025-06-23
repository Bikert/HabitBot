import { TelegramWebApp } from '../telegram'
import { useDebugInformationVisibility } from '../stores/visibilityStores'

export function DebugView() {
  const user = TelegramWebApp.initDataUnsafe?.user
  const debug = useDebugInformationVisibility((state) => state.visible)
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
