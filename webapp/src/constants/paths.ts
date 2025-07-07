import type { DateApiString } from '../types/DateFormat'
import { isValidDateString } from '../utils/date'

export const PATHS = Object.freeze({
  root: '/',
  editHabit: (id?: string) => {
    if (!id) return '/habit'
    return `/habit/${id}`
  },
  habitsList: '/habits-list',
  settings: '/settings',
  day: (date?: DateApiString) => {
    if (isValidDateString(date)) return `/day/${date}`
    return '/day'
  },
  bodyMeasurements: '/body-measurements',
})
