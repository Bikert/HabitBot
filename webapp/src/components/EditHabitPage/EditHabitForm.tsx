import type { HabitsCreateHabitDto, HabitsHabitDto } from '@habit-bot/api-client'
import { type FormEvent, useActionState, useCallback, useState } from 'react'
import { colors, type DayOfWeek, daysOfWeek, daysOfWeekLabels, repeatTypes } from '../../constants/HabitOptions'
import { randomElement } from '../../utils/randomElement'
import { EmojiInput } from '../EmojiInput'
import { TelegramWebApp } from '../../telegram'
import { TextField } from '../common/TextField'
import { Form } from '../common/Form'
import { Button } from '../common/Button'
import { ColorSwatchPicker, ColorSwatchPickerItem } from '../common/ColorSwatchPicker'
import { Key, parseColor } from 'react-aria-components'
import { Label } from '../common/Field'
import { ToggleButtonGroup } from '../common/ToggleButtonGroup'
import { ToggleButton } from '../common/ToggleButton'
import { toast } from '../common/Toast'

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
  const [repeatType, setRepeatType] = useState(new Set<Key>([existing?.repeatType === 'weekly' ? 'weekly' : 'daily']))
  const [selectedDays, setSelectedDays] = useState(
    new Set<Key>((existing?.daysOfWeek?.split(',').filter((d) => d) as DayOfWeek[]) ?? []),
  )

  const handleSubmit = useCallback(
    (e: FormEvent) => {
      if (repeatType.has('weekly') && selectedDays.size === 0) {
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
        daysOfWeek: [...selectedDays].join(','),
        // TODO: fix expected format on BE
        // firstDate: getCurrentDateApiString(),
        firstDate: new Date().toISOString(),
        repeatType: repeatType.has('weekly') ? 'weekly' : 'daily',
        color: color.toString('hex'),
      }
      submit(data)
      return data
    },
    {} as Partial<HabitsCreateHabitDto>,
  )

  const showDaysPicker = repeatType.has('weekly')
  return (
    <>
      <div className="mx-auto max-w-md rounded-3xl bg-surface-container-low p-5 shadow-lg">
        <Form autoComplete="off" className="" onSubmit={handleSubmit} action={formAction}>
          <div className="flex flex-col text-center">
            <div
              className={`${pending ? 'animate-spin [animation-duration:1000ms]' : ''} cursor-pointer text-5xl duration-75`}
            >
              <EmojiInput value={emoji} setValue={setEmoji} />
            </div>
            <h2>{title === '' ? 'Habit' : title}</h2>
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
            <ToggleButtonGroup
              className="flex justify-stretch overflow-hidden select-none"
              selectionMode="single"
              disallowEmptySelection
              selectedKeys={repeatType}
              onSelectionChange={(keys) => {
                setRepeatType(keys)
                TelegramWebApp.HapticFeedback.selectionChanged()
              }}
            >
              {repeatTypes.map((type) => (
                <ToggleButton
                  id={type}
                  key={type}
                  size="md"
                  className="grow rounded-xs first:rounded-s-full last:rounded-e-full"
                >
                  {type.charAt(0).toUpperCase() + type.slice(1)}
                </ToggleButton>
              ))}
            </ToggleButtonGroup>

            <div
              className="grid overflow-y-hidden duration-250"
              style={{
                gridTemplateRows: showDaysPicker ? '1fr' : '0fr',
              }}
            >
              <div className="overflow-y-hidden">
                <ToggleButtonGroup
                  selectionMode="multiple"
                  selectedKeys={selectedDays}
                  onSelectionChange={setSelectedDays}
                  className="flex justify-center"
                >
                  {daysOfWeek.map((day) => (
                    <ToggleButton key={day} size="xs" id={day} isDisabled={!showDaysPicker}>
                      {daysOfWeekLabels[day]}
                    </ToggleButton>
                  ))}
                </ToggleButtonGroup>
              </div>
            </div>
          </div>

          <Button
            isPending={loading}
            isDisabled={loading}
            type="submit"
            size="lg"
            className="w-full cursor-pointer select-none disabled:cursor-not-allowed disabled:opacity-50"
          >
            {submitButtonLabel}
          </Button>
        </Form>
      </div>
    </>
  )
}
