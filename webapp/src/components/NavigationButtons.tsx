import { useViewportHeight } from '../utils/useViewportHeight'
import classNames from 'classnames'
import { NavLink, To } from 'react-router'
import type { PropsWithChildren } from 'react'
import { PATHS } from '../constants/paths'

function NavigationButton({ to, children }: PropsWithChildren<{ to: To }>) {
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

export function NavigationButtons() {
  const viewportHeight = useViewportHeight()
  const navFixed = viewportHeight && viewportHeight > 500
  const buttons = (
    <>
      <NavigationButton to={PATHS.editHabit()}>
        <span className="material-icons">add</span>
      </NavigationButton>
      <NavigationButton to={PATHS.habitsList}>
        <span className="material-icons">checklist</span>
      </NavigationButton>
      <NavigationButton to={PATHS.day()}>
        <span className="material-icons">today</span>
      </NavigationButton>
      <NavigationButton to={PATHS.settings}>
        <span className="material-icons">tune</span>
      </NavigationButton>
    </>
  )
  return (
    <>
      {/* TODO: fix the buttons positioning */}
      {/* fake to add some space under the fixed element. */}
      <div className={classNames('rounded-t-4xl py-4 select-none', 'invisible', navFixed ? 'static flex' : 'hidden')}>
        <NavigationButton to="" />
      </div>
      <div
        className={classNames(
          'bg-tg-secondary-bg right-0 left-0 flex w-full justify-center gap-3 rounded-t-4xl py-4 select-none',
          navFixed ? 'bottom-tg-content-safe-bottom fixed' : 'static',
        )}
      >
        {buttons}
      </div>
    </>
  )
}
