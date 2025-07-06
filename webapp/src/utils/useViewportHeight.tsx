import { useEffect, useState } from 'react'
import { throttle } from './throttle'

const UPDATE_DELAY_MS = 200

export const useViewportHeight = () => {
  const [viewportHeight, setViewportHeight] = useState(window.visualViewport?.height ?? window.innerHeight)

  useEffect(() => {
    const updater = throttle(setViewportHeight, UPDATE_DELAY_MS)
    const visualViewport = window.visualViewport
    if (visualViewport) {
      const listener = () => {
        updater(visualViewport.height)
      }
      visualViewport.addEventListener('resize', listener)
      return () => {
        visualViewport.removeEventListener('resize', listener)
      }
    } else {
      const listener = () => {
        updater(window.innerHeight)
      }
      window.addEventListener('resize', listener)
      return () => {
        window.removeEventListener('resize', listener)
      }
    }
  }, [])

  return viewportHeight
}
