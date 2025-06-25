import { type LoaderFunction, NavLink, replace, useLoaderData, useNavigate, useParams } from 'react-router'
import { delay } from '../utils/delay'
import { getCurrentDate, getRelativeDate, isValidDateString } from '../utils/date'
import classNames from 'classnames'
import type { HabitColor } from '../constants/HabitOptions'
import { TelegramWebApp } from '../telegram'

type Habit = {
  id: number
  title: string
  emoji: string
  description: string
  color: HabitColor
}

type HabitOnDate = {
  habit: Habit
  completed: boolean
}

type DayViewLoaderData = {
  habitsOnDate: HabitOnDate[]
}

const fakeHabitOnDate: HabitOnDate[] = [
  {
    habit: {
      id: 1,
      title: 'Read',
      description: '',
      emoji: '📚',
      color: '#63e6be',
    },
    completed: true,
  },
  {
    habit: {
      id: 2,
      title: 'Study',
      description: '',
      emoji: '📖',
      color: '#c77dff',
    },
    completed: false,
  },
  {
    habit: {
      id: 3,
      title: 'Mop the house',
      description: '',
      emoji: '🧹',
      color: '#f06595',
    },
    completed: false,
  },
]

export const dayDataLoader: LoaderFunction = async ({ params }) => {
  const date = params['date']
  if (!isValidDateString(date)) return replace('/day/' + getCurrentDate())
  await delay(500)
  return { habitsOnDate: structuredClone(fakeHabitOnDate) }

  // const response = await fetch('/api/habits/day/' + params['date'])
  // const habits = await response.json()
  // return {
  //   habits,
  // }
}

export const DayView = () => {
  const { habitsOnDate }: DayViewLoaderData = useLoaderData()
  const navigate = useNavigate()
  const params = useParams()
  const date = params['date']!

  async function toggleHabit(id: Habit['id']) {
    const existing = habitsOnDate.find((h) => h.habit.id === id)
    if (!existing) {
      return
    }
    const res = await fetch(`/api/habits/${id}`, {
      method: 'PATCH',
      headers: {
        'content-type': 'application/json',
        'x-telegram-init-data': TelegramWebApp.initData,
      },
      body: JSON.stringify({ completed: !existing.completed }),
    })
    if (!res.ok) {
      console.log('failed to toggle habit')
      const originalExisting = fakeHabitOnDate.find((h) => h.habit.id === id)
      if (!originalExisting) {
        return
      }
      originalExisting.completed = !existing.completed
    }
    navigate('.', { replace: true })
  }

  return (
    <>
      <div className="flex justify-center gap-2">
        <NavLink to={`../${getRelativeDate(date, -1)}`} className="select-none">
          {'<'}
        </NavLink>
        <div>Date: {date}</div>
        <NavLink to={`../${getRelativeDate(date, 1)}`} className="select-none">
          {'>'}
        </NavLink>
      </div>
      <ul className="m-4 flex list-none flex-col gap-2">
        {habitsOnDate.map(({ habit, completed }) => (
          <li
            key={habit.id}
            className={classNames('flex justify-between gap-2 rounded-xl p-3')}
            style={{ backgroundColor: `color-mix(in srgb, ${habit.color}, transparent 40%)` }}
            onClick={() => toggleHabit(habit.id)}
          >
            <div className="flex grow flex-row items-center">
              <div>{habit.emoji}</div>
              <div
                className={classNames(
                  'grow',
                  completed &&
                    'after:bg-tg-text relative box-border line-through after:absolute after:top-1/2 after:block after:h-[1px] after:w-full after:px-11 after:content-[""]',
                )}
              >
                {habit.title}
              </div>
            </div>
            <div className="flex items-center">
              <input
                type="checkbox"
                className="checked:bg-tg-button h-5 w-5 cursor-pointer appearance-none rounded-full border-3"
                onChange={() => toggleHabit(habit.id)}
                checked={completed}
              />
            </div>
          </li>
        ))}
      </ul>
    </>
  )
}
