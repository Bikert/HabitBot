import './App.css';
import React, {useEffect, useState} from 'react';
import WebApp from '@twa-dev/sdk';
import HabitForm from "./HabitForm";

function App() {
  const [user, setUser] = useState();

  useEffect(() => {
    // Показываем кнопку Telegram, если она есть
    WebApp.ready();
    WebApp.MainButton.setText('Отправить');
    WebApp.MainButton.onClick(() => {
      WebApp.sendData(JSON.stringify({ message: 'Привет от WebApp!' }));
    });
    WebApp.MainButton.show();
    const tg = window.Telegram?.WebApp;
    // Получаем initDataUnsafe, где лежит user
    const user = tg.initDataUnsafe?.user;

    if (user) {
      setUser(user)
      console.log("Telegram User Info:", user);
    } else {
      console.warn("No user info found in initDataUnsafe");
    }

  }, []);


  return (
      <div>
        {user ? user.id : ""}
        <HabitForm />
    </div>

  );
}

export default App;
