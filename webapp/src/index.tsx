import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App'
import { createBrowserRouter, replace, RouterProvider } from 'react-router'
import EditHabitPage, { editHabitLoader } from './components/EditHabitPage'
import { ConfigPage } from './components/ConfigPage'
import { ErrorBoundary } from './components/ErrorBoundary'
import { getCurrentDateApiString } from './utils/date'
import { DayView, dayViewLoader } from './components/DayView'
import { delay } from './utils/delay'
import '@material-design-icons/font'
import { PATHS } from './constants/paths'
import { HabitsListPage } from './components/HabitsListPage'

if (!sessionStorage['initialLocation']) {
  sessionStorage['initialLocation'] = window.location.href
}

const router = createBrowserRouter(
  [
    {
      path: '/',
      Component: App,
      children: [
        {
          index: true,
          loader: async () => replace(PATHS.day()),
        },
        {
          path: PATHS.editHabit(':id?'),
          loader: editHabitLoader,
          Component: EditHabitPage,
        },
        {
          path: PATHS.habitsList,
          Component: HabitsListPage,
        },
        {
          path: PATHS.settings,
          Component: ConfigPage,
        },
        {
          path: PATHS.day(),
          HydrateFallback: () => <div>Loading... (router hydrate fallback)</div>,
          ErrorBoundary: ErrorBoundary,
          children: [
            {
              index: true,
              loader: async () => {
                // HACK: give some time to router to understand transition is started, so isPending initialised
                await delay(1)
                return replace(PATHS.day(getCurrentDateApiString()))
              },
            },
            {
              path: ':date',
              loader: dayViewLoader,
              Component: DayView,
            },
          ],
        },
        {
          path: '*',
          loader: () => replace('/'),
        },
      ],
    },
  ],
  {},
)

const root = ReactDOM.createRoot(document.getElementById('root')!)
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
