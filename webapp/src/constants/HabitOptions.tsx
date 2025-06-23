export const daysOfWeek = Object.freeze(['ПН', 'ВТ', 'СР', 'ЧТ', 'ПТ', 'СБ', 'ВС'] as const)
export type DayOfWeek = (typeof daysOfWeek)[number]
export const repeatTypes = ['daily', 'weekly'] as const
export type RepeatType = (typeof repeatTypes)[number]
export const colors = ['#b2f2bb', '#c77dff', '#ffa94d', '#63e6be', '#fa5252', '#fff3bf', '#868e96', '#f06595'] as const
export type HabitColor = (typeof colors)[number]
