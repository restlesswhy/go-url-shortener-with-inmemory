# Укорачиватель ссылок

### Требования

Укорачиватель ссылок
Необходимо реализовать сервис, который должен предоставлять API по созданию сокращённых ссылок следующего формата:
- Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться только одна сокращенная ссылка.
- Ссылка должна быть длинной 10 символов
- Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание)
Сервис должен быть написан на Go и принимать следующие запросы по gRPC:
1. Метод Create, который будет сохранять оригинальный URL в базе и возвращать сокращённый
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный URL
Решение должно быть предоставлено в «конечном виде», а именно: Сервис должен быть распространён в виде Docker-образа. Ожидается, что сервис позволяет использовать для хранения postgresql*. И in-memory хранилище, держащее данные в памяти сервиса (т.е. Redis или любое другое внешнее хранилище не подойдет). Какое хранилище использовать, указывается параметром при запуске сервиса. API должно быть описано в proto файле
Покрыть реализованный функционал Unit-тестами


### Для запуска приложения:

```
make run
```

### Замечание

Билдится только со второго раза так как контейнер с базой в первый раз долго развертывается и миграции не проходят, со второго раза и далее все успешно работает. 
Мой укорачиватель несколько иной нежели необходимо, не так понял задание. В моем варианте мы не выбираем базу изначально, а запускаем две одновременно. То есть если мы имеем урл локально, то мы его выдаем, чтобы зря не ходить в базу. В комментариях в usecase все описано. Данное решение подразумевает что запись локально должна хранится установленное время, но я не успел внедрить это.

# To do:

 - Временное нахождение в локальном хранилище.
 - Unit-тесты
 - Более надежный способ укорачивания ссылки (скорре всего при помощи добаления uuid к записи)
 - Параллельный поиск в базах данных
 - Поработать над обработкой ошибок и логгированием