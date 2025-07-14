import { FormEvent, useCallback, useState } from 'react'
import { bodyMetricApi } from '../../api/bodyMetricApi'
import { BodyMetricsBodyMetricDTO } from '@habit-bot/api-client'
import { useMutation } from '@tanstack/react-query'
import { toast } from 'sonner'
import { NumberField } from '../common/TextField'
import { Button } from '../common/Button'

const MEASUREMENTS_MAP = {
  weight: {
    label: 'Вес (кг)',
    placeholder: 'Введите вес в килограммах',
  },
  bicepsLeft: {
    label: 'Бицепс левый (см)',
    placeholder: 'Введите обхват левого бицепса в сантиметрах',
  },
  bicepsRight: {
    label: 'Бицепс правый (см)',
    placeholder: 'Введите обхват правого бицепса в сантиметрах',
  },
  chest: {
    label: 'Грудь (см)',
    placeholder: 'Введите обхват груди в сантиметрах',
  },
  waist: {
    label: 'Талия (см)',
    placeholder: 'Введите обхват талии в сантиметрах',
  },
  belly: {
    label: 'Живот (см)',
    placeholder: 'Введите обхват живота в сантиметрах',
  },
  hips: {
    label: 'Бёдра (см)',
    placeholder: 'Введите обхват бёдер в сантиметрах',
  },
  thighMaxLeft: {
    label: 'Бедро верх (левое) (см)',
    placeholder: 'Введите обхват верхней части левого бедра в сантиметрах',
  },
  thighMaxRight: {
    label: 'Бедро верх (правое) (см)',
    placeholder: 'Введите обхват верхней части правого бедра в сантиметрах',
  },
  thighLowLeft: {
    label: 'Бедро низ (левое) (см)',
    placeholder: 'Введите обхват нижней части левого бедра в сантиметрах',
  },
  thighLowRight: {
    label: 'Бедро низ (правое) (см)',
    placeholder: 'Введите обхват нижней части правого бедра в сантиметрах',
  },
}

export function BodyMeasurementsForm() {
  const [measurements, setMeasurements] = useState<BodyMetricsBodyMetricDTO>({ date: new Date().toISOString() })

  const handleSubmit = useCallback((e: FormEvent) => {
    e.preventDefault()
    const result: BodyMetricsBodyMetricDTO = {
      ...measurements,
      date: new Date().toISOString(),
    }
    mutation.mutate(result)
  }, [])

  const mutation = useMutation({
    mutationFn: (data: BodyMetricsBodyMetricDTO) => bodyMetricApi.apiBodyMetricCreatePost({ metric: data }),
    onSuccess: () => {
      toast.success(`Успешно сохранено!`)
    },
    onError: (error) => {
      toast.error('Возникла ошибка ', { description: error.message })
    },
  })

  return (
    <form className="m-5 overflow-y-auto" onSubmit={handleSubmit}>
      <div className="mx-auto max-w-md rounded-3xl bg-tg-secondary-bg p-5 shadow-lg">
        {Object.entries(MEASUREMENTS_MAP).map(([key, { label, placeholder }]) => (
          <NumberField
            key={key}
            name={key}
            label={label}
            description={placeholder}
            value={
              measurements[key as keyof BodyMetricsBodyMetricDTO] !== undefined
                ? (measurements[key as keyof BodyMetricsBodyMetricDTO] as number)
                : 0
            }
            onChange={(v) => {
              setMeasurements((prev) => ({ ...prev, [key]: v }))
            }}
          />
        ))}
        <Button type="submit">Сохранить</Button>
      </div>
    </form>
  )
}
