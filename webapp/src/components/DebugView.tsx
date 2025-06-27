import { TelegramWebApp } from '../telegram'
import { useShowDebugInformation } from '../stores/featureFlagsStores'

export function DebugView() {
  const user = TelegramWebApp.initDataUnsafe?.user
  const debug = useShowDebugInformation((state) => state.active)
  if (!debug) return null
  return (
    <div className="flex max-w-screen flex-col gap-2 overflow-x-hidden px-2 py-4 whitespace-nowrap">
      <div>Version: {TelegramWebApp.version}</div>
      <div className="flex gap-4">
        <button
          className="bg-tg-button cursor-pointer rounded-xl"
          onClick={() => {
            if (user?.id) {
              return navigator.clipboard.writeText(user.id.toString())
            }
          }}
        >
          📋User ID📋
        </button>
        <div>{user?.id}</div>
      </div>
      <div className="flex gap-4">
        <button
          className="bg-tg-button cursor-pointer rounded-xl"
          onClick={() => navigator.clipboard.writeText(TelegramWebApp.initData)}
        >
          📋Init data📋
        </button>
        <div className="overflow-x-hidden text-ellipsis">{TelegramWebApp.initData}</div>
      </div>
      <div className="flex gap-4">
        <button
          className="bg-tg-button text-tg-button-text cursor-pointer rounded-xl"
          onClick={() => navigator.clipboard.writeText(window.location.href)}
        >
          📋Location📋
        </button>
        <div className="overflow-x-hidden text-ellipsis">{window.location.href}</div>
      </div>
    </div>
  )
}
