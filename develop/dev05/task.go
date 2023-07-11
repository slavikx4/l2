package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// grepFlag структура содержащая значения флагов программы
type grepFlag struct {
	A, B, C       int
	c, i, v, F, n bool
}

func main() {

	// парсинг флагов
	gf := grepFlag{}

	flag.IntVar(&gf.A, "A", 0, "output N line after")
	flag.IntVar(&gf.B, "B", 0, "output N line before")
	flag.IntVar(&gf.C, "C", 0, "output N line before and after")
	flag.BoolVar(&gf.c, "c", false, "output counter line equal")
	flag.BoolVar(&gf.i, "i", false, "ignore register")
	flag.BoolVar(&gf.v, "v", false, "output inverse")
	flag.BoolVar(&gf.F, "F", false, "exact match")
	flag.BoolVar(&gf.n, "n", false, "output number line with equal")
	flag.Parse()

	// создание паттерна
	re := setPattern(&gf)

	//открытие файла с входными данными
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// преобразования []byte в string
	text := string(data)

	// обработка флага v
	if gf.v {

		result := re.Split(text, -1)

		for _, match := range result {
			fmt.Println(match)
		}

		// обработка остальных флагов
	} else {

		// поиск информации про индексы нахождений
		matches := re.FindAllStringIndex(text, -1)

		if matches == nil {

			fmt.Println("have not matches")

		} else {

			uniqueLines := map[int]struct{}{}

			lines := strings.Split(text, "\n")

			for _, match := range matches {

				startIndex := match[0]
				numberLine := getNumberLine(&lines, startIndex)

				uniqueLines[numberLine] = struct{}{}

				if gf.A != 0 {
					printAfter(numberLine, &gf, &lines, false)
				}

				if gf.B != 0 {
					printBefore(numberLine, &gf, &lines, false)
				}

				if gf.C != 0 {
					printAfterAndBefore(numberLine, &gf, &lines)
				}
			}

			if gf.c {
				printCounterUniqueLines(&uniqueLines)
			}
		}
	}

}

// getNumberLine функция поиска номера строки по индексу символа
func getNumberLine(lines *[]string, index int) int {

	sumChars := 0

	for i, line := range *lines {
		sumChars += len(line) + 1
		if sumChars > index {
			return i + 1
		}
	}
	return -1
}

// printAfter функция печати n строк после нахождения
func printAfter(numberLine int, gf *grepFlag, lines *[]string, generalPass bool) {

	var isPass bool

	for i := numberLine; (i <= numberLine+gf.A) && (i-1 < len(*lines)); i++ {

		if gf.n {
			fmt.Printf("%d\t%s\t\n", i, (*lines)[i-1])
		} else {
			fmt.Printf("%s\t\n", (*lines)[i-1])
		}
		isPass = true
	}

	if isPass && !generalPass {
		fmt.Println("--")
	}
}

// printBefore функция печати n строк до нахождения
func printBefore(numberLine int, gf *grepFlag, lines *[]string, generalPass bool) {

	var isPass bool

	for i := numberLine - gf.B; i <= numberLine && i-1 >= 0; i++ {

		if gf.n {
			fmt.Printf("%d\t%s\t\n", i, (*lines)[i-1])
		} else {
			fmt.Printf("%s\t\n", (*lines)[i-1])
		}
		isPass = true
	}

	if isPass && !generalPass {
		fmt.Println("--")
	}
}

// printAfterAndBefore функция для печати n строк до и после нахождения
func printAfterAndBefore(numberLine int, gf *grepFlag, lines *[]string) {

	printAfter(numberLine, gf, lines, true)

	printBefore(numberLine, gf, lines, true)
}

// printCounterUniqueLines функция возвращает количество уникальных строк
func printCounterUniqueLines(uniqueLines *map[int]struct{}) {
	fmt.Println(len(*uniqueLines))
}

// setPattern функция возвращает паттерн в соответсвии с задаными флагами
func setPattern(gf *grepFlag) *regexp.Regexp {

	pattern := flag.Arg(0)

	if !gf.F {

		if gf.i {
			pattern = "(?i)" + pattern
		}

	} else {

		pattern = `\b` + pattern + `\b`
	}

	re := regexp.MustCompile(pattern)

	return re
}
