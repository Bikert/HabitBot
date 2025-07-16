import { useNavigate } from 'react-router'
import { useMutation } from '@tanstack/react-query'
import { habitsApi } from '../../api/habitsApi'
import { toast } from '../common/Toast'
import { useCallback, useState } from 'react'
import { Button, type ButtonProps } from '../common/Button'
import { DialogTrigger, Heading } from 'react-aria-components'
import { Modal } from '../common/Modal'
import { Dialog } from '../common/Dialog'
import { InfoIcon } from 'lucide-react'
import { chain } from 'react-aria'
import { useRegisterBlockerCallback } from '../../utils/useRegisterBlockerCallback'

type DisableHabitButtonProps = Omit<ButtonProps, 'onPress'> & { habitId: string }

export function DisableHabitButton({ habitId, ...buttonProps }: DisableHabitButtonProps) {
  const navigate = useNavigate()
  const { mutate } = useMutation({
    mutationFn: async () => {
      return habitsApi.apiHabitGroupIdDisablePost({
        groupId: habitId,
      })
    },
    onSuccess: (habitResponse) => {
      toast.success(habitResponse.message ?? 'Habit disabled')
      navigate(-1)
    },
    onError: (error) => {
      console.error(error)
      toast.error('Failed to save habit', { description: error.message })
    },
  })
  const [dialogOpened, setDialogOpened] = useState(false)

  useRegisterBlockerCallback({
    blockerCallback: useCallback(() => setDialogOpened(false), []),
    isBlocked: dialogOpened,
  })

  return (
    <>
      <Button
        {...buttonProps}
        type={buttonProps.type ?? 'button'}
        color={buttonProps.color ?? 'destructive'}
        onPress={() => setDialogOpened(true)}
      >
        {buttonProps.children ?? 'Disable'}
      </Button>
      <DialogTrigger isOpen={dialogOpened} onOpenChange={setDialogOpened}>
        <Modal>
          <Dialog>
            {({ close }) => (
              <>
                <Heading slot="title" className="my-0 text-xl leading-6 font-semibold text-on-surface">
                  Disable the habit?
                </Heading>
                <div className="absolute top-6 right-6 h-6 w-6 stroke-2 text-on-surface">
                  <InfoIcon aria-hidden />
                </div>
                <div className="mt-4 flex justify-between gap-2">
                  <Button variant="tonal" onPress={close}>
                    Cancel
                  </Button>
                  <div className="flex justify-end gap-2">
                    <Button autoFocus color="destructive" onPress={chain(mutate, close)}>
                      Disable
                    </Button>
                  </div>
                </div>
              </>
            )}
          </Dialog>
        </Modal>
      </DialogTrigger>
    </>
  )
}
