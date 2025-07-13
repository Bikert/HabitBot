import type { HabitsCreateHabitDto, HabitsHabitDto } from '@habit-bot/api-client'
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
import { TextField } from '../common/TextField'
import { Form } from '../common/Form'
import { Button } from '../common/Button'
import { ColorSwatchPicker, ColorSwatchPickerItem } from '../common/ColorSwatchPicker'
import { parseColor } from 'react-aria-components'
import { Label } from '../common/Field'

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
  const [color, setColor] = useState(parseColor(existing?.color ?? randomElement(colors as unknown as string[])))
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
  console.log(color)

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
        color: color.toString('hex'),
      }
      submit(data)
      return data
    },
    {} as Partial<HabitsCreateHabitDto>,
  )

  return (
    <>
      <div className="mx-auto max-w-md rounded-3xl bg-surface-container-low p-5 shadow-lg">
        <Form autoComplete="off" className="overflow-y-auto" onSubmit={handleSubmit} action={formAction}>
          <div className="text-center">
            <div
              className={`${pending ? 'animate-spin [animation-duration:1000ms]' : ''} cursor-pointer text-5xl duration-75`}
            >
              <EmojiInput value={emoji} setValue={setEmoji} />
            </div>
            <h2 className="my-2">{title === '' ? 'Habit' : title}</h2>
          </div>

          <TextField
            autoFocus
            name="title"
            label="Title"
            description="Name your new task"
            value={title}
            onChange={setTitle}
            isRequired
          />

          <TextField
            type="text"
            name="description"
            label="Description"
            description="Describe it"
            value={description}
            onChange={setDescription}
          />

          <div className="flex flex-col gap-1">
            <Label>Card Color</Label>
            <div className="flex justify-center-safe">
              <ColorSwatchPicker
                value={color}
                onChange={(c) => {
                  setColor(c)
                  TelegramWebApp.HapticFeedback.selectionChanged()
                }}
              >
                {colors.map((c) => (
                  <ColorSwatchPickerItem key={c} color={c} />
                ))}
              </ColorSwatchPicker>
            </div>
          </div>

          <div>
            <Label>Repeat</Label>

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
                        ? 'border-transparent bg-tg-button text-white'
                        : 'border-gray-400 bg-tg-bg text-gray-600'
                    }`}
                  >
                    {daysOfWeekLabels[day]}
                  </div>
                ))}
              </div>
            </div>
          </div>

          <Button
            isPending={loading}
            isDisabled={loading}
            type="submit"
            className="w-full cursor-pointer rounded-3xl py-3.5 text-lg font-bold transition-colors select-none disabled:cursor-not-allowed disabled:opacity-50"
          >
            {submitButtonLabel}
          </Button>
        </Form>
      </div>
    </>
  )
}
