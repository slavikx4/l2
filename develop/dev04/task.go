package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// входные данные
	words := &[]string{"пятак", "Пятка", "Тяпка", "листок", "листок", "слиток", "столик", "мир"}

	// поиск анаграм
	anagramSets := FindAnagramSets(words)

	// Вывод результата
	for key, value := range anagramSets {
		fmt.Printf("Множество анаграмм для слова %s: %v\n", key, value)
	}
}

// FindAnagramSets функция возвращает анаграмы
func FindAnagramSets(words *[]string) map[string][]string {

	// создание множества анаграм
	anagramSets := make(map[string][]string)

	for _, word := range *words {
		// Приводим слово к нижнему регистру и сортируем его буквы
		sortedWord := sortString(strings.ToLower(word))

		// Добавляем отсортированное слово в мапу anagramSets
		if _, ok := anagramSets[sortedWord]; !ok {
			anagramSets[sortedWord] = []string{word}
		} else {
			anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
		}
	}

	result := make(map[string][]string)

	// Удаляем множества из одного элемента и сортируем массивы слов
	for _, value := range anagramSets {
		if len(value) > 1 {
			tempRes := []string{}
			sort.Strings(value)
			for i := range value {
				value[i] = strings.ToLower(value[i])
			}

			for i := 1; i < len(value); i++ {
				if value[i-1] != value[i] {
					tempRes = append(tempRes, value[i-1])
				}
			}

			l := len(value)
			if value[l-2] != value[l-1] {
				tempRes = append(tempRes, value[l-1])
			}
			result[tempRes[0]] = tempRes
		}
	}

	return result
}

// sortString возвращает строку в которой отсортированы символы
func sortString(str string) string {

	buffer := []rune(str)

	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i] < buffer[j]
	})

	return string(buffer)
}
