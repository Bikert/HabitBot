import {
  useEmulateSlowConnection,
  useFullscreenState,
  useShowDebugInformation,
  useShowDemoButtons,
  useShowHeader,
  useVerticalSwipes,
} from '../stores/featureFlagsStores'

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
    <div className="text-tg-button-text fl flex w-full flex-wrap justify-center gap-2 *:rounded-xl *:p-2">
      <button
        className={demoButtonsActive ? 'bg-tg-button-text text-tg-button' : 'bg-tg-button text-tg-button-text'}
        onClick={toggleDemoButtons}
      >
        Demo buttons
      </button>
      <button
        className={headerActive ? 'bg-tg-button-text text-tg-button' : 'bg-tg-button text-tg-button-text'}
        onClick={toggleHeader}
      >
        Header
      </button>
      <button
        className={debugActive ? 'bg-tg-button-text text-tg-button' : 'bg-tg-button text-tg-button-text'}
        onClick={toggleDebug}
      >
        Debug information
      </button>
      <button
        className={fullscreenActive ? 'bg-tg-button-text text-tg-button' : 'bg-tg-button text-tg-button-text'}
        onClick={toggleFullscreen}
      >
        Fullscreen
      </button>
      <button
        className={verticalSwipesActive ? 'bg-tg-button-text text-tg-button' : 'bg-tg-button text-tg-button-text'}
        onClick={toggleVerticalSwipes}
      >
        Telegram vertical swipes
      </button>
      <button
        className={
          emulateSlowConnectionActive ? 'bg-tg-button-text text-tg-button' : 'bg-tg-button text-tg-button-text'
        }
        onClick={toggleEmulateSlowConnection}
      >
        Emulate slow connection
      </button>
    </div>
  )
}
