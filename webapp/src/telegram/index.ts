import './telegram-web-app.js'
import type { WebApp } from './types'

// @ts-expect-error js lib filling the global object
export const TelegramWebApp = window.Telegram.WebApp as WebApp
