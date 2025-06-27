import { useViewportHeight } from '../utils/useViewportHeight'
import classNames from 'classnames'
import { NavLink } from 'react-router'

export function NavigationButtons() {
  const viewportHeight = useViewportHeight()
  const navFixed = viewportHeight && viewportHeight > 500
  return (
    <div
      className={classNames(
        'right-0 left-0 mx-auto flex w-fit gap-3 select-none',
        navFixed ? 'bottom-tg-content-safe-bottom fixed mb-4' : 'static pb-4',
      )}
    >
      <NavLink
        to="/habit"
        className={({ isActive, isPending }) =>
          classNames(
            'text-tg-button-text ring-tg-bg flex h-14 w-14 items-center justify-center rounded-full text-2xl ring-2',
            isActive ? 'bg-tg-button-text' : 'bg-tg-button',
            isPending && 'pointer-events-none animate-spin',
          )
        }
      >
        ➕
      </NavLink>
      <NavLink
        to="/day"
        className={({ isActive, isPending }) =>
          classNames(
            'text-tg-button-text ring-tg-bg flex h-14 w-14 items-center justify-center rounded-full text-2xl ring-2',
            isActive ? 'bg-tg-button-text' : 'bg-tg-button',
            isPending && 'pointer-events-none animate-spin',
          )
        }
      >
        📋
      </NavLink>
      <NavLink
        to="/config"
        className={({ isActive, isPending }) =>
          classNames(
            'text-tg-button-text ring-tg-bg flex h-14 w-14 items-center justify-center rounded-full text-2xl ring-2',
            isActive ? 'bg-tg-button-text' : 'bg-tg-button',
            isPending && 'pointer-events-none animate-spin',
          )
        }
      >
        🔧
      </NavLink>
    </div>
  )
}
