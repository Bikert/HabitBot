function dateToDateSting(date: Date) {
  return `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`
}

export function getCurrentDate() {
  const date = new Date()
  return dateToDateSting(date)
}

export function getRelativeDate(date: string, days: number) {
  if (!isValidDateString(date)) return date
  const [year, month, day] = date.split('-').map(Number) as [number, number, number]
  const newDate = new Date(year, month - 1, day + days)
  return dateToDateSting(newDate)
}

export function isValidDateString(date: unknown): boolean {
  if (typeof date !== 'string') return false
  const datePattern = /^\d{4}-\d{1,2}-\d{1,2}$/
  if (!datePattern.test(date)) return false

  const [year, month, day] = date.split('-').map(Number) as [number, number, number]
  const dateObj = new Date(year, month - 1, day)

  return dateObj.getFullYear() === year && dateObj.getMonth() === month - 1 && dateObj.getDate() === day
}
