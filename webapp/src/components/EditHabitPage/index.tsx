import { BlockerFunction, type LoaderFunction, replace, useBlocker, useNavigate, useParams } from 'react-router'
import { delay } from '../../utils/delay'
import { queryClient } from '../../api/queryClient'
import { habitsApi } from '../../api/habitsApi'
import { queryOptions, useMutation, useSuspenseQuery } from '@tanstack/react-query'
import type { HabitsCreateHabitDto } from '@habit-bot/api-client'
import { toast } from 'sonner'
import { DialogTrigger, Heading } from 'react-aria-components'
import { EditHabitForm } from './EditHabitForm'
import { useCallback, useEffect, useRef, useState } from 'react'
import { InfoIcon } from 'lucide-react'
import { chain } from 'react-aria'
import { Modal } from '../common/Modal'
import { Button } from '../common/Button'
import { Dialog } from '../common/Dialog'

function habitQueryOptions(id: string) {
  return queryOptions({
    queryKey: ['habit', id],
    queryFn: () => {
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

function CreateHabit() {
  const navigate = useNavigate()
  const mutation = useMutation({
    mutationFn: async (habit: HabitsCreateHabitDto) => {
      return habitsApi.apiHabitCreatePost({
        request: habit,
      })
    },
    onSuccess: (response) => {
      toast.success(`Habit ${response.name} created`)
      navigate(-1)
    },
    onError: (error) => {
      console.error(error)
      toast.error('Failed to save habit', { description: error.message })
    },
  })

  return <EditHabitForm loading={mutation.isPending} submit={mutation.mutate} submitButtonLabel="Create" />
}

function EditHabit({ id }: { id: string }) {
  const { data: existing } = useSuspenseQuery(habitQueryOptions(id))
  const navigate = useNavigate()

  const submitDataRef = useRef<HabitsCreateHabitDto>(existing)
  const [dialogOpened, setDialogOpened] = useState(false)
  const onFormSubmit = useCallback((data: HabitsCreateHabitDto) => {
    submitDataRef.current = data
    setDialogOpened(true)
  }, [])

  const newVersionMutation = useMutation({
    mutationFn: async () => {
      return habitsApi.apiHabitGroupIdPost({
        groupId: id,
        request: submitDataRef.current,
      })
    },
    onSuccess: (habitResponse) => {
      toast.success(`Habit ${habitResponse.name} updated`)
      navigate(-1)
    },
    onError: (error) => {
      console.error(error)
      toast.error('Failed to save habit', { description: error.message })
    },
  })

  const updateVersionMutation = useMutation({
    mutationFn: async () => {
      return habitsApi.apiHabitGroupIdVersionIdPut({
        versionId: existing.versionId,
        groupId: id,
        request: submitDataRef.current,
      })
    },
    onSuccess: (habitResponse) => {
      toast.success(`Habit ${habitResponse.name} updated`)
      navigate(-1)
    },
    onError: (error) => {
      console.error(error)
      toast.error('Failed to save habit', { description: error.message })
    },
  })

  const blocker = useBlocker(
    useCallback<BlockerFunction>(({ historyAction }) => historyAction === 'POP' && dialogOpened, [dialogOpened]),
  )
  useEffect(() => {
    if (blocker.state === 'blocked') {
      setDialogOpened(false)
      blocker.reset()
    }
  }, [blocker])

  return (
    <>
      <DialogTrigger isOpen={dialogOpened} onOpenChange={setDialogOpened}>
        <Modal>
          <Dialog>
            {({ close }) => (
              <>
                <Heading slot="title" className="my-0 text-xl leading-6 font-semibold text-on-surface">
                  How to apply the changes?
                </Heading>
                <div className="absolute top-6 right-6 h-6 w-6 stroke-2 text-on-surface">
                  <InfoIcon aria-hidden />
                </div>
                <p className="mt-3 text-on-surface-variant">Change all: both future and past habits will be updated.</p>
                <p className="mt-3 text-on-surface-variant">Change only future: only future habits will be updated.</p>
                <div className="mt-4 flex justify-between gap-2">
                  <Button variant="destructive" onPress={close}>
                    Cancel
                  </Button>
                  <div className="flex justify-end gap-2">
                    <Button variant="secondary" onPress={chain(updateVersionMutation.mutate, close)}>
                      Change all
                    </Button>
                    <Button variant="primary" autoFocus onPress={chain(newVersionMutation.mutate, close)}>
                      Change only future
                    </Button>
                  </div>
                </div>
              </>
            )}
          </Dialog>
        </Modal>
      </DialogTrigger>
      <EditHabitForm
        loading={newVersionMutation.isPending || updateVersionMutation.isPending}
        submit={onFormSubmit}
        submitButtonLabel="Update"
        existing={existing}
      />
    </>
  )
}

export default function EditHabitPage() {
  const id = useParams()['id']
  return !id ? <CreateHabit /> : <EditHabit id={id} />
}
