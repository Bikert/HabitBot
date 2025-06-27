import { Outlet, useLocation, useNavigate } from 'react-router'
import { BackButton, MainButton, SecondaryButton } from './telegram/components'
import { DebugView } from './components/DebugView'
import { useShowDemoButtons, useShowHeader } from './stores/featureFlagsStores'
import { NavigationButtons } from './components/NavigationButtons'
import { useTelegramInit } from './utils/useTelegramInit'
import { useNavigateBackOrClose } from './utils/useNavigateBackOrClose'
import { SettingsButton } from './telegram/components/SettingsButton'

function App() {
  const showHeader = useShowHeader((state) => state.active)
  const showDemoButtons = useShowDemoButtons((state) => state.active)
  useTelegramInit()
  const goBack = useNavigateBackOrClose()
  const navigate = useNavigate()
  const location = useLocation()

  return (
    <div className="mt-tg-content-safe-top mb-tg-content-safe-bottom ml-tg-content-safe-left mr-tg-content-safe-right max-h-svh p-2">
      {showHeader && <h1 className="p-2 text-center text-3xl font-bold">HabitBot {location.pathname}</h1>}
      <BackButton onClick={goBack} />
      <SettingsButton onClick={() => navigate('/config')} />
      <Outlet />
      {showDemoButtons && <MainButton text="submit" />}
      {showDemoButtons && <SecondaryButton text="secondary" />}
      <DebugView />
      <NavigationButtons />
    </div>
  )
}

export default App
