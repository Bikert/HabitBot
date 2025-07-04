import type { DateApiString } from '../../types/DateFormat'
import { habitsOnDateQueryKey } from './queryKey'
import { delay } from '../../utils/delay'
import { habitsApi } from '../../api/habitsApi'
import { queryOptions } from '@tanstack/react-query'

export function habitsOnDateQueryOptions(date: DateApiString, emulateSlowConnection: boolean = false) {
  return queryOptions({
    queryKey: habitsOnDateQueryKey(date),
    queryFn: async () => {
      if (emulateSlowConnection) {
        await delay(1500)
      }
      return habitsApi.apiHabitCompletionDateGet({
        date,
      })
    },
    staleTime: 5_000,
  })
}
