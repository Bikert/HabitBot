import React from 'react'
import ReactDOM from 'react-dom/client'
import '@material-design-icons/font'
import './index.css'
import App from './App'
import { createBrowserRouter, replace, RouterProvider } from 'react-router'
import EditHabitPage, { editHabitLoader } from './components/EditHabitPage'
import { ConfigPage } from './components/ConfigPage'
import { ErrorBoundary } from './components/ErrorBoundary'
import { getCurrentDateApiString } from './utils/date'
import { DayView, dayViewLoader } from './components/DayView'
import { delay } from './utils/delay'
import { PATHS } from './constants/paths'
import { HabitsListPage } from './components/HabitsListPage'
import { BodyMeasurementsPage } from './components/BodyMeasurements/BodyMeasurementsPage'
import { TelegramWebApp } from './telegram'

if (!sessionStorage['initialLocation']) {
  sessionStorage['initialLocation'] = window.location.href
}

window.document.documentElement.dataset['theme'] = TelegramWebApp.colorScheme
TelegramWebApp.onEvent('themeChanged', () => {
  window.document.documentElement.dataset['theme'] = TelegramWebApp.colorScheme
})

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
          path: PATHS.bodyMeasurements,
          Component: BodyMeasurementsPage,
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
