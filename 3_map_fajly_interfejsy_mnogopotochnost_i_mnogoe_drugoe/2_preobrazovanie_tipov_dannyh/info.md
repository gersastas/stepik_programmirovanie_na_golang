# Преобразование типов данных 

На своем пути разработчика вы будете часто встречаться с задачами, в которых нельзя без преобразования типов. Преобразование одного типа может положительно сказаться как на удобстве разработки, так и на производительности самого приложения. А может быть вам просто нужно вызвать функцию которая принимает только определенный тип данных.

Поговорим о **Golang** с точки зрения типов. Это статически типизированный язык, понятнее говоря, типы данных в Go (и других статически типизированных языках) связаны с переменной, а не с ее значением. Вы уже усвоили, что в переменной типа int могут храниться только целые числа. Но Go позволяет преобразовывать целочисленные типы, числа с плавающей запятой, строки, байты и так далее. 

# Приведение целочисленных типов в Go

Выбор целочисленного типа зависит от производительности, однако в некоторых ситуациях необходимо преобразовать один тип в другой. Например, Go часто генерирует числовые значения как `int`, что может не соответствовать типу вашего входного значения. Если вы используете `int32` и пытаетесь выполнить операцию с `int`, то вам нужно будет привести их типы для корректной работы.

В Go преобразования между целыми типами выполняются довольно легко. Для этого необходимо обернуть переменную в скобки и перед ней указать требуемый тип. Рассмотрим пример преобразования `int8` в `int32`:

```go
var index int8 = 15
var bigIndex int32
bigIndex = int32(index)

fmt.Println(bigIndex)         // 15
fmt.Printf("%T \n", bigIndex) // int32

                  
```

Аналогично можно конвертировать типы в другие, например:

```go
var a int32 = 22
var b uint64
b = uint64(a)

fmt.Println(b)         // 22
fmt.Printf("%T \n", b) // uint64

                  
```

**Примечание:** в Go можно использовать параметр `%T` в функции `Printf`, чтобы вывести тип переменной.

## Преобразование типов с меньшим количеством бит

Go позволяет преобразовывать типы с большим количеством бит в типы с меньшим количеством бит, например, из `int64` в `int8`:

```go
var big int64 = 64
var little int8

little = int8(big)
fmt.Println(little) // 64

                  
```

Однако важно помнить, что при преобразовании целых чисел может произойти потеря данных, если результат превышает допустимый диапазон для целевого типа. Например:

```go
var big int64 = 129
var little = int8(big)

fmt.Println(little)  // -127 (перенос данных)

                  
```

Для того чтобы узнать максимальные значения для различных целочисленных типов, можно воспользоваться константами из пакета `math`:

```go
fmt.Println(math.MaxInt8)   // 127
fmt.Println(math.MaxUint8)  // 255
fmt.Println(math.MaxInt16)  // 32767
fmt.Println(math.MaxUint16) // 65535
// ...
```

#  Приведение целых чисел и чисел с плавающей точкой

Преобразование целого числа в число с плавающей точкой ничем не отличается от преобразования целого числа. Можно использовать встроенную конверсию, обернув `float64()` или `float32()` вокруг целого числа: 

```go
var x int64 = 57
var y float64 = float64(x)
fmt.Print(y) // 57
                  
```

**Преобразование чисел с плавающей точкой в целые числа** 

Go может преобразовывать float в int. Но делает это с потерей точности. Синтаксис преобразования не меняется.

```go
var f float64 = 56.231
var i int = int(f)
fmt.Println(f) // 56.231
fmt.Println(i) // 56
                  
```

**Числа, конвертируемые с помощью деления** 

```css
a := 5 / 2
fmt.Println(a) // 2
                  
```

Если при делении используются числовые типы с плавающей точкой, тогда все остальные типы будут автоматически объявляться как числа с плавающей точкой:
 

```css
a := 5.0 / 2
fmt.Println(a) //2.5
```

#  Конвертация строк в байты/rune и обратно

Строка в Go это срез байтов, поэтому мы можем конвертировать байты в строку и наоборот:

```go
package main

import (
    "fmt"
)

func main() {
    a := "str"

    b := []byte(a)

    c := string(b)

    fmt.Println(a) // str

    fmt.Println(b) // [115 116 114] - побайтовый срез

    fmt.Println(c) // str
}
                  
```

Первая строка вывода - значение переменной "a", вторая - ее побайтовый срез, третья - значение переменной "c", являющейся результатом конвертации байтов в строку.

Тоже самое работает и со срезами типа **rune**:

```go
package main

import (
    "fmt"
)

func main() {
    a := "строка"

    b := []rune(a) // срез рун

    c := string(b)

    fmt.Println(a) // строка

    fmt.Println(b) // [1089 1090 1088 1086 1082 1072] - срез рун

    fmt.Println(c) // строка
}
```

# Конвертация в строки

Для начала рассмотрим конвертацию целых чисел в строки. **Golang** - язык со статической и строгой типизацией. Он не позволит вам сложить строку и число. Например напишем такой код:

```go
package main

import (
    "fmt"
)

func main() {
    user := "ученик"
    steps := 4

    fmt.Println("Поздравляю, " + user + "! Ты прошел " + steps + " шага по приведению типов.")
}
                  
```

Мы получим ошибку во время компиляции:

```
invalid operation: ("Поздравляю, " + user + "! Ты прошел ") + steps (mismatched types string and int)
```

При конвертации чисел в строки очень удобно использовать пакет `strconv`, он обладает методом `Itoa`, превращающим числовое значение (int) переменной в строковое (string).

В теории звучит очень сложно, на деле - просто. Рассмотрим на примере.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := strconv.Itoa(2020) // int -> string
	fmt.Printf("%T \n", a) // тип - string
	fmt.Println(a) // 2020
}
                  
```

А теперь вернемся к примеру выше. Чтобы исправить прошлую ошибку - надо привести steps к строке. 

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	user := "ученик"
	steps := 4

	fmt.Println("Поздравляю, " + user + "! Ты прошел " + strconv.Itoa(steps) + " шага по приведению типов.")
}

                  
```

Да, в этом примере не обязательно использовать `strconv`, мы могли бы просто перечислить в Print все переменные как аргументы функции, но могут возникнуть ситуации где никак не обойтись без конвертации. Мы лишь показали как это делать.

**Интересное дополнение**, метод `Itoa` это всего-лишь обертка для `FormatInt`: (кусок исходного кода пакета `strconv`)

```go
// Itoa is equivalent to FormatInt(int64(i), 10).
func Itoa(i int) string {
	return FormatInt(int64(i), 10)
}
                  
```

То-есть вызывая метод `Itoa` мы по сути вызываем `FormatInt` который принимает систему счисления в качестве 2 аргумента, но туда сразу передается - десятичная система счисления.

Но никто нам не мешает напрямую вызывать`FormatInt`, полезно если работаем с разными системами счисления:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// приставка '0x' означает что число в шестнадцатеричной системе счисления
	var a int64 = 0xB // 'B' в шестнадцатеричной это 11 в десятичной системе
	fmt.Println(strconv.FormatInt(a, 10)) // 11
	fmt.Println(strconv.FormatInt(a, 16)) // b
}

                  
```

## Конвертация целых беззнаковых чисел в строку


По аналогии с `FormatInt` есть такой метод как `FormatUint`, пример:

```go
var a uint64 = 10101
res := strconv.FormatUint(a, 10)
fmt.Println(res) // 10101
                  
```

###  Конвертация чисел с плавающей запятой в строку 

Для этого есть функция `FormatFloat`:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a float64 = 1.0123456789

	// 1 параметр - число для конвертации
	// fmt - форматирование
	// prec - точность (кол-во знаков после запятой)
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	fmt.Println(strconv.FormatFloat(a, 'f', 2, 64)) // 1.01

	// если мы хотим учесть все цифры после запятой, то можем в prec передать -1
	fmt.Println(strconv.FormatFloat(a, 'f', -1, 64)) // 1.0123456789

	// Возможные форматы fmt:
	// 'f' (-ddd.dddd, no exponent),
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'g' ('e' for large exponents, 'f' otherwise),
	// 'G' ('E' for large exponents, 'f' otherwise),
	// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
	// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
	var b float64 = 2222 * 1023 * 245 * 2 * 52
	fmt.Println(strconv.FormatFloat(b, 'e', -1, 64)) // 5.791874088e+10
}

                  
```

 Так же можно использовать пакет "fmt". Он обладает удобным методом Sprintf. Вот [шпаргалка](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) по всему пакету.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(fmt.Sprint(20.19)) // Краткая форма

    a := 20.20
    fmt.Println(fmt.Sprintf("%f", a)) // Полная форма
}
                  
```

**Внимание! Использовать fmt для конвертации нежелательно из-за того что производительность ниже по сравнению с strconv.**

## **Конвертация bool в string**

**Тут все просто:**

```go
var a = true
res := strconv.FormatBool(a)
fmt.Println(res)     	// true
fmt.Printf("%T", res)   // string
```

# Конвертация строк в другие типы

Рассмотрим для начала конвертацию **строк в целые числа** на примерах:

```go
package main

import (
    "fmt"
)

func main() {
    foo := "10"
    bar := "15"
    baz := foo - bar
    fmt.Println(baz)
}

                  
```

Если вы попробуете запустить этот код, то вы столкнетесь с ошибкой:

```go
  invalid operation: foo - bar (operator - not defined on string)
  
                  
```

Встретите вы её, так как операнд вычитания не является действительным для строк. Это можно исправить, использовав метод пакета `strconv` - `Atoi`:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	foo := "10"
	bar := "15"
	barInt, err := strconv.Atoi(bar)
	if err != nil {
		panic(err)
	}
	fooInt, err := strconv.Atoi(foo)
	if err != nil {
		panic(err)
	}
	baz := barInt - fooInt
	fmt.Println(baz) //5
}


                  
```

#### ***Важно: при конвертации строки, которая не содержит в себе число - ваша программа выдаст вам ошибку*** 

```go
strconv.Atoi: parsing "not a number": invalid syntax

                  
```

Так как метод Atoi кроме результата возвращает еще и ошибку, то мы можем легко это проверить (вспоминаем урок 2 модуля про ошибки):
 

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "323str"
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err) // strconv.Atoi: parsing "323str": invalid syntax
	} else {
		fmt.Println(result)
	}
}
                  
```

## Конвертация string в float с помощью метода `ParseFloat`:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "23.23456"
	// как и в прошлом шаге, здесь 2 параметр - bitSize
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	// но нужно понимать что метод все равно вернет float64
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)        			 // 23.23456
	fmt.Printf("%T \n", result)  // float64

	// Конкретный пример для разных bitSize:
	s = "1.0000000012345678"
	//  не будем обрабатывать ошибки в примерах, но на практике так не делайте ;)
	result32, _ := strconv.ParseFloat(s, 32)
	result64, _ := strconv.ParseFloat(s, 64)
	fmt.Println("bitSize32:", result32)  // вывод 1 (не уместились)
	fmt.Println("bitSize64:", result64)  // вывод  1.0000000012345678
}
                  
```

**Полезно знать!**
Так же по аналогии с примерами выше есть методы ParseUint, ParseInt, ParseBool.

**Кстати, метод** `Atoi` эквивалентен `ParseInt(s, 10, 0), конвертированному в int.`
Примеры:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "-12345"
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil { // не забываем проверить ошибку
		panic(err)
	}

	fmt.Println(res) // -12345

	s = "12345"
	res2, err := strconv.ParseUint(s, 10, 64)
	if err != nil {  // не забываем проверить ошибку
		panic(err)
	}
	fmt.Println(res2) // 12345
}

                  
```

 

## **Конвертация string в bool**

```go
s := "true"
res, err := strconv.ParseBool(s)
if err != nil { // не забываем проверить ошибку
	panic(err)
}
fmt.Println(res)      // true
fmt.Printf("%T", res)  // bool
```