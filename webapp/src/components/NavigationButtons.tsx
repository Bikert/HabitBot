import classNames from 'classnames'
import { NavLink, To } from 'react-router'
import type { PropsWithChildren } from 'react'
import { PATHS } from '../constants/paths'

export function NavigationButton({ to, children }: PropsWithChildren<{ to: To }>) {
  return (
    <NavLink
      to={to}
      className={({ isActive, isPending }) =>
        classNames(
          'flex h-14 w-14 items-center justify-center rounded-full text-2xl',
          isActive ? 'bg-tg-bg pointer-events-none' : 'bg-tg-secondary-bg',
          isPending && 'pointer-events-none animate-spin',
        )
      }
    >
      {children}
    </NavLink>
  )
}

export function AddHabitButton() {
  return (
    <div className="ring-tg-bg shadow-tg-bg absolute right-1/12 bottom-1/12 rounded-full shadow-xl ring-2 select-none">
      <NavigationButton to={PATHS.editHabit()}>
        <span className="material-icons">add</span>
      </NavigationButton>
    </div>
  )
}

export function NavigationButtons() {
  return (
    <div className="flex w-full justify-center gap-3 py-4 select-none">
      <NavigationButton to={PATHS.habitsList}>
        <span className="material-icons">checklist</span>
      </NavigationButton>
      <NavigationButton to={PATHS.day()}>
        <span className="material-icons">today</span>
      </NavigationButton>
      <NavigationButton to={PATHS.settings}>
        <span className="material-icons">tune</span>
      </NavigationButton>
    </div>
  )
}
