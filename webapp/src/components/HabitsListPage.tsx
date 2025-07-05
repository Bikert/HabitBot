import { useSuspenseQuery } from '@tanstack/react-query'
import { habitsApi } from '../api/habitsApi'
import classNames from 'classnames'
import type { HabitsHabitDto } from '@habit-bot/api-client'
import { NavLink } from 'react-router'
import { AddHabitButton } from './NavigationButtons'

function HabitRow({ habit }: { habit: HabitsHabitDto }) {
  return (
    <li
      className="rounded-xl p-3"
      style={{
        backgroundColor: `color-mix(in srgb, ${habit.color}, transparent 40%)`,
      }}
    >
      <NavLink to={`/habit/${habit.id}`} className={classNames('flex gap-2 rounded-xl')}>
        <div className={classNames('relative flex min-w-0 grow flex-row items-center justify-stretch')}>
          <span>{habit.icon}</span>
          <span className="min-w-0 overflow-x-clip text-nowrap text-ellipsis">{habit.name}</span>
        </div>
        <span className="material-icons">edit</span>
      </NavLink>
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
      <AddHabitButton />
    </div>
  )
}
