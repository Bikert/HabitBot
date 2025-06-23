import { TelegramWebApp } from './index'
import { StateStorage } from 'zustand/middleware'
import { promisify } from '../utils/promisify'
import type { DeviceStorage } from './types'

const preferLocalStorage = true as boolean

export const tgStorage = TelegramWebApp.DeviceStorage
const TelegramDeviceStorage: StateStorage = tgStorage
  ? {
      getItem: promisify(tgStorage.getItem),
      setItem: promisify(tgStorage.setItem),
      removeItem: promisify(tgStorage.removeItem),
    }
  : localStorage

const DeviceStorage = preferLocalStorage ? window.localStorage : TelegramDeviceStorage
export { DeviceStorage }
