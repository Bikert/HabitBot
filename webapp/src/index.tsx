import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App'
import { createBrowserRouter, redirect, replace, RouterProvider } from 'react-router'
import HabitForm from './components/HabitForm'
import { ConfigForm } from './components/ConfigForm'
import { ErrorBoundary } from './components/ErrorBoundary'
import { getCurrentDate } from './utils/date'
import { dayDataLoader, DayView } from './components/DayView'

const router = createBrowserRouter([
  {
    path: '/',
    Component: App,
    children: [
      {
        index: true,
        Component: HabitForm,
      },
      {
        path: 'habit/:id?',
        Component: HabitForm,
      },
      {
        path: 'config',
        Component: ConfigForm,
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
