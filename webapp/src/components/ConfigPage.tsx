import {
  useEmulateSlowConnection,
  useFullscreenState,
  useShowDebugInformation,
  useShowDemoButtons,
  useShowHeader,
  useVerticalSwipes,
} from '../stores/featureFlagsStores'
import type { PropsWithChildren } from 'react'
import classNames from 'classnames'

function ConfigButton({ active, toggle, children }: PropsWithChildren<{ active: boolean; toggle: { (): void } }>) {
  return (
    <button
      className={classNames(
        'ring-tg-button ring-2',
        active ? 'bg-tg-button text-tg-button-text' : 'bg-tg-secondary-bg text-tg-button',
      )}
      onClick={toggle}
    >
      {children}
    </button>
  )
}

export function ConfigPage() {
  const toggleDemoButtons = useShowDemoButtons((state) => state.toggle)
  const demoButtonsActive = useShowDemoButtons((state) => state.active)

  const toggleHeader = useShowHeader((state) => state.toggle)
  const headerActive = useShowHeader((state) => state.active)

  const toggleDebug = useShowDebugInformation((state) => state.toggle)
  const debugActive = useShowDebugInformation((state) => state.active)

  const toggleFullscreen = useFullscreenState((state) => state.toggle)
  const fullscreenActive = useFullscreenState((state) => state.active)

  const toggleVerticalSwipes = useVerticalSwipes((state) => state.toggle)
  const verticalSwipesActive = useVerticalSwipes((state) => state.active)

  const toggleEmulateSlowConnection = useEmulateSlowConnection((state) => state.toggle)
  const emulateSlowConnectionActive = useEmulateSlowConnection((state) => state.active)

  return (
    <div className="text-tg-button-text fl flex w-full flex-wrap justify-center gap-3 *:rounded-xl *:p-2">
      <ConfigButton active={demoButtonsActive} toggle={toggleDemoButtons}>
        Demo buttons
      </ConfigButton>
      <ConfigButton active={headerActive} toggle={toggleHeader}>
        Header
      </ConfigButton>
      <ConfigButton active={debugActive} toggle={toggleDebug}>
        Debug information
      </ConfigButton>
      <ConfigButton active={fullscreenActive} toggle={toggleFullscreen}>
        Fullscreen
      </ConfigButton>
      <ConfigButton active={verticalSwipesActive} toggle={toggleVerticalSwipes}>
        Telegram vertical swipes
      </ConfigButton>
      <ConfigButton active={emulateSlowConnectionActive} toggle={toggleEmulateSlowConnection}>
        Emulate slow connection
      </ConfigButton>
    </div>
  )
}
