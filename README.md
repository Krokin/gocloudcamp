## Вопросы
```1. Опишите самую интересную задачу в программировании, которую вам приходилось решать?```

Необходимо было сделать пагинацию в REST API с информацией из бд, интерес в том, что для оптимизации запросов, 
нужно хранить данные о предыдущих запросах конкретного пользователя, но у REST архитектуры ведь есть принцип как "Отсутствие сохранения состояния",
Поэтому этот вызов мне хорошо запомнился.

```2. Расскажите о своем самом большом факапе? Что вы предприняли для решения проблемы?```

Писали в команде проект на С, во время обучения, около месяца в сумме потратили, и в ночь перед сдачей,
решили окончательно проверить, прогнать тесты, пофиксить стиль, гитом тогда не пользовались по неопытности,
и в итоге случайно удалил весь проект со всеми тестами, готовыми функциями, в итоге пришлось очень быстро
по памяти восстанавливать наработки в старой версии, которая все-таки нашлась в телеграмме. Это была веселая ночь.
На следующих проектах, git стал встречать меня по утрам и провожать по вечерам. 

```3. Каковы ваши ожидания от участия в буткемпе?```

1. Улушить навыки программирования на Go и понимание архитектуры высоконагруженных сервисов. Улучшить понимание и расширить свой стек технологий. 
2. Поработать над реальными проектами вместе с менторами, чтобы получить фидбек по hard и soft скилам.
3. Получить опыт командной работы и практику на реальный задачах.
## Часть 1
Разработан модуль для обеспечения работы с плейлистом. Модуль обладает следующими возможностями:
    
    Play - начинает воспроизведение
    Pause - приостанавливает воспроизведение
    AddSong - добавляет в конец плейлиста песню
    Next воспроизвести след песню
    Prev воспроизвести предыдущую песню

Находится в каталоге Part_One

## Часть 2
#### API для музыкального плейлиста на базе протокола gRPC

Сервис позволяет управлять музыкальным плейлистом по API.
Может выполнять CRUD операции с песнями в плейлисте, а также воспроизводить, приостанавливать, переходить к следующему и предыдущему трекам.

Методы взаимодейстия (подробнее в /proto):

    Play - запускает композицию
	Pause - останавливает композицию
	Next - переходит на следующию композицию
	Prev - переходит на предыдущую композицию
	Create - добавляет композицию в плейлист
	Delete - удаляет композицию из плейлиста
    Update - обновляет информацию о композиции
	ReadSong - возвращает композицию выбранную/играющую в данный момент
	ReadPlaylist - возвращает все композиции плейлиста
	
Персистентность данных обеспечена за счет сохранения json файла при завершении программы и загрузка, если возможно, при запуске сервера.
#### Пример использования сервиса
    Клиентский код находиться в Part_Two/cmd/client_test

Запуск симуляции обращения к сервису клиента:

Клиент:

sqren

Логи сервера:

скрин


#### Запуск сервиса:
```git clone git@github.com:Krokin/gocloudcamp.git```

```docker-compose up -d```