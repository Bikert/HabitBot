import { BodyMeasurementsForm } from './bodyMetrices/BodyMeasurementsForm'

export const BodyMeasurementsPage = () => {
  return (
    <div className="mx-auto max-w-md p-4">
      <h1 className="mb-6 text-2xl font-bold text-tg-text">Body Measurements</h1>
      <BodyMeasurementsForm />
    </div>
  )
}
