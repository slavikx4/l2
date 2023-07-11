package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// парсинг флагов
	var filepath string
	var k int
	var n, r, u bool

	flag.StringVar(&filepath, "filepath", "input.txt", "enter filepath")
	flag.IntVar(&k, "k", 1, "enter number column")
	flag.BoolVar(&n, "n", false, "sort by numeric value")
	flag.BoolVar(&r, "r", false, "sort in reverse")
	flag.BoolVar(&u, "u", false, "do not output duplicate")
	flag.Parse()
	k--

	// открытие файла
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	// чтение файла
	bufferData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	_ = file.Close()

	// запись в одну большую строку
	lineData := string(bufferData)

	// создаём таблицу
	var table [][]string

	// разбиваем одну большую строку на подстроки
	for _, el := range strings.Split(lineData, "\n") {

		// берём список слов из подстроки
		tokens := strings.Fields(el)

		// добавляем в таблицу
		table = append(table, tokens)
	}

	// сортировка
	if n {

		sort.Slice(table, func(i, j int) bool {
			iv, err := strconv.Atoi(table[i][k])
			if err != nil {
				panic(err)
			}
			jv, err := strconv.Atoi(table[j][k])
			if err != nil {
				panic(err)
			}
			res := iv < jv
			if r {
				return !res
			}
			return res
		})
	} else {

		sort.Slice(table, func(i, j int) bool {

			res := table[i][k] < table[j][k]

			if r {
				return !res
			}
			return res
		})
	}

	// выбор только уникальных значений
	var uniqueLines []string
	var result string

	if u {

		uniqueLines = append(uniqueLines, strings.Join(table[0], " "))

		for i := 1; i < len(table); i++ {

			if strings.Join(table[i], " ") != strings.Join(table[i-1], " ") {

				uniqueLines = append(uniqueLines, strings.Join(table[i], " "))
			}
		}

		result = strings.Join(uniqueLines, "\n")

	} else {

		builder := &strings.Builder{}

		for _, el := range table {

			builder.WriteString(strings.Join(el, " ") + "\n")
		}

		result = builder.String()
	}

	// создание файла с ответом
	outputFile, err := os.OpenFile("outputFile.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	// запись в файл с ответом
	if _, err := io.WriteString(outputFile, result); err != nil {
		panic(err)
	}
	_ = outputFile.Close()
}
