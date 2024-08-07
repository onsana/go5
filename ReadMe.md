# Readme 

## Lesson 1

## Lesson 2

1. Розробити програму «Зоопарк».
Завдання: 5 чи більше звірів повтікали, наглядач повинен їх зібрати. Кожну сутність (наглядач, звір, клітка тощо) представляти окремою структурою (zookeeper, animal, cage). Користуємось ембдінгом і методами.
2. Зареєструватися на https://github.com/.
3. Створити публічний репозиторій github і додати в нього виконану домашню роботу. Посилання на репозиторій зберегти у рішення домашньої роботи.

## Lesson 3

Розробити гру-текстовий квест «Новий світ».

Ваш персонаж прокидається в невідомому місці з деякими речами. Він нічого не памʼятає. У нього є можливість піти одним з кількох шляхів (усі перелічені сутності — структури). Ситуація розвивається залежно від обраного рішення.

Ігровий режим: текстом пишеться ситуація і пропонуються текстові варіанти, які може обрати гравець. Гравець пише один з варіантів і читає, як у цьому випадку розвивається ситуація.

Можливий сценарій:
Стівен прокинувся біля входу в печеру.
Він лише памʼятає своє імʼя.
Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж.
У печері темно, тому Стівен іде стежкою, яка веде від печери в ліс.
У лісі Стівен натикається на мертве тіло дивної тварини. Він обирає нічого з цим не робити й іти далі.
Через деякий час Стівен приходить до безлюдного табору. Він вже втомлений і вирішує відпочити, а не йти далі.
У найближчому наметі він знаходить сейф з кодовим замком з двох чисел.
Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає.
Стівен непритомніє. А все могло бути зовсім інакше.

Треба використати Scan, type struct, if/else, switch/case, for loop.

https://pkg.go.dev/fmt#Scanf

## Lesson 4

1. Пошук для текстового редактора. Створити slice string з текстом, який користувач ввів у текстовий редактор. Написати функцію, яка приймає на вхід рядок для пошуку та знаходить у текстовому редакторі всі рядки, які містять рядок пошуку.  Використовуючи цю функцію, додати можливість пошуку тексту в текстовому редакторі та вивести на екран усі відповідні результати. Розширена задача: ініціалізувати текс в редакторі не через код програми, а зчитавши рядки тексту з файлу (приклади читання файлу https://www.tutorialspoint.com/golang-program-to-read-the-content-of-a-file-line-by-line)

2. Унікальні структури. Створити тип структура, що містить одне поле (наприклад `ID`). Написати функцію, яка на вхід приймає слайс з елементами створеного типу, а повертає слайс того ж типу лише з унікальними значеннями (структури з дублікатами значення поля відкидаються). Результ функції має бути відсортований у порядку зростання значень поля структури. Додаткові умови: Не використовувати бібліотеки для пошуку унікальних значень. Використати можливості стандартної бібліотеки `sort` для сортування. Приклад: [{3}, {2}, {1}, {2}] -> [{1}, {2}, {3}]

## Lesson 5

Імплементувати пошук для текстового редактора (аналогічно до завдання в HW4) використовуючи індекс слів у мапі. Тобто, для текстового редактора реалізувати методи "проіндексувати текст по словам", та "пошук усіх рядків за словом".

## Lesson 6

Створити на Golang інтерфейс «Публічний транспорт», який має методи «Приймати пасажирів» та «Висаджувати пасажирів», і реалізувати його для типів «Автобус», «Потяг», «Літак».

Створити тип «Маршрут», який містить список транспортних засобів, які необхідні для проходження по заданому маршруту. Тип «Маршрут» має мати методи «Додавати транспортний засіб до маршруту» та «Показувати список транспортних засобів на маршруті». Тепер цей маршрут мусить пройти ваш подорожувальник («Пасажир») із виводом його подорожі на екран.

Файли різних груп об‘єктів зберігати в різних пакетах.

## Lesson 7

Написати дві програми:

1. Яка створює 3 горутини. Перша горутина генерує випадкові числа й надсилає їх через канал у другу горутину. Друга горутина отримує випадкові числа та знаходить їх середнє значення, після чого надсилає його в третю горутину через канал. Третя горутина виводить середнє значення на екран.

2. Яка створює 2 горутини. Перша горутина генерує випадкові числа в заданому діапазоні й надсилає їх через канал у другу горутину. Друга горутина отримує випадкові числа і знаходить найбільше й найменше число, після чого надсилає їх назад у першу горутину через канал. Перша горутина виводить найбільше й найменше числа на екран.

## Lesson 8

Створити програму для симуляції групи людей, які одночасно грають в ігри на великому екрані (моделюємо гру схожу на kahoot https://www.youtube.com/watch?v=az1xm2Ij7rA). Програма має використовувати горутину-генератор, який кожні 10 секунд генерує новий ігровий раунд (питання та варіанти відповідей) та відправляє його до горутин-гравців через канал. Гравці отримують новий ігровий раунд та вводять свої відповіді через окремий канал. Далі горутина-лічильник перевіряє правильність відповідей та повертає результат (кількість відповідей по варіантах та/або загальний результат гри по гравцях) у головну горутину через окремий канал, яка виводить результат раунду на екран. Якщо користувач перериває програму, то програма має коректно завершувати роботу з використанням контексту.

## Lesson 9

Створити вебсервер для перегляду інформації щодо класу школи.

Користувач повинен мати можливість отримувати загальну інформацію про клас (список учнів, назва класу).

Додаткові вимоги:

• інформація про учнів має зберігатися в оперативній пам'яті та бути доступною під час кожного запиту;

• отримання інформації про учня (наприклад, середній бал по предметах) має здійснюватись методом GET на адресі "/student/{id}", де {id} — унікальний ідентифікатор учня;

• дані можна отримати, лише якщо користувач є вчителем у цьому класі.

## Lesson 10

Створити сервер з REST API для перегляду списку cправ.

Користувач повинен мати можливості:

* переглядати список завдань

* додати нове завдання

* змінити існуюче завдання (наприклад, відмітити виконаним)

* видалити завдання

Додаткові вимоги:

* список завдань має бути збережений в оперативній пам'яті та бути доступним під час кожного запиту

* сервер має відповідати та приймати дані у форматі JSON

* можна використати стандартну бібліотеку net/http або спробувати популярні бібліотеки/фреймворки для web (echo, chi, gorilla, etc.)

## Lesson 12

1. Виконати пошук телефонних номерів у файлі з даними контактів. Задача: створити регулярний вислів, який можна використовувати для знаходження телефонних номерів, записаних у різних форматах. Наприклад, ви можете почати з використання вислову, який знаходить номери телефонів, що складаються з 10 цифр, а потім розширити його, додавши підтримку різних форматів, наприклад, номери з круглими дужками, пробілами та дефісами.

2. Реалізувати пошук слів із певним шаблоном у текстовому файлі. Задача: створити регулярний вислів, який можна використовувати для знаходження слів, які відповідають певному шаблону. Наприклад, вислів, який знаходить слова, що починаються на голосні літери та закінчуються на приголосні, або слова, що складаються з двох однакових букв, розділених будь-яким символом. Завдання творче.

## Lesson 13

Консольний менеджер паролів

Написати консольну програму для зберігання паролів (спрощений аналог утиліти pass в UNIX). Шифрування паролів в цій роботі не реалізуємо.

Функціонал:

* вивести назви збережених паролів

* зберегти пароль за назвою (введення паролю у через fmt.Scan)

* дістати збережений пароль

Додаткові умови:

* використовуємо tracer bullet development, тобто пишемо ітеративно

* зберігати стан у файлі (щоб паролі можна було дивитися між запусками)

* використати рекомендовану структуру пакетів (cmd, internal, …)