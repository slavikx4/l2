package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// UnpackString функция распаковки строки
func UnpackString(data string) (string, error) {

	// стринг билдер  для птимально быстрой конкатенации строк
	answer := strings.Builder{}

	// преобразуем в слайс рун для обработки юникод таблицы
	runeSlice := []rune(data)
	lenSlice := len(runeSlice)
	var lastRune, currentRune rune

	// проходимся по каждому символу
	for i := 0; i < lenSlice; i++ {

		currentRune = runeSlice[i]

		// в зависимости от того, цифра или буква очередной символ
		if unicode.IsLetter(currentRune) {

			if unicode.IsDigit(lastRune) {
				lastRune = 0
			}
			answer.WriteRune(lastRune)

		} else if unicode.IsDigit(lastRune) && unicode.IsDigit(currentRune) {

			return "", errors.New("некорректная строка")

		} else {

			multiplier, _ := strconv.Atoi(string(currentRune))
			answer.WriteString(strings.Repeat(string(lastRune), multiplier))
		}

		lastRune = currentRune
	}

	// проверка последнего символа
	if unicode.IsLetter(lastRune) {
		answer.WriteRune(lastRune)
	}

	// заменяем null значения на стыке конкактенации
	res := strings.Replace(answer.String(), "\x00", "", -1)

	// возвращаем ответ
	return res, nil
}

func main() {
	inputs := []string{"a4bc2d5e4", "abcd", "45", ""}

	for _, el := range inputs {
		res, err := UnpackString(el)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}
