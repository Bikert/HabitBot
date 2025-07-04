import { Configuration, DefaultConfig, HabitsApi } from '@habit-bot/api-client'
import { TelegramWebApp } from '../telegram'

export const habitsApi = new HabitsApi(
  new Configuration({
    ...DefaultConfig,
    basePath: window.location.origin,
    headers: {
      'x-telegram-init-data': TelegramWebApp.initData,
    },
  }),
)
