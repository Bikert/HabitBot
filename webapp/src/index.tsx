import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App'
import { createBrowserRouter, redirect, replace, RouterProvider } from 'react-router'
import HabitPage from './components/HabitPage'
import { ConfigPage } from './components/ConfigPage'
import { ErrorBoundary } from './components/ErrorBoundary'
import { getCurrentDate } from './utils/date'
import { dayDataLoader, DayView } from './components/DayView'

if (!sessionStorage['initialLocation']) {
  sessionStorage['initialLocation'] = window.location.href
}

const router = createBrowserRouter([
  {
    path: '/',
    Component: App,
    children: [
      {
        index: true,
        loader: async () => replace('/habit' + window.location.hash),
      },
      {
        path: 'habit/:id?',
        Component: HabitPage,
      },
      {
        path: 'config',
        Component: ConfigPage,
      },
      {
        path: 'day',
        HydrateFallback: () => <div>Loading...</div>,
        ErrorBoundary: ErrorBoundary,
        children: [
          {
            index: true,
            loader: async () => redirect('/day/' + getCurrentDate()),
          },
          {
            path: ':date',
            loader: dayDataLoader,
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
])

const root = ReactDOM.createRoot(document.getElementById('root')!)
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
