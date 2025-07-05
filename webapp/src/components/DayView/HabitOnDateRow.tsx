import { useMutation, useQueryClient } from '@tanstack/react-query'
import { delay } from '../../utils/delay'
import { habitsApi } from '../../api/habitsApi'
import classNames from 'classnames'
import type { HabitsHabitCompletionDto } from '@habit-bot/api-client'
import { habitsOnDateQueryKey } from './queryKey'
import { useEmulateSlowConnection } from '../../stores/featureFlagsStores'
import type { DateApiString } from '../../types/DateFormat'
import { NavLink } from 'react-router'

type HabitOnDateRowProps = {
  completed: boolean
  habit: HabitsHabitCompletionDto['habit']
  date: DateApiString
}

export function HabitOnDateRow({ completed, habit, date }: HabitOnDateRowProps) {
  const queryClient = useQueryClient()
  const emulateSlowConnection = useEmulateSlowConnection((state) => state.active)

  const { mutate: updateHabitCompletion, isPending } = useMutation({
    mutationFn: async (newCompleted: boolean) => {
      if (emulateSlowConnection) {
        await delay(1500)
      }
      if (!habit.versionId) {
        throw new Error('versionId is required')
      }
      await habitsApi.apiHabitVersionIdDatePatch({
        date: date,
        versionId: habit.versionId,
        request: {
          completed: newCompleted,
        },
      })
      return {
        versionId: habit.versionId,
        completed: newCompleted,
      }
    },
    onSuccess: () =>
      queryClient.invalidateQueries({
        queryKey: habitsOnDateQueryKey(date),
      }),
  })

  return (
    <li
      key={habit.versionId}
      className={classNames(
        'flex gap-2 rounded-xl p-3',
        isPending && 'pointer-events-none animate-pulse',
        completed && 'opacity-50',
      )}
      style={{
        backgroundColor: `color-mix(in srgb, ${habit.color}, transparent 40%)`,
      }}
    >
      <div
        onClick={() => updateHabitCompletion(!completed)}
        className={classNames('relative flex min-w-0 grow flex-row items-center justify-stretch')}
      >
        <span>{habit.icon}</span>
        <span className="min-w-0 overflow-x-clip text-nowrap text-ellipsis">{habit.name}</span>
        {completed && <div className="bg-tg-text absolute top-1/2 box-border h-[1px] w-full px-11" />}
      </div>
      <NavLink to={`/habit/${habit.id}`} className="flex">
        <span className="material-icons">edit</span>
      </NavLink>
    </li>
  )
}
