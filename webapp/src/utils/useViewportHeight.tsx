import { useEffect, useState } from 'react'

export const useViewportHeight = () => {
  const [viewportHeight, setViewportHeight] = useState(window.visualViewport?.height ?? window.innerHeight)

  useEffect(() => {
    const visualViewport = window.visualViewport
    if (visualViewport) {
      const listener = () => {
        setViewportHeight(visualViewport.height)
      }
      visualViewport.addEventListener('resize', listener)
      return () => {
        visualViewport.removeEventListener('resize', listener)
      }
    } else {
      const listener = () => {
        setViewportHeight(window.innerHeight)
      }
      window.addEventListener('resize', listener)
      return () => {
        window.removeEventListener('resize', listener)
      }
    }
  }, [])

  return viewportHeight
}
