function dateToDateSting(date: Date) {
  // do not use toISOString here - it formats in UTC+0 (Z) timezone instead of local, effectively adding off-by-one error
  return [
    date.getFullYear().toString(),
    (date.getMonth() + 1).toString().padStart(2, '0'),
    date.getDate().toString().padStart(2, '0'),
  ].join('-')
}

export function getCurrentDate() {
  const date = new Date()
  return dateToDateSting(date)
}

export function reformatDate(date: string) {
  return getRelativeDate(date, 0)
}

export function getRelativeDate(date: string, days: number) {
  if (!isValidDateString(date)) return date
  const [year, month, day] = date.split('-').map(Number) as [number, number, number]
  const newDate = new Date(year, month - 1, day + days)
  return dateToDateSting(newDate)
}

export function toDate(dateString: string) {
  if (!isValidDateString(dateString)) throw new Error('Invalid date string')
  const [year, month, day] = dateString.split('-').map(Number) as [number, number, number]
  return new Date(year, month - 1, day)
}

export function isValidDateString(date: unknown): boolean {
  if (typeof date !== 'string') return false
  const datePattern = /^\d{4}-\d{2}-\d{2}$/
  if (!datePattern.test(date)) return false

  const [year, month, day] = date.split('-').map(Number) as [number, number, number]
  const dateObj = new Date(year, month - 1, day)

  return dateObj.getFullYear() === year && dateObj.getMonth() === month - 1 && dateObj.getDate() === day
}
