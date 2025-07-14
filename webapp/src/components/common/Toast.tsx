import { toast as sonnerToast } from 'sonner'
import { type ReactNode } from 'react'

interface ToastProps {
  id: string | number
  title: string
  description?: string
  button?: {
    label: string
    onClick?: () => void
  }
  icon?: ReactNode
}

type ToastParams = Omit<ToastProps, 'id' | 'title'>

export function toast(toast: Omit<ToastProps, 'id'>) {
  return sonnerToast.custom((id) => (
    <Toast id={id} title={toast.title} description={toast.description} button={toast.button} icon={toast.icon} />
  ))
}

toast.success = function (title: string, parameters?: ToastParams) {
  toast({
    title,
    icon: <span className="material-icons text-inverse-on-surface">check_circle</span>,
    ...parameters,
  })
}

toast.info = function (title: string, parameters?: ToastParams) {
  toast({
    title,
    icon: <span className="material-icons text-inverse-on-surface">info</span>,
    ...parameters,
  })
}

toast.warning = function (title: string, parameters?: ToastParams) {
  toast({
    title,
    icon: <span className="material-icons text-inverse-on-surface">warning</span>,
    ...parameters,
  })
}

toast.error = function (title: string, parameters?: ToastParams) {
  toast({
    title,
    icon: <span className="material-icons text-inverse-on-surface">error</span>,
    ...parameters,
  })
}

function Toast(props: ToastProps) {
  const { title, description, button, id, icon } = props
  return (
    <div className="flex items-center gap-1 rounded-xs bg-inverse-surface p-4 shadow-lg ring-1 shadow-shadow ring-black/5 select-none min-[600px]:w-(--width)">
      {!!icon && <div className="flex items-center">{icon}</div>}
      <div className="flex flex-1 grow items-center">
        <div className="w-full grow">
          <p className="w-full text-sm font-medium text-inverse-on-surface">{title}</p>
          {!!description && <p className="mt-1 text-sm text-gray-500">{description}</p>}
        </div>
      </div>
      {!!button && (
        <div className="ml-5 shrink-0 rounded-md text-sm font-medium text-indigo-600 hover:text-indigo-500 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-hidden">
          <button
            className="rounded px-3 py-1 text-sm font-semibold text-inverse-primary"
            onClick={() => {
              if (button?.onClick) {
                button.onClick()
              }
              sonnerToast.dismiss(id)
            }}
          >
            {button.label}
          </button>
        </div>
      )}
    </div>
  )
}
