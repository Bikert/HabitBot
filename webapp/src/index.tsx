import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App'
import { createBrowserRouter, redirect, replace, RouterProvider } from 'react-router'
import EditHabitPage, { editHabitLoader } from './components/EditHabitPage'
import { ConfigPage } from './components/ConfigPage'
import { ErrorBoundary } from './components/ErrorBoundary'
import { getCurrentDateApiString, isValidDateString } from './utils/date'
import { DayView, dayViewLoader } from './components/DayView'
import { delay } from './utils/delay'
import { queryClient } from './api/queryClient'

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
          loader: async () => replace('/day' + window.location.hash),
        },
        {
          path: 'habit/:id?',
          loader: editHabitLoader,
          Component: EditHabitPage,
        },
        {
          path: 'config',
          Component: ConfigPage,
        },
        {
          path: 'day',
          HydrateFallback: () => <div>Loading... (router hydrate fallback)</div>,
          ErrorBoundary: ErrorBoundary,
          children: [
            {
              index: true,
              loader: async () => {
                // HACK: give some time to router to understand transition is started, so isPending initialised
                await delay(1)
                return redirect('/day/' + getCurrentDateApiString())
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
export { habitsApi } from './api/habitsApi'
