import type { HabitsCreateHabitDto, HabitsHabitDto, HabitsUpdateHabitDto } from '@habit-bot/api-client'
import { type FormEvent, useActionState, useCallback, useState } from 'react'
import {
  colors,
  type DayOfWeek,
  daysOfWeek,
  daysOfWeekLabels,
  type HabitColor,
  type RepeatType,
  repeatTypes,
} from '../../constants/HabitOptions'
import { randomElement } from '../../utils/randomElement'
import { toast } from 'sonner'
import { EmojiInput } from '../EmojiInput'
import { TelegramWebApp } from '../../telegram'

interface EditHabitFormProps {
  existing?: HabitsHabitDto
  submit: (habit: HabitsCreateHabitDto) => void
  submitButtonLabel: string
  loading: boolean
}

export function EditHabitForm({ existing, submit, submitButtonLabel, loading }: EditHabitFormProps) {
  const [title, setTitle] = useState(existing?.name ?? '')
  const [description, setDescription] = useState(existing?.desc ?? '')
  const [emoji, setEmoji] = useState(existing?.icon ?? '⭐')
  const [color, setColor] = useState<HabitColor>(
    (existing?.color as HabitColor) ?? randomElement(colors as unknown as HabitColor[]),
  )
  const [repeatType, setRepeatType] = useState<RepeatType>(existing?.repeatType === 'weekly' ? 'weekly' : 'daily')
  const [selectedDays, setSelectedDays] = useState<DayOfWeek[]>(
    (existing?.daysOfWeek?.split(',').filter((d) => d) as DayOfWeek[]) ?? [],
  )

  const toggleDay = (day: DayOfWeek) => {
    setSelectedDays((prev) => (prev.includes(day) ? prev.filter((d) => d !== day) : [...prev, day]))
  }

  const handleSubmit = useCallback(
    (e: FormEvent) => {
      if (repeatType === 'weekly' && selectedDays.length === 0) {
        e.stopPropagation()
        e.preventDefault()
        toast.error('Please select at least one day')
      }
    },
    [repeatType, selectedDays],
  )

  const [, formAction, pending] = useActionState(
    async (previousState: Partial<HabitsCreateHabitDto>, formData: FormData) => {
      const title = formData.get('title') as string
      const description = (formData.get('description') ?? '') as string
      const data: HabitsCreateHabitDto = {
        icon: emoji,
        desc: description,
        name: title,
        daysOfWeek: selectedDays?.join(','),
        // TODO: fix expected format on BE
        // firstDate: getCurrentDateApiString(),
        firstDate: new Date().toISOString(),
        repeatType,
        color,
      }
      submit(data)
      return data
    },
    {} as Partial<HabitsCreateHabitDto>,
  )

  return (
    <form autoComplete="off" className="m-5 overflow-y-auto" onSubmit={handleSubmit} action={formAction}>
      <div className="bg-tg-secondary-bg mx-auto max-w-md rounded-3xl p-5 shadow-lg">
        <div className="mb-5 text-center">
          <div
            className={`${pending ? 'animate-spin [animation-duration:1000ms]' : ''} cursor-pointer text-5xl duration-75`}
          >
            <EmojiInput value={emoji} setValue={setEmoji} />
          </div>
          <h2 className="my-2">{title === '' ? 'Привычка' : title}</h2>
        </div>

        <input
          autoFocus
          type="text"
          name="title"
          placeholder="Name your new task"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
          className="bg-tg-bg outline-tg-accent-text mb-3 w-full rounded-xl p-3 text-base"
        />

        <input
          type="text"
          name="description"
          placeholder="Describe it"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="outline-tg-accent-text bg-tg-bg mb-3 w-full rounded-xl p-3 text-base"
        />

        <div className="mb-5">
          <b>Card Color</b>
          <div className="mt-2 flex justify-center-safe gap-2">
            {colors.map((c) => (
              <div
                key={c}
                onClick={() => {
                  setColor(c)
                  TelegramWebApp.HapticFeedback.selectionChanged()
                }}
                className={`border-tg-secondary-bg outline-tg-accent-text box-border h-7 w-7 cursor-pointer rounded-full ${color === c ? 'outline-4' : 'border-2'}`}
                style={{
                  backgroundColor: c,
                }}
              />
            ))}
          </div>
        </div>

        <div className="mb-5">
          <b>Repeat</b>

          <div className="mt-2 flex overflow-hidden rounded-xl select-none">
            {repeatTypes.map((type) => (
              <div
                key={type}
                onClick={() => {
                  setRepeatType(type)
                  TelegramWebApp.HapticFeedback.selectionChanged()
                }}
                className={`flex-1 cursor-pointer py-2.5 text-center font-semibold transition-colors ${repeatType === type ? 'bg-tg-button text-tg-button-text' : 'bg-tg-bg text-tg-link'}`}
              >
                {type.charAt(0).toUpperCase() + type.slice(1)}
              </div>
            ))}
          </div>

          <div
            className="mt-3 grid overflow-y-hidden duration-250"
            style={{
              gridTemplateRows: repeatType === 'weekly' ? '1fr' : '0fr',
            }}
          >
            <div className="flex justify-center gap-1 overflow-y-hidden">
              {daysOfWeek.map((day) => (
                <div
                  key={day}
                  onClick={() => {
                    toggleDay(day)
                    TelegramWebApp.HapticFeedback.selectionChanged()
                  }}
                  className={`cursor-pointer rounded-full px-2 py-0.5 font-medium select-none ${
                    selectedDays.includes(day)
                      ? 'bg-tg-button border-transparent text-white'
                      : 'bg-tg-bg border-gray-400 text-gray-600'
                  }`}
                >
                  {daysOfWeekLabels[day]}
                </div>
              ))}
            </div>
          </div>
        </div>

        <button
          disabled={loading}
          type="submit"
          className="bg-tg-button text-tg-button-text w-full cursor-pointer rounded-3xl py-3.5 text-lg font-bold transition-colors select-none disabled:cursor-not-allowed disabled:opacity-50"
        >
          {submitButtonLabel}
        </button>
      </div>
    </form>
  )
}
