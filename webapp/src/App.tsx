import { Outlet, useNavigate } from 'react-router'
import { BackButton, MainButton, SecondaryButton } from './telegram/components'
import { DebugView } from './components/DebugView'
import { useShowDemoButtons, useShowHeader } from './stores/featureFlagsStores'
import { NavigationButtons } from './components/NavigationButtons'
import { useTelegramInit } from './utils/useTelegramInit'
import { useNavigateBackOrClose } from './utils/useNavigateBackOrClose'
import { SettingsButton } from './telegram/components/SettingsButton'
import { QueryClientProvider } from '@tanstack/react-query'
import { Suspense } from 'react'
import { queryClient } from './api/queryClient'
import { useViewportHeight } from './utils/useViewportHeight'
import { PATHS } from './constants/paths'
import { Toaster } from 'sonner'
import { useTelegramTheme } from './stores/useTelegramTheme'
import { useSharedNavigationBlocker } from './utils/useSharedNavigationBlocker'

function AppHeader() {
  const showHeader = useShowHeader((state) => state.active)
  return (
    <div className="bg-background backdrop-opacity-50">
      <div className="min-h-tg-content-safe-top-inset">
        {showHeader && <h1 className="p-2 text-center text-3xl font-bold">HabitBot {location.pathname}</h1>}
      </div>
    </div>
  )
}

function AppFooter() {
  return (
    <div className="rounded-t-4xl bg-tg-bottom-bar-bg pb-tg-safe-bottom">
      <NavigationButtons />
    </div>
  )
}

function App() {
  const theme = useTelegramTheme((state) => state.theme)
  const showDemoButtons = useShowDemoButtons((state) => state.active)
  useTelegramInit()
  const goBack = useNavigateBackOrClose()
  const navigate = useNavigate()

  const viewportHeight = useViewportHeight()
  const fixedLayoutElements = viewportHeight && viewportHeight > 560

  useSharedNavigationBlocker()

  return (
    <QueryClientProvider client={queryClient}>
      {/* Out-of-page native telegram elements */}
      <BackButton onClick={goBack} />
      <SettingsButton onClick={() => navigate(PATHS.settings)} />
      {showDemoButtons && <MainButton text="submit" />}
      {showDemoButtons && <SecondaryButton text="secondary" />}
      {/* Layout container */}
      <div className="relative box-border flex h-full flex-col pt-tg-safe-top pr-tg-safe-right pl-tg-safe-left">
        <Toaster theme={theme} position="top-right" swipeDirections={['left', 'right']} />
        {fixedLayoutElements && <AppHeader />}
        {/* Non-scrollable positioned container for absolute elements */}
        <div className="relative flex min-h-0 shrink grow">
          {/* Scrollable container */}
          <div className="flex min-w-0 grow flex-col overflow-y-auto">
            {!fixedLayoutElements && <AppHeader />}
            {/* Content goes here */}
            <div className="w-full max-w-xl grow self-center py-4">
              <Suspense fallback={<div>Loading...</div>}>
                <Outlet />
              </Suspense>
              <DebugView />
            </div>
            {!fixedLayoutElements && <AppFooter />}
          </div>
        </div>
        {fixedLayoutElements && <AppFooter />}
      </div>
    </QueryClientProvider>
  )
}

export default App
