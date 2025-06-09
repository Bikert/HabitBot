DELETE FROM habits
WHERE title IN (
                'Норма шагов',
                'Зарядка',
                'Соблюдение КБЖУ',
                'Сон до 23:00',
                'Тренировка'
    )
  AND isDefault = 1;