import { type LoaderFunction, NavLink, replace, useLoaderData, useNavigate, useParams } from 'react-router'
import { getCurrentDate, getRelativeDate, isValidDateString } from '../utils/date'
import classNames from 'classnames'
import { TelegramWebApp } from '../telegram'
import { Configuration, DefaultConfig, HabitsApi, type HabitsHabitCompletionDto } from '@habit-bot/api-client'

type DayViewLoaderData = {
  habitsOnDate: HabitsHabitCompletionDto[]
}

const api = new HabitsApi(
  new Configuration({
    ...DefaultConfig,
    basePath: window.location.origin,
    headers: {
      'x-telegram-init-data': TelegramWebApp.initData,
    },
  }),
)

export const dayDataLoader: LoaderFunction = async ({ params }) => {
  const date = params['date']
  if (!isValidDateString(date)) return replace('/day/' + getCurrentDate())

  const habitCompletions = await api.apiHabitCompletionDateGet({ date: date! })
  console.log(habitCompletions)
  return { habitsOnDate: habitCompletions }
}

type VersionId = NonNullable<DayViewLoaderData['habitsOnDate'][number]['habit']>['versionId']
export const DayView = () => {
  const { habitsOnDate }: DayViewLoaderData = useLoaderData()
  const navigate = useNavigate()
  const params = useParams()
  const date = params['date']!

  console.log(habitsOnDate)
  async function toggleHabit(versionId: VersionId) {
    if (!versionId) {
      throw new Error('versionId is required')
    }
    const existing = habitsOnDate.find((h) => h.habit?.versionId === versionId)
    if (!existing) {
      throw new Error('habit not found')
    }
    await api.apiHabitVersionIdDatePatch({
      date: date,
      versionId: versionId,
      request: {
        completed: !existing.completed,
      },
    })
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
        {habitsOnDate.map(({ completed, habit }) => (
          <li
            key={habit.versionId}
            className={classNames('flex justify-between gap-2 rounded-xl p-3')}
            style={{ backgroundColor: `color-mix(in srgb, ${habit.color}, transparent 40%)` }}
            onClick={() => toggleHabit(habit.versionId)}
          >
            <div className="flex grow flex-row items-center justify-stretch">
              <div>{habit.icon}</div>
              <div
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
                onChange={() => toggleHabit(habit.versionId)}
                checked={completed}
              />
            </div>
          </li>
        ))}
      </ul>
    </>
  )
}
