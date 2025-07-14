import { FormEvent, useCallback, useState } from 'react'
import { bodyMetricApi } from '../../api/bodyMetricApi'
import { BodyMetricsBodyMetricDTO } from '@habit-bot/api-client'
import { useMutation } from '@tanstack/react-query'
import { toast } from 'sonner'

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
  const [weight, setWeight] = useState<number | ''>('')
  const [bicepsLeft, setBicepsLeft] = useState<number | ''>('')
  const [bicepsRight, setBicepsRight] = useState<number | ''>('')
  const [chest, setChest] = useState<number | ''>('')
  const [waist, setWaist] = useState<number | ''>('')
  const [belly, setBelly] = useState<number | ''>('')
  const [hips, setHips] = useState<number | ''>('')
  const [thighMaxLeft, setThighMaxLeft] = useState<number | ''>('')
  const [thighMaxRight, setThighMaxRight] = useState<number | ''>('')
  const [thighLowLeft, setThighLowLeft] = useState<number | ''>('')
  const [thighLowRight, setThighLowRight] = useState<number | ''>('')

  const handleSubmit = useCallback((e: FormEvent) => {
    e.preventDefault()
    const measurementsRaw = {
      date: new Date().toISOString(),
      weight,
      bicepsLeft,
      bicepsRight,
      chest,
      waist,
      belly,
      hips,
      thighMaxLeft,
      thighMaxRight,
      thighLowLeft,
      thighLowRight,
    }

    const filtered = Object.fromEntries(Object.entries(measurementsRaw).filter(([, value]) => value !== ''))

    const measurements: BodyMetricsBodyMetricDTO = {
      ...filtered,
      date: new Date().toISOString(),
    }
    mutation.mutate(measurements)
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
      <div className="bg-tg-secondary-bg mx-auto max-w-md rounded-3xl p-5 shadow-lg">
        <label>{MEASUREMENTS_MAP.weight.label}</label>
        <input
          type="number"
          name={MEASUREMENTS_MAP.weight.label}
          placeholder={MEASUREMENTS_MAP.weight.placeholder}
          value={weight}
          onChange={(e) => setWeight(e.target.value === '' ? '' : Number(e.target.value))}
        />

        <input
          type="number"
          name={MEASUREMENTS_MAP.bicepsLeft.label}
          placeholder={MEASUREMENTS_MAP.bicepsLeft.placeholder}
          value={bicepsLeft}
          onChange={(e) => setBicepsLeft(e.target.value === '' ? '' : Number(e.target.value))}
        />

        <input
          type="number"
          name={MEASUREMENTS_MAP.bicepsRight.label}
          placeholder={MEASUREMENTS_MAP.bicepsRight.placeholder}
          value={bicepsRight}
          onChange={(e) => setBicepsRight(e.target.value === '' ? '' : Number(e.target.value))}
        />

        <input
          type="number"
          name={MEASUREMENTS_MAP.chest.label}
          placeholder={MEASUREMENTS_MAP.chest.placeholder}
          value={chest}
          onChange={(e) => setChest(e.target.value === '' ? '' : Number(e.target.value))}
        />
        <input
          type="number"
          name={MEASUREMENTS_MAP.waist.label}
          placeholder={MEASUREMENTS_MAP.waist.placeholder}
          value={waist}
          onChange={(e) => setWaist(e.target.value === '' ? '' : Number(e.target.value))}
        />

        <input
          type="number"
          name={MEASUREMENTS_MAP.belly.label}
          placeholder={MEASUREMENTS_MAP.belly.placeholder}
          value={belly}
          onChange={(e) => setBelly(e.target.value === '' ? '' : Number(e.target.value))}
        />

        <input
          type="number"
          name={MEASUREMENTS_MAP.hips.label}
          placeholder={MEASUREMENTS_MAP.hips.placeholder}
          value={hips}
          onChange={(e) => setHips(e.target.value === '' ? '' : Number(e.target.value))}
        />

        <input
          type="number"
          name={MEASUREMENTS_MAP.thighMaxLeft.label}
          placeholder={MEASUREMENTS_MAP.thighMaxLeft.placeholder}
          value={thighMaxLeft}
          onChange={(e) => setThighMaxLeft(e.target.value === '' ? '' : Number(e.target.value))}
        />

        <input
          type="number"
          name={MEASUREMENTS_MAP.thighMaxRight.label}
          placeholder={MEASUREMENTS_MAP.thighMaxRight.placeholder}
          value={thighMaxRight}
          onChange={(e) => setThighMaxRight(e.target.value === '' ? '' : Number(e.target.value))}
        />
        <input
          type="number"
          name={MEASUREMENTS_MAP.thighLowLeft.label}
          placeholder={MEASUREMENTS_MAP.thighLowLeft.placeholder}
          value={thighLowLeft}
          onChange={(e) => setThighLowLeft(e.target.value === '' ? '' : Number(e.target.value))}
        />
        <input
          type="number"
          name={MEASUREMENTS_MAP.thighLowRight.label}
          placeholder={MEASUREMENTS_MAP.thighLowRight.placeholder}
          value={thighLowRight}
          onChange={(e) => setThighLowRight(e.target.value === '' ? '' : Number(e.target.value))}
        />

        <button type="submit">Сохранить</button>
      </div>
    </form>
  )
}
