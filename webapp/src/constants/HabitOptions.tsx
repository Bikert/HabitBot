export const daysOfWeek = Object.freeze(['mon', 'tue', 'wed', 'thu', 'fri', 'sat', 'sun'] as const)
export type DayOfWeek = (typeof daysOfWeek)[number]
export const daysOfWeekLabels = {
  mon: 'ПН',
  tue: 'ВТ',
  wed: 'СР',
  thu: 'ЧТ',
  fri: 'ПТ',
  sat: 'СБ',
  sun: 'ВС',
}
export const repeatTypes = ['daily', 'weekly'] as const
export type RepeatType = (typeof repeatTypes)[number]
export const colors = ['#b2f2bb', '#c77dff', '#ffa94d', '#63e6be', '#fa5252', '#fff3bf', '#868e96', '#f06595'] as const
export type HabitColor = (typeof colors)[number]
