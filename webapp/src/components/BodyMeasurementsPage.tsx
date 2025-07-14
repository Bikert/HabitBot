import React from 'react'
import { BodyMeasurementsForm } from './bodyMetrices/BodyMeasurementsForm'

export const BodyMeasurementsPage = () => {
  return (
    <div className="mx-auto max-w-md p-4">
      <h1 className="text-tg-text mb-6 text-2xl font-bold">Body Measurements</h1>
      <BodyMeasurementsForm />
    </div>
  )
}
