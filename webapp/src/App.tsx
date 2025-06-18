import './App.css'
import { useEffect, useState } from 'react'
import WebApp from '@twa-dev/sdk'
import HabitForm from './HabitForm'

function App() {
  const [user, setUser] = useState<typeof WebApp.initDataUnsafe.user>()

  useEffect(() => {
    // Показываем кнопку Telegram, если она есть
    WebApp.ready()
    WebApp.MainButton.setText('Отправить')
    WebApp.MainButton.onClick(() => {
      WebApp.sendData(JSON.stringify({ message: 'Привет от WebApp!' }))
    })
    WebApp.MainButton.show()
    // Получаем initDataUnsafe, где лежит user
    const user = WebApp.initDataUnsafe?.user

    if (user) {
      setUser(user)
      console.log('Telegram User Info:', user)
    } else {
      console.warn('No user info found in initDataUnsafe')
    }
  }, [])

  return (
    <div>
      {user?.id}
      {WebApp.initData}
      <div>{window.location.href}</div>
      <HabitForm />
    </div>
  )
}

export default App
