import React, { useState } from 'react'

type Measurement = {
  value: string
  label: string
  unit: string
}

export const BodyMeasurementsPage = () => {
  const [measurements, setMeasurements] = useState<Record<string, string>>({})

  const measurementFields: Measurement[] = [
    { value: 'weight', label: 'Weight', unit: 'kg' },
    { value: 'chest', label: 'Chest', unit: 'cm' },
    { value: 'waist', label: 'Waist', unit: 'cm' },
    { value: 'hips', label: 'Hips', unit: 'cm' },
    { value: 'biceps', label: 'Biceps', unit: 'cm' },
    { value: 'thighs', label: 'Thighs', unit: 'cm' },
  ]

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // Handle submission logic here
    console.log('Measurements:', measurements)
  }

  const handleInputChange = (field: string, value: string) => {
    setMeasurements((prev) => ({
      ...prev,
      [field]: value,
    }))
  }

  return (
    <div className="mx-auto max-w-md p-4">
      <h1 className="mb-6 text-2xl font-bold text-tg-text">Body Measurements</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        {measurementFields.map(({ value, label, unit }) => (
          <div key={value} className="flex items-center gap-4">
            <label htmlFor={value} className="min-w-[80px] text-sm text-on-surface">
              {label}
            </label>
            <div className="flex grow items-center">
              <input
                type="number"
                id={value}
                step="0.1"
                value={measurements[value] || ''}
                onChange={(e) => handleInputChange(value, e.target.value)}
                className="grow rounded-lg bg-surface-container-low px-3 py-2 text-on-surface-variant focus:ring-2 focus:ring-secondary focus:outline-none"
                placeholder={`Enter ${label.toLowerCase()}`}
              />
              <span className="ml-2 min-w-[30px]">{unit}</span>
            </div>
          </div>
        ))}

        <button
          type="submit"
          className="mt-6 w-full rounded-lg bg-primary px-4 py-2 text-on-primary transition-opacity hover:opacity-90"
        >
          Save Measurements
        </button>
      </form>
    </div>
  )
}
