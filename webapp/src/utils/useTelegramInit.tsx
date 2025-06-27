import { useEffect } from 'react'
import { TelegramWebApp } from '../telegram'
import { useFullscreenState, useVerticalSwipes } from '../stores/featureFlagsStores'

export function useTelegramInit() {
  useEffect(() => {
    TelegramWebApp.expand()
    TelegramWebApp.ready()
  }, [])

  const fullscreen = useFullscreenState((state) => state.active)
  useEffect(() => {
    if (fullscreen && !TelegramWebApp.isFullscreen) {
      TelegramWebApp.requestFullscreen()
    }
    if (!fullscreen && TelegramWebApp.isFullscreen) {
      TelegramWebApp.exitFullscreen()
    }
  }, [fullscreen])

  const verticalSwipes = useVerticalSwipes((state) => state.active)
  useEffect(() => {
    if (verticalSwipes && !TelegramWebApp.isVerticalSwipesEnabled) {
      TelegramWebApp.enableVerticalSwipes()
    }
    if (!verticalSwipes && TelegramWebApp.isVerticalSwipesEnabled) {
      TelegramWebApp.disableVerticalSwipes()
    }
  }, [verticalSwipes])
}
