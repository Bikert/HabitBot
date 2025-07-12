function noop() {}

export function throttle<TV, T extends TV[]>(fn: (...args: T) => void, delay: number, triggerImmediately = true) {
  let timer = 0
  let scheduledAction = noop
  const actionAndClear = () => {
    scheduledAction()
    timer = 0
    scheduledAction = noop
  }

  return function (...args: T) {
    const action = () => fn(...args)
    if (timer) {
      scheduledAction = action
      return
    }

    timer = window.setTimeout(actionAndClear, delay)

    if (triggerImmediately) {
      scheduledAction = noop
      action()
    } else {
      scheduledAction = action
    }
  }
}
