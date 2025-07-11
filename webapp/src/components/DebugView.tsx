import { TelegramWebApp } from '../telegram'
import { useShowDebugInformation } from '../stores/featureFlagsStores'
import { toast } from 'sonner'
import { Button } from './common/Button'

export function DebugView() {
  const user = TelegramWebApp.initDataUnsafe?.user
  const debug = useShowDebugInformation((state) => state.active)
  if (!debug) return null
  const currentLocation = window.location.href
  const initialLocation = sessionStorage.getItem('initialLocation')
  return (
    <div className="flex max-w-screen flex-col gap-2 overflow-x-hidden px-2 py-4 text-xs whitespace-nowrap">
      <div>Version: {TelegramWebApp.version}</div>
      <div className="flex items-baseline gap-4">
        <Button
          size="xs"
          variant="secondary"
          className="cursor-pointer"
          onClick={async () => {
            if (user?.id) {
              await navigator.clipboard.writeText(user.id.toString())
              toast.info('Copied user ID')
            }
          }}
        >
          📋User ID
        </Button>
        <div>{user?.id}</div>
      </div>
      <div className="flex items-baseline gap-4">
        <Button
          size="xs"
          variant="secondary"
          className="cursor-pointer"
          onClick={async () => {
            await navigator.clipboard.writeText(TelegramWebApp.initData)
            toast.info('Copied init data')
          }}
        >
          📋Init data
        </Button>
        <div className="overflow-x-hidden text-ellipsis">{TelegramWebApp.initData}</div>
      </div>
      <div className="flex items-baseline gap-4">
        <Button
          size="xs"
          variant="secondary"
          className="cursor-pointer"
          onClick={async () => {
            await navigator.clipboard.writeText(currentLocation)
            toast.info('Copied current location')
          }}
        >
          📋Location
        </Button>
        <div className="overflow-x-hidden text-ellipsis">{currentLocation}</div>
      </div>
      <div className="flex items-baseline gap-4">
        <Button
          size="xs"
          variant="secondary"
          className="cursor-pointer"
          onClick={async () => {
            await navigator.clipboard.writeText(initialLocation ?? '')
            toast.info('Copied initial location')
          }}
        >
          📋Initial location
        </Button>
        <div className="overflow-x-hidden text-ellipsis">{initialLocation}</div>
      </div>
    </div>
  )
}
