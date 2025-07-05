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
import { AddHabitButton } from '../NavigationButtons'

interface DayViewInternalProps {
  date: DateApiString
}

function getDay(date: Date) {
  const options = { weekday: 'short' } as const
  const format = new Intl.DateTimeFormat('en-US', options)
  return format.format(date)
}

function HabitDateNavLink({ date: dateApiString }: { date: DateApiString }) {
  const date = toDate(dateApiString)
  const dayOfWeek = getDay(date)
  const dayOfMonth = date.getDate()
  return (
    <NavLink
      style={{
        viewTransitionName: 'match-element',
      }}
      className={({ isActive, isPending }) =>
        classNames(
          'rounded-xl px-2 py-1',
          isActive ? 'bg-tg-accent-text pointer-events-none' : 'bg-tg-button',
          isPending && 'pointer-events-none',
        )
      }
      to={`../${dateApiString}`}
      viewTransition
    >
      <div className="flex flex-col items-center text-center">
        <div className="relative">
          <div className="absolute right-0 left-0">{dayOfWeek}</div>
          <div className="invisible">Mon</div>
        </div>
        <div className="bg-tg-bg flex h-8 w-8 items-center justify-center rounded-full">
          <p>{dayOfMonth}</p>
        </div>
      </div>
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
      <div className="flex justify-center gap-2">
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
      <AddHabitButton />
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
