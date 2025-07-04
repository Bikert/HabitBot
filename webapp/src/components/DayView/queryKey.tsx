import type { DateApiString } from '../../types/DateFormat'

export function habitsOnDateQueryKey(date: DateApiString) {
  return ['habitsOnDate', date] as const
}
