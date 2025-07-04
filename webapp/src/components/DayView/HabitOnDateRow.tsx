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
        'flex justify-between gap-2 rounded-xl p-3',
        isPending && 'pointer-events-none animate-pulse',
      )}
      style={{
        backgroundColor: `color-mix(in srgb, ${habit.color}, transparent 40%)`,
      }}
    >
      <div className="flex grow flex-row items-center justify-stretch">
        <NavLink to={`/habit/${habit.id}`}>✏️</NavLink>
        <div
          onClick={() => updateHabitCompletion(!completed)}
          className={classNames(
            'grow text-center',
            completed &&
              'after:bg-tg-text relative box-border line-through after:absolute after:top-1/2 after:block after:h-[1px] after:w-full after:px-11 after:content-[""]',
          )}
        >
          {habit.icon}
          {habit.name}
          {habit.icon}
        </div>
      </div>
      <div className="flex items-center">
        <input
          type="checkbox"
          className="checked:bg-tg-button h-5 w-5 cursor-pointer appearance-none rounded-full border-3"
          onChange={() => updateHabitCompletion(!completed)}
          checked={completed}
        />
      </div>
    </li>
  )
}
