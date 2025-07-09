import { TelegramWebApp } from '../telegram'
import { useShowDebugInformation } from '../stores/featureFlagsStores'
import { toast } from 'sonner'

export function DebugView() {
  const user = TelegramWebApp.initDataUnsafe?.user
  const debug = useShowDebugInformation((state) => state.active)
  if (!debug) return null
  const currentLocation = window.location.href
  const initialLocation = sessionStorage.getItem('initialLocation')
  return (
    <div className="flex max-w-screen flex-col gap-2 overflow-x-hidden px-2 py-4 whitespace-nowrap">
      <div>Version: {TelegramWebApp.version}</div>
      <div className="flex gap-4">
        <button
          className="bg-tg-button cursor-pointer rounded-xl"
          onClick={async () => {
            if (user?.id) {
              await navigator.clipboard.writeText(user.id.toString())
              toast.info('Copied user ID')
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
          onClick={async () => {
            await navigator.clipboard.writeText(TelegramWebApp.initData)
            toast.info('Copied init data')
          }}
        >
          📋Init data📋
        </button>
        <div className="overflow-x-hidden text-ellipsis">{TelegramWebApp.initData}</div>
      </div>
      <div className="flex gap-4">
        <button
          className="bg-tg-button cursor-pointer rounded-xl"
          onClick={async () => {
            await navigator.clipboard.writeText(currentLocation)
            toast.info('Copied current location')
          }}
        >
          📋Location📋
        </button>
        <div className="overflow-x-hidden text-ellipsis">{currentLocation}</div>
      </div>
      <div className="flex gap-4">
        <button
          className="bg-tg-button cursor-pointer rounded-xl"
          onClick={async () => {
            await navigator.clipboard.writeText(initialLocation ?? '')
            toast.info('Copied initial location')
          }}
        >
          📋Initial location📋
        </button>
        <div className="overflow-x-hidden text-ellipsis">{initialLocation}</div>
      </div>
    </div>
  )
}
