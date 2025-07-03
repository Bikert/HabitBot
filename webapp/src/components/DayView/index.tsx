import { type LoaderFunction, NavLink, redirect, useNavigate, useParams } from 'react-router'
import { getCurrentDateApiString, getRelativeDate, isValidDateString, toDate } from '../../utils/date'
import classNames from 'classnames'
import { useEffect } from 'react'
import { useSuspenseQuery } from '@tanstack/react-query'
import { HabitOnDateRow } from './HabitOnDateRow'
import { useEmulateSlowConnection } from '../../stores/featureFlagsStores'
import type { DateApiString } from '../../types/DateFormat'
import { habitsOnDateQueryOptions } from './habitsOnDateQueryOptions'
import { delay } from '../../utils/delay'
import { queryClient } from '../../api/queryClient'

interface DayViewInternalProps {
  date: DateApiString
}

function HabitDateNavLink({ date: dateApiString }: { date: DateApiString }) {
  const date = toDate(dateApiString)
  const shortDateString = `${date.getDate().toString().padStart(2, '0')}.${(date.getMonth() + 1).toString().padStart(2, '0')}`
  const fullDateString = `${shortDateString}.${date.getFullYear()}`
  return (
    <NavLink
      style={{
        viewTransitionName: 'match-element',
      }}
      className={({ isActive, isPending }) =>
        classNames(
          'rounded-xl px-2 py-1',
          isActive ? 'bg-tg-secondary-bg pointer-events-none' : 'bg-tg-button',
          isPending && 'pointer-events-none',
        )
      }
      to={`../${dateApiString}`}
      viewTransition
    >
      {({ isActive }) => (isActive ? fullDateString : shortDateString)}
    </NavLink>
  )
}

export const dayViewLoader: LoaderFunction = async ({ params }) => {
  const date = params['date']
  if (!isValidDateString(date)) {
    return redirect('/day/' + getCurrentDateApiString())
  }
  await queryClient.prefetchQuery(habitsOnDateQueryOptions(date))
  return delay(1)
}

export const DayViewInternal = ({ date }: DayViewInternalProps) => {
  const emulateSlowConnection = useEmulateSlowConnection((state) => state.active)
  const { data: habitsOnDate, isFetching } = useSuspenseQuery(habitsOnDateQueryOptions(date, emulateSlowConnection))

  return (
    <>
      <div className="flex justify-center gap-2 *:first:before:content-['<'] *:last:after:content-['>']">
        {[-2, -1, 0, 1, 2].map((offset) => {
          const relativeDate = getRelativeDate(date, offset)
          return <HabitDateNavLink key={relativeDate} date={relativeDate} />
        })}
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
