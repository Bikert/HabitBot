import { BodyMetricsApi, Configuration, DefaultConfig } from '@habit-bot/api-client'
import { TelegramWebApp } from '../telegram'

export const bodyMetricApi = new BodyMetricsApi(
  new Configuration({
    ...DefaultConfig,
    basePath: window.location.origin,
    headers: {
      'x-telegram-init-data': TelegramWebApp.initData,
    },
  }),
)
