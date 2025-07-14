import { FormEvent, useCallback } from 'react'
import { bodyMetricApi } from '../../api/bodyMetricApi'
import { BodyMetricsBodyMetricDTO } from '@habit-bot/api-client'
import { useMutation } from '@tanstack/react-query'
import { Button } from '../common/Button'
import { useImmer } from 'use-immer'
import { Form } from '../common/Form'
import { NumberField } from '../common/NumberField'
import { toast } from '../common/Toast'

type BodyMetrics = Omit<BodyMetricsBodyMetricDTO, 'date'>

const MEASUREMENTS_FIELDS: { key: keyof BodyMetrics; label: string; placeholder?: string }[] = [
  { key: 'weight', label: 'Вес (кг)', placeholder: 'Введите вес в килограммах' },
  { key: 'bicepsLeft', label: 'Бицепс левый (см)', placeholder: 'Введите обхват левого бицепса в сантиметрах' },
  { key: 'bicepsRight', label: 'Бицепс правый (см)', placeholder: 'Введите обхват правого бицепса в сантиметрах' },
  { key: 'chest', label: 'Грудь (см)', placeholder: 'Введите обхват груди в сантиметрах' },
  { key: 'waist', label: 'Талия (см)', placeholder: 'Введите обхват талии в сантиметрах' },
  { key: 'belly', label: 'Живот (см)', placeholder: 'Введите обхват живота в сантиметрах' },
  { key: 'hips', label: 'Бёдра (см)', placeholder: 'Введите обхват бёдер в сантиметрах' },
  {
    key: 'thighMaxLeft',
    label: 'Бедро верх (левое) (см)',
    placeholder: 'Введите обхват верхней части левого бедра в сантиметрах',
  },
  {
    key: 'thighMaxRight',
    label: 'Бедро верх (правое) (см)',
    placeholder: 'Введите обхват верхней части правого бедра в сантиметрах',
  },
  {
    key: 'thighLowLeft',
    label: 'Бедро низ (левое) (см)',
    placeholder: 'Введите обхват нижней части левого бедра в сантиметрах',
  },
  {
    key: 'thighLowRight',
    label: 'Бедро низ (правое) (см)',
    placeholder: 'Введите обхват нижней части правого бедра в сантиметрах',
  },
]

export function BodyMeasurementsForm() {
  const [measurements, updateMeasurements] = useImmer<BodyMetrics>({})

  const mutation = useMutation({
    mutationFn: (data: BodyMetricsBodyMetricDTO) => bodyMetricApi.apiBodyMetricCreatePost({ metric: data }),
    onSuccess: () => {
      toast.success(`Успешно сохранено!`)
    },
    onError: (error) => {
      toast.error('Возникла ошибка ', { description: error.message })
    },
  })

  const handleSubmit = useCallback(
    (e: FormEvent) => {
      e.preventDefault()
      const result: BodyMetricsBodyMetricDTO = {
        ...measurements,
        date: new Date().toISOString(),
      }
      mutation.mutate(result)
    },
    [measurements, mutation],
  )

  return (
    <div className="mx-auto max-w-md rounded-3xl bg-surface-container-low p-5 shadow-lg">
      <Form className="" onSubmit={handleSubmit}>
        {MEASUREMENTS_FIELDS.map(({ key, label, placeholder }) => (
          <NumberField
            key={key}
            name={key}
            label={label}
            minValue={0}
            step={0.1}
            formatOptions={{
              useGrouping: false,
            }}
            description={placeholder}
            value={measurements[key]}
            onChange={(v) =>
              updateMeasurements((prev) => {
                prev[key] = v
              })
            }
          />
        ))}
        <Button type="submit">Сохранить</Button>
      </Form>
    </div>
  )
}
