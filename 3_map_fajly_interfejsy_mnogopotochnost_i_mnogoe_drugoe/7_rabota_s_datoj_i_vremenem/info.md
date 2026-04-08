# Работа с временем в Go

Модуль `time` стандартной библиотеки предоставляет инструменты для работы с датой и временем, включая создание объектов времени, их форматирование и преобразование.

## type Time

### Создание структуры Time

Структура `Time` в Go используется для представления конкретных даты и времени. Мы можем создавать экземпляры этой структуры с помощью различных функций из пакета `time`.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Получаем текущее время
	now := time.Now()

	// Создаем время с помощью конкретных значений
	currentTime := time.Date(
		2020,     // год
		time.May, // месяц
		15,       // день
		10,       // часы
		13,       // минуты
		12,       // секунды
		45,       // наносекунды
		time.UTC, // временная зона
	)

	// Создаем время, используя секунды и наносекунды, прошедшие с начала эпохи Unix
	unixTime := time.Unix(
		150000, // секунды
		1,      // наносекунды
	)

	// Форматируем и выводим время в строковом виде
	fmt.Println(now.Format("02-01-2006 15:04:05"))         // 15-05-2020 09:58:16
	fmt.Println(currentTime.Format("02-01-2006 15:04:05")) // 15-05-2020 10:13:12
	fmt.Println(unixTime.Format("02-01-2006 15:04:05"))    // 02-01-1970 22:40:00
}

                  
```

В примере мы используем метод `Format` для преобразования времени в строку в определенном формате. Этот метод принимает строку-шаблон, и время будет преобразовано в строку, соответствующую этому шаблону.

### Конвертирование строк в структуру Time

Часто возникает необходимость преобразовать строковое представление даты и времени в объект `Time`. Для этого в Go есть две полезные функции: `Parse` и `ParseInLocation`.

Шаблон, используемый для парсинга строкового времени, всегда основан на следующей дате и времени: `"Mon Jan 2 15:04:05 MST 2006"` (понедельник, 2 января 2006 года, 15:04:05, время в часовом поясе MST). Это может показаться странным в первый раз, но со временем вы привыкнете к этому формату.

Вот как можно использовать эти функции для парсинга:

```go
// Функция Parse парсит строку в соответствии с заданным шаблоном
firstTime, err := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")
if err != nil {
	panic(err)
}

// LoadLocation находит временную зону в справочнике IANA
// https://www.iana.org/time-zones
loc, err := time.LoadLocation("Asia/Yekaterinburg")
if err != nil {
	panic(err)
}

// Функция ParseInLocation парсит строку в указанной временной зоне
secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:10pm", loc)
if err != nil {
	panic(err)
}

fmt.Println(firstTime.Format("02-01-2006 15:04:05"))  // 15-05-2020 17:45:00
fmt.Println(secondTime.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:10

                  
```

Как и раньше, мы используем метод `Format` для вывода времени в строковом виде. Важно, что аргумент `Format` должен быть в формате, понятном для Go.

Теперь вы можете легко работать с датами и временем, парсить их из строк и форматировать обратно в нужный вид. Это полезные инструменты для работы с временными данными в различных приложениях и системах.

### Методы структуры Time

#### Методы, возвращающие отдельные элементы структуры

Таких методов довольно много и в целом они не должны вызвать никаких проблем, для большей части этих методов мы сделаем короткие примеры:

```go
current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

// func (t Time) Date() (year int, month Month, day int)
fmt.Println(current.Date()) // 2020 May 15

// func (t Time) Year() int
fmt.Println(current.Year()) // 2020

// func (t Time) Month() Month
fmt.Println(current.Month()) // May

// func (t Time) Day() int
fmt.Println(current.Day()) // 15

// func (t Time) Clock() (hour, min, sec int)
fmt.Println(current.Clock()) // 17 45 12

// func (t Time) Hour() int
fmt.Println(current.Hour()) //17

// func (t Time) Minute() int
fmt.Println(current.Minute()) // 45

// func (t Time) Second() int
fmt.Println(current.Second()) // 12

// func (t Time) Unix() int64
fmt.Println(current.Unix()) // 1589546712

// func (t Time) Weekday() Weekday
fmt.Println(current.Weekday()) // Friday

// func (t Time) YearDay() int
fmt.Println(current.YearDay()) // 136
                  
```

Какие-то дополнительные комментарии к представленным методам не требуются — имена методов и возвращаемые значения говорят сами за себя.

#### Конвертирование структуры Time в строку

С методом Format мы уже знакомы.

```go
// func (t Time) Format(layout string) string
current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
fmt.Println(current.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:12
                  
```

#### Сравнение структур Time

```go
firstTime := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
secondTime := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

// func (t Time) After(u Time) bool
// true если позже
fmt.Println(firstTime.After(secondTime)) // true

// func (t Time) Before(u Time) bool
// true если раньше
fmt.Println(firstTime.Before(secondTime)) // false

// func (t Time) Equal(u Time) bool
// true если равны
fmt.Println(firstTime.Equal(secondTime)) // false
                  
```

#### Методы, изменяющие структуру Time

```go
now := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

// func (t Time) Add(d Duration) Time
// изменяет дату в соответствии с параметром - "продолжительностью"
future := now.Add(time.Hour * 12) // перемещаемся на 12 часов вперед

// func (t Time) AddDate(years int, months int, days int) Time
// изменяет дату в соответствии с параметрами - количеством лет, месяцев и дней
past := now.AddDate(-1, -2, -3) // перемещаемся на 1 год, два месяца и 3 дня назад

// func (t Time) Sub(u Time) Duration
// вычисляет время, прошедшее между двумя датами
fmt.Println(future.Sub(past)) // 10332h0m0s
                  
```

Обратите внимание, что в методах Add и AddDate могут использоваться и отрицательные значения, это позволяет не только «добавлять» время (что видно из названий методов), но и «отнимать» его.

## type Month

В предыдущем шаге остался один неразрешенный момент — указания на месяц. Это всего лишь объявленные на уровне модуля time константы, которые выглядят следующим образом:
 

```go
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
```

## type Duration

В предыдущем шаге мы увидели такой тип как Duration — продолжительность. Рассмотрим его подробнее. Внутри Duration представляет из себя int64, определяющий количество наносекунд, прошедших между двумя моментами времени.

Создается экземпляр типа Duration одной из следующих функций:

```go
now := time.Now()
past := now.AddDate(0, 0, -1)
future := now.AddDate(0, 0, 1)

// func Since(t Time) Duration
// вычисляет период между текущим моментом и заданным временем в прошлом
fmt.Println(time.Since(past).Round(time.Second)) // 24h0m0s

// func Until(t Time) Duration
// вычисляет период между текущим моментом и заданным временем в будущем
fmt.Println(time.Until(future).Round(time.Second)) // 24h0m0s

// func ParseDuration(s string) (Duration, error)
// преобразует строку в Duration с использованием аннотаций:
// "ns" - наносекунды,
// "us" - микросекунды,
// "ms" - миллисекунды,
// "s" - секунды,
// "m" - минуты,
// "h" - часы.
dur, err := time.ParseDuration("1h12m3s")
if err != nil {
	panic(err)
}
fmt.Println(dur.Round(time.Hour).Hours()) // 1
                  
```

Время — вещь текучая (и в общем, конечно, не вещь вовсе), поэтому не всегда нам удается получить то значение, какое мы ожидаем. Чтобы увидеть конкретный результат, который мы ожидали получить, мы в дополнение к рассматриваемой функции использовали метод Round, округляющий значение до ближайшего целого с заданной точностью.

У типа Duration помимо метода Round, который мы рассмотрели выше, есть ряд других методов, позволяющих вернуть часть значения: часы, минуты, секунды и пр.

```go
func (d Duration) Hours() float64
func (d Duration) Minutes() float64
func (d Duration) Seconds() float64
func (d Duration) Milliseconds() int64
func (d Duration) Microseconds() int64
func (d Duration) Nanoseconds() int64
                  
```

Завершая разговор об этом типе отметим, что модуль time содержит ряд констант типа Duration:

```go
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
                  
```

*Для целей решения задач на платформе Stepik необходимо учесть, что ряд методов типа Duration были включены в стандартную библиотеку в версиях 1.9 - 1.13 (Milliseconds, Microseconds, Round), таким образом их использовании в решении приведет к возникновению ошибки (по нашему мнению, таких задач в курсе нет).*