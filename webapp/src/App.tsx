import { useEffect, useState } from 'react'
import { Outlet, useLocation, useSearchParams } from 'react-router'
import { BackButton, MainButton, SecondaryButton } from './telegram/components'
import { TelegramWebApp } from './telegram'
import { Debug } from './components/Debug'
import { useDebugStore } from './stores/useDebugStore'
import { useHeaderVisibility } from './stores/useHeaderVisibility'

function App() {
  const [searchParams] = useSearchParams()
  const showAdditionalButtons = searchParams.get('debug') === 'true'
  const toggleDebug = useDebugStore((state) => state.toggle)
  const toggleHeader = useHeaderVisibility((state) => state.toggle)
  const showHeader = useHeaderVisibility((state) => state.visible)
  const [showDemoButtons, setShowDemoButtons] = useState(false)
  const location = useLocation()
  useEffect(() => {
    TelegramWebApp.expand()
    TelegramWebApp.ready()
  }, [])

  return (
    <div>
      {showHeader && <h1 className="p-3 text-center text-3xl font-bold">HabitBot {location.pathname}</h1>}

      <Outlet />

      {showAdditionalButtons && (
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
      <Debug />

      {showDemoButtons && (
        <>
          <MainButton text="submit" />
          <SecondaryButton text="secondary" />
          <BackButton />
        </>
      )}
    </div>
  )
}

export default App
