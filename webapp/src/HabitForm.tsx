import { useState, useEffect } from 'react'
import TelegramWebApp from '@twa-dev/sdk'

const daysOfWeek = Object.freeze(['ПН', 'ВТ', 'СР', 'ЧТ', 'ПТ', 'СБ', 'ВС'] as const)
type DayOfWeek = (typeof daysOfWeek)[number]
const colors = ['#b2f2bb', '#c77dff', '#ffa94d', '#63e6be', '#fa5252', '#fff3bf', '#868e96', '#f06595']

export default function HabitForm() {
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [color, setColor] = useState(colors[0])
  const [repeatType, setRepeatType] = useState('daily') // daily | weekly
  const [selectedDays, setSelectedDays] = useState<DayOfWeek[]>([])
  const [webAppHeight, setWebAppHeight] = useState(window.innerHeight)

  useEffect(() => {
    if (TelegramWebApp) {
      TelegramWebApp.expand()
      setWebAppHeight(TelegramWebApp.viewportHeight)

      const onResize = () => {
        setWebAppHeight(TelegramWebApp.viewportHeight)
      }
      TelegramWebApp.onEvent('viewportChanged', onResize)
      return () => {
        TelegramWebApp.offEvent('viewportChanged', onResize)
      }
    } else {
      const onResize = () => setWebAppHeight(window.innerHeight)
      window.addEventListener('resize', onResize)
      return () => window.removeEventListener('resize', onResize)
    }
  }, [])

  const toggleDay = (day: DayOfWeek) => {
    setSelectedDays((prev) => (prev.includes(day) ? prev.filter((d) => d !== day) : [...prev, day]))
  }

  const handleSubmit = (e) => {
    e.preventDefault()
    const data = { title, description, color, repeatType, selectedDays }
    console.log('Создана привычка:', data)
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.close()
    }
  }

  return (
    <form
      className="habit-form-component"
      onSubmit={handleSubmit}
      style={{
        height: webAppHeight,
      }}
    >
      <div className="habit-form">
        <div className="title">
          <div className="emoji">⭐</div>
          <h2>{title === '' ? 'Привычка' : title}</h2>
          {/* <small>Click to change the emoji</small> */}
        </div>

        {/* Название */}
        <input
          type="text"
          placeholder="Name your new task"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />

        {/* Описание */}
        <input
          type="text"
          placeholder="Describe it"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />

        {/* Цвет карточки */}
        <div style={{ marginBottom: 20 }}>
          <b>Card Color</b>
          <div style={{ marginTop: 8, display: 'flex', gap: 10 }}>
            {colors.map((c) => (
              <div
                key={c}
                onClick={() => setColor(c)}
                style={{
                  width: 28,
                  height: 28,
                  borderRadius: '50%',
                  backgroundColor: c,
                  border: c === color ? '3px solid #000' : '2px solid #ccc',
                  cursor: 'pointer',
                  boxSizing: 'border-box',
                }}
              />
            ))}
          </div>
        </div>

        {/* Repeat */}
        <div style={{ marginBottom: 20 }}>
          <b>Repeat</b>

          {/* Tabs */}
          <div
            style={{
              display: 'flex',
              marginTop: 8,
              borderRadius: 12,
              backgroundColor: '#e9f0ff',
              overflow: 'hidden',
              userSelect: 'none',
            }}
          >
            {['daily', 'weekly'].map((type) => (
              <div
                key={type}
                onClick={() => setRepeatType(type)}
                style={{
                  flex: 1,
                  padding: '10px 0',
                  textAlign: 'center',
                  cursor: 'pointer',
                  backgroundColor: repeatType === type ? '#ffb57d' : 'transparent',
                  color: repeatType === type ? '#fff' : '#666',
                  fontWeight: '600',
                  transition: 'background-color 0.3s',
                }}
              >
                {type.charAt(0).toUpperCase() + type.slice(1)}
              </div>
            ))}
          </div>

          {/* Если weekly, показываем дни */}
          {repeatType === 'weekly' && (
            <div style={{ marginTop: 12, display: 'flex', gap: 6, justifyContent: 'center' }}>
              {daysOfWeek.map((day) => (
                <div
                  key={day}
                  onClick={() => toggleDay(day)}
                  style={{
                    padding: '3px 8px',
                    borderRadius: 20,
                    border: selectedDays.includes(day) ? 'none' : '1px solid #bbb',
                    backgroundColor: selectedDays.includes(day) ? '#ffb57d' : '#fff',
                    color: selectedDays.includes(day) ? '#fff' : '#444',
                    cursor: 'pointer',
                    userSelect: 'none',
                    fontWeight: '500',
                  }}
                >
                  {day}
                </div>
              ))}
            </div>
          )}
        </div>

        {/* Кнопка */}
        <button
          type="submit"
          style={{
            width: '100%',
            padding: 14,
            borderRadius: 30,
            backgroundColor: '#ffb57d',
            color: '#fff',
            fontSize: 18,
            fontWeight: '700',
            border: 'none',
            boxShadow: '0 4px 8px rgba(255, 181, 125, 0.6)',
            cursor: 'pointer',
            userSelect: 'none',
            transition: 'background-color 0.3s',
          }}
          onMouseDown={(e) => (e.currentTarget.style.backgroundColor = '#e6985d')}
          onMouseUp={(e) => (e.currentTarget.style.backgroundColor = '#ffb57d')}
          onMouseLeave={(e) => (e.currentTarget.style.backgroundColor = '#ffb57d')}
        >
          Add
        </button>
      </div>
    </form>
  )
}
