import { useEffect } from 'react'
import { NavLink, Outlet, useLocation } from 'react-router'
import { BackButton, MainButton, SecondaryButton } from './telegram/components'
import { TelegramWebApp } from './telegram'
import { DebugView } from './components/DebugView'
import { useDemoButtonsVisibility, useHeaderVisibility } from './stores/visibilityStores'
import { useViewportHeight } from './utils/useViewportHeight'
import classNames from 'classnames'

function App() {
  const showHeader = useHeaderVisibility((state) => state.visible)
  const showDemoButtons = useDemoButtonsVisibility((state) => state.visible)
  const location = useLocation()
  useEffect(() => {
    TelegramWebApp.expand()
    TelegramWebApp.ready()
  }, [])
  const viewportHeight = useViewportHeight()
  const navFixed = viewportHeight && viewportHeight > 500

  return (
    <div className="max-h-svh">
      {showHeader && <h1 className="p-3 text-center text-3xl font-bold">HabitBot {location.pathname}</h1>}
      <Outlet />
      {showDemoButtons && (
        <>
          <MainButton text="submit" />
          <SecondaryButton text="secondary" />
          <BackButton />
        </>
      )}
      <DebugView />
      <div
        className={classNames(
          'right-0 left-0 mx-auto flex w-fit gap-3 select-none',
          navFixed ? 'fixed bottom-4' : 'static pb-4',
        )}
      >
        <NavLink
          to="/habit"
          className={({ isActive, isPending }) =>
            classNames(
              'text-tg-button-text ring-tg-bg flex h-14 w-14 items-center justify-center rounded-full text-2xl ring-2',
              isActive ? 'bg-tg-button-text' : 'bg-tg-button',
              isPending && 'pointer-events-none animate-spin',
            )
          }
        >
          ➕
        </NavLink>
        <NavLink
          to="/day"
          className={({ isActive, isPending }) =>
            classNames(
              'text-tg-button-text ring-tg-bg flex h-14 w-14 items-center justify-center rounded-full text-2xl ring-2',
              isActive ? 'bg-tg-button-text' : 'bg-tg-button',
              isPending && 'pointer-events-none animate-spin',
            )
          }
        >
          📋
        </NavLink>
        <NavLink
          to="/config"
          className={({ isActive, isPending }) =>
            classNames(
              'text-tg-button-text ring-tg-bg flex h-14 w-14 items-center justify-center rounded-full text-2xl ring-2',
              isActive ? 'bg-tg-button-text' : 'bg-tg-button',
              isPending && 'pointer-events-none animate-spin',
            )
          }
        >
          🔧
        </NavLink>
      </div>
    </div>
  )
}

export default App
