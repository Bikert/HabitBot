import { NavLink, useNavigate, useParams } from 'react-router'
import { getCurrentDateApiString, getRelativeDate, isValidDateString, toDate } from '../../utils/date'
import classNames from 'classnames'
import { useEffect } from 'react'
import { useSuspenseQuery } from '@tanstack/react-query'
import { delay } from '../../utils/delay'
import { habitsApi } from '../../api/habitsApi'
import { HabitOnDateRow } from './HabitOnDateRow'
import { habitsOnDateQueryKey } from './queryKey'
import { useEmulateSlowConnection } from '../../stores/featureFlagsStores'
import type { DateApiString } from '../../types/DateFormat'

interface DayViewInternalProps {
  date: DateApiString
}

function HabitDateNavLink({ date: dateApiString }: { date: DateApiString }) {
  const date = toDate(dateApiString)
  const shortDateString = `${date.getDate().toString().padStart(2, '0')}.${(date.getMonth() + 1).toString().padStart(2, '0')}`
  const fullDateString = `${shortDateString}.${date.getFullYear()}`
  return (
    <NavLink
      className={({ isActive, isPending }) =>
        classNames(
          'rounded-xl px-2 py-1',
          isActive ? 'bg-tg-secondary-bg pointer-events-none' : 'bg-tg-button',
          isPending && 'pointer-events-none animate-pulse',
        )
      }
      to={`../${dateApiString}`}
    >
      {({ isActive }) => (isActive ? fullDateString : shortDateString)}
    </NavLink>
  )
}

export const DayViewInternal = ({ date }: DayViewInternalProps) => {
  const emulateSlowConnection = useEmulateSlowConnection((state) => state.active)
  const { data: habitsOnDate, isFetching } = useSuspenseQuery({
    queryKey: habitsOnDateQueryKey(date),
    queryFn: async () => {
      if (emulateSlowConnection) {
        await delay(1500)
      }
      return habitsApi.apiHabitCompletionDateGet({
        date,
      })
    },
    staleTime: 30_000,
  })

  return (
    <>
      <div className="flex justify-center gap-2 *:first:before:content-['<'] *:last:after:content-['>']">
        <HabitDateNavLink date={getRelativeDate(date, -2)} />
        <HabitDateNavLink date={getRelativeDate(date, -1)} />
        <HabitDateNavLink date={getRelativeDate(date, 0)} />
        <HabitDateNavLink date={getRelativeDate(date, 1)} />
        <HabitDateNavLink date={getRelativeDate(date, 2)} />
      </div>
      <ul className={classNames('m-4 flex list-none flex-col gap-2', isFetching && '*:opacity-50')}>
        {habitsOnDate.map((habitOnDate) => (
          <HabitOnDateRow
            key={`${date}:${habitOnDate.habit.versionId}`}
            date={date}
            habit={habitOnDate.habit}
            completed={habitOnDate.completed}
          />
        ))}
      </ul>
    </>
  )
}

export const DayView = () => {
  const date = useParams()['date']
  const navigate = useNavigate()

  const isValidDate = isValidDateString(date)
  useEffect(() => {
    if (!isValidDate) navigate('/day/' + getCurrentDateApiString(), { replace: true })
  }, [isValidDate, navigate])

  if (!isValidDate) {
    return null
  }

  return <DayViewInternal date={date} />
}
