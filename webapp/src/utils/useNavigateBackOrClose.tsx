import { useNavigate } from 'react-router'
import { useCallback } from 'react'
import { TelegramWebApp } from '../telegram'

export function useNavigateBackOrClose() {
  const navigate = useNavigate()
  return useCallback(() => {
    if (history.state?.idx === 0) {
      TelegramWebApp.close()
    } else {
      navigate(-1)
    }
  }, [navigate])
}
