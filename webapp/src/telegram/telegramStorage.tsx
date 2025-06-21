import { TelegramWebApp } from './index'
import { StateStorage } from 'zustand/middleware'
import { promisify } from '../utils/promisify'

export const tgStorage = TelegramWebApp.DeviceStorage
console.log('window.Telegram.WebApp.DeviceStorage', tgStorage)
const TelegramDeviceStorage: StateStorage = tgStorage
  ? {
      getItem: promisify(tgStorage.getItem),
      setItem: promisify(tgStorage.setItem),
      removeItem: promisify(tgStorage.removeItem),
    }
  : localStorage

export { TelegramDeviceStorage as DeviceStorage }
