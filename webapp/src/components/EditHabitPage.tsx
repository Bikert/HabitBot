import { type FormEvent, useActionState, useCallback, useState } from 'react'
import { TelegramWebApp } from '../telegram'
import {
  colors,
  type DayOfWeek,
  daysOfWeek,
  daysOfWeekLabels,
  type HabitColor,
  type RepeatType,
  repeatTypes,
} from '../constants/HabitOptions'
import { EmojiInput } from './EmojiInput'
import { type LoaderFunction, replace, useNavigate, useParams } from 'react-router'
import { delay } from '../utils/delay'
import { queryClient } from '../api/queryClient'
import { habitsApi } from '../api/habitsApi'
import { queryOptions, useMutation, useSuspenseQuery } from '@tanstack/react-query'
import type { HabitsCreateHabitDto, HabitsUpdateHabitDto } from '@habit-bot/api-client'
import { useEmulateSlowConnection } from '../stores/featureFlagsStores'
import { randomElement } from '../utils/randomElement'
import { toast } from 'sonner'

function habitQueryOptions(id?: string) {
  return queryOptions({
    queryKey: ['habit', id],
    queryFn: () => {
      if (!id) return Promise.resolve(null)
      return habitsApi.apiHabitGroupIdGet({ groupId: id })
    },
    staleTime: 10_000,
  })
}

export const editHabitLoader: LoaderFunction = async ({ params }) => {
  const { id } = params
  if (!id) {
    return await delay(1)
  }
  try {
    const habit = await queryClient.fetchQuery(habitQueryOptions(id))
    if (!habit) return replace('/habit')
  } catch (e) {
    console.error(e)
    return replace('/habit')
  }
}

export default function EditHabitPage() {
  const id = useParams()['id']
  const { data: existing } = useSuspenseQuery(habitQueryOptions(id))
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

  const navigate = useNavigate()
  const emulateSlowConnection = useEmulateSlowConnection((state) => state.active)
  const toggleDay = (day: DayOfWeek) => {
    setSelectedDays((prev) => (prev.includes(day) ? prev.filter((d) => d !== day) : [...prev, day]))
  }
  const editMutation = useMutation({
    mutationFn: async (habit: HabitsCreateHabitDto | HabitsUpdateHabitDto) => {
      if (emulateSlowConnection) {
        await delay(1500)
      }
      if ('id' in habit && habit.id) {
        return habitsApi.apiHabitGroupIdPost({
          groupId: habit.id,
          request: habit,
        })
      }
      return habitsApi.apiHabitCreatePost({
        request: habit,
      })
    },
    onSuccess: (habitResponse, habitRequest) => {
      if ('id' in habitRequest && habitRequest.id) {
        toast.success(`Habit ${habitResponse.name} updated`)
      } else {
        toast.success(`Habit ${habitResponse.name} created`)
      }
      navigate(-1)
    },
    onError: (error) => {
      console.error(error)
      toast.error('Failed to save habit', { description: error.message })
    },
  })

  const handleSubmit = useCallback(
    (e: FormEvent) => {
      if (repeatType === 'weekly' && selectedDays.length === 0) {
        e.preventDefault()
        TelegramWebApp.showAlert('Please select at least one day')
      }
    },
    [repeatType, selectedDays],
  )

  const [, formAction, pending] = useActionState(
    async (previousState: Partial<HabitsCreateHabitDto>, formData: FormData) => {
      const title = formData.get('title') as string
      const description = (formData.get('description') ?? '') as string
      const data: HabitsUpdateHabitDto | HabitsCreateHabitDto = {
        id: id,
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
      editMutation.mutate(data)
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
          disabled={editMutation.isPending}
          type="submit"
          className="bg-tg-button text-tg-button-text w-full cursor-pointer rounded-3xl py-3.5 text-lg font-bold transition-colors select-none disabled:cursor-not-allowed disabled:opacity-50"
        >
          {id ? 'Update' : 'Add'}
        </button>
      </div>
    </form>
  )
}
