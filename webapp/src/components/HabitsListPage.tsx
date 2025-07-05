import { useSuspenseQuery } from '@tanstack/react-query'
import { habitsApi } from '../api/habitsApi'
import classNames from 'classnames'
import type { HabitsHabitDto } from '@habit-bot/api-client'
import { NavLink } from 'react-router'

function HabitRow({ habit }: { habit: HabitsHabitDto }) {
  return (
    <li
      key={habit.versionId}
      className={classNames('flex justify-between gap-2 rounded-xl p-3 align-baseline')}
      style={{
        backgroundColor: `color-mix(in srgb, ${habit.color}, transparent 40%)`,
      }}
    >
      <div>{habit.icon}</div>
      <div className="flex grow flex-row items-center justify-stretch">
        <div className={classNames('grow text-center')}>{habit.name}</div>
      </div>
      <div className="flex items-center">
        <NavLink to={`/habit/${habit.id}`}>
          <span className="material-icons">edit</span>
        </NavLink>
      </div>
    </li>
  )
}
export function HabitsListPage() {
  const { data, error, isError, isPending, isFetching } = useSuspenseQuery({
    queryKey: ['habits-list'],
    queryFn: () => habitsApi.apiHabitGet(),
  })
  return (
    <div>
      {isPending && 'Loading...'}
      {isError && <pre>{JSON.stringify(error)}</pre>}
      {data && (
        <ul className={classNames('m-4 flex list-none flex-col gap-2', isFetching && '*:opacity-50')}>
          {data.map((habit) => (
            <HabitRow key={habit.id} habit={habit} />
          ))}
        </ul>
      )}
    </div>
  )
}
