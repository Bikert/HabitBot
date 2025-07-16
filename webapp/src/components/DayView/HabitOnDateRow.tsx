import { useMutation, useQueryClient } from '@tanstack/react-query'
import { delay } from '../../utils/delay'
import { habitsApi } from '../../api/habitsApi'
import classNames from 'classnames'
import type { HabitsHabitCompletionDto } from '@habit-bot/api-client'
import { habitsOnDateQueryKey } from './queryKey'
import { useEmulateSlowConnection } from '../../stores/featureFlagsStores'
import type { DateApiString } from '../../types/DateFormat'
import { NavLink } from 'react-router'
import { useCallback } from 'react'
import { TelegramWebApp } from '../../telegram'

type HabitOnDateRowProps = {
  completed: boolean
  habit: HabitsHabitCompletionDto['habit']
  date: DateApiString
  flippedProps: object
}

export const HabitOnDateRow = ({ completed, habit, date, flippedProps }: HabitOnDateRowProps) => {
  const queryClient = useQueryClient()
  const emulateSlowConnection = useEmulateSlowConnection((state) => state.active)

  const mutation = useMutation({
    mutationFn: async (newCompleted: boolean) => {
      if (!habit.versionId) {
        throw new Error('versionId is required')
      }
      if (emulateSlowConnection) {
        await delay(1500)
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
    onSettled: () => {
      return queryClient.invalidateQueries({
        queryKey: habitsOnDateQueryKey(date),
      })
    },
  })

  const { mutate, isPending } = mutation
  const updateCompletionStatus = useCallback(() => {
    TelegramWebApp.HapticFeedback.selectionChanged()
    mutate(!completed)
  }, [mutate, completed])

  const optimisticCompleted = mutation.isPending ? mutation.variables : completed

  return (
    <li
      {...flippedProps}
      className={classNames(
        'flex gap-2 rounded-xl px-3',
        isPending ? 'pointer-events-none animate-pulse' : 'cursor-pointer',
        !isPending && completed && 'opacity-50',
      )}
      style={{
        backgroundColor: `color-mix(in srgb, ${habit.color}, transparent 40%)`,
      }}
    >
      <div
        onClick={updateCompletionStatus}
        className={classNames('relative flex min-w-0 grow flex-row items-center justify-stretch py-3')}
      >
        <span>{habit.icon}</span>
        <span className="min-w-0 overflow-x-clip text-nowrap text-ellipsis">{habit.name}</span>
        {optimisticCompleted && <div className="absolute top-1/2 box-border h-[1px] w-full bg-surface px-11" />}
      </div>
      <NavLink to={`/habit/${habit.id}`} className="flex items-center">
        <span className="material-icons">edit</span>
      </NavLink>
    </li>
  )
}
