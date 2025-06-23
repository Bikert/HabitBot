import { useCallback, useEffect, useState } from 'react'
import { Outlet, useLocation, useSearchParams } from 'react-router'
import { BackButton, MainButton, SecondaryButton } from './telegram/components'
import { TelegramWebApp } from './telegram'
import { Debug } from './components/Debug'
import { useDebugStore } from './stores/useDebugStore'
import { useHeaderVisibility } from './stores/useHeaderVisibility'
import { useViewportHeight } from './utils/useViewportHeight'
import classNames from 'classnames'

function App() {
  const [searchParams, setSearchParams] = useSearchParams()
  const showDebugButtons = searchParams.get('debug') === 'true'
  const toggleDebugButtons = useCallback(() => {
    setSearchParams((prev) => ({ debug: prev.get('debug') === 'true' ? 'false' : 'true' }))
  }, [setSearchParams])
  const toggleDebug = useDebugStore((state) => state.toggle)
  const toggleHeader = useHeaderVisibility((state) => state.toggle)
  const showHeader = useHeaderVisibility((state) => state.visible)
  const [showDemoButtons, setShowDemoButtons] = useState(false)
  const location = useLocation()
  useEffect(() => {
    TelegramWebApp.expand()
    TelegramWebApp.ready()
  }, [])
  const viewportHeight = useViewportHeight()
  const navFixed = viewportHeight && viewportHeight > 600

  return (
    <div className="max-h-svh">
      {showHeader && <h1 className="p-3 text-center text-3xl font-bold">HabitBot {location.pathname}</h1>}
      <Outlet />
      {showDebugButtons && (
        <div className="text-tg-button-text flex w-full justify-center gap-2">
          <button className="bg-tg-button rounded-l-xl p-2" onClick={() => setShowDemoButtons((show) => !show)}>
            toggle buttons
          </button>
          <button className="bg-tg-button p-2" onClick={toggleHeader}>
            toggle header
          </button>
          <button className="bg-tg-button rounded-r-xl p-2" onClick={toggleDebug}>
            toggle debug information
          </button>
        </div>
      )}
      {showDemoButtons && (
        <>
          <MainButton text="submit" />
          <SecondaryButton text="secondary" />
          <BackButton />
        </>
      )}
      <Debug />
      <div
        className={classNames('right-0 left-0 mx-auto flex w-fit gap-3', navFixed ? 'fixed bottom-4' : 'static pb-4')}
      >
        <button
          onClick={() => (window.location.href = '/habit')}
          className="bg-tg-button text-tg-button-text ring-tg-bg flex h-14 w-14 items-center justify-center rounded-full text-2xl ring-2"
        >
          ➕
        </button>
        <button
          onClick={() => (window.location.href = '/')}
          className="bg-tg-button text-tg-button-text ring-tg-bg flex h-14 w-14 items-center justify-center rounded-full text-2xl ring-2"
        >
          📋
        </button>
        <button
          onClick={toggleDebugButtons}
          className="bg-tg-button text-tg-button-text ring-tg-bg flex h-14 w-14 items-center justify-center rounded-full text-2xl ring-2"
        >
          🐛
        </button>
      </div>
    </div>
  )
}

export default App
