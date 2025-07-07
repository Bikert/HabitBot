export function throttle<TV, T extends TV[]>(fn: (...args: T) => void, delay: number) {
  let timer = 0
  let delayedArgs = [] as unknown as T
  const callback = () => {
    fn(...delayedArgs)
  }

  return function (...args: T) {
    delayedArgs = args
    if (!timer) {
      timer = window.setTimeout(() => {
        callback()
        timer = 0
      }, delay)
    }
  }
}
