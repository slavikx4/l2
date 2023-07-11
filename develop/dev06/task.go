package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// cutFlag структура содержащая флаги программы
type cutFlag struct {
	f int
	d string
	s bool
}

func main() {

	// парсинг флагов запуска программы
	cf := &cutFlag{}

	flag.IntVar(&cf.f, "f", 0, "select fields (columns)")
	flag.StringVar(&cf.d, "d", " ", "use a different separator")
	flag.BoolVar(&cf.s, "s", false, "delimited lines only")
	flag.Parse()

	// сканер для чтения из STDIN
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		text := scanner.Text()

		lines := strings.Split(text, "\n")

		counterCharInLine := make([]int, 0)

		words := make([][]string, 0)

		for _, line := range lines {

			counterCharInLine = append(counterCharInLine, len(line))

			words = append(words, strings.Split(line, cf.d))
		}

		printResult(&words, &counterCharInLine, cf)
	}
}

// printResult функция выводит результат в соответсвии с задаными флагами программы
func printResult(words *[][]string, counterCharInLine *[]int, cf *cutFlag) {

	for i, line := range *words {

		if cf.f > 0 {
			if len(line) >= cf.f {
				fmt.Println(line[cf.f-1] + "\t")
			}
		} else {

			if cf.s {

				var tempCounterChar int

				for _, word := range line {
					tempCounterChar += len(word)
				}

				if tempCounterChar != (*counterCharInLine)[i] {
					for _, word := range line {
						fmt.Print(word + "\t")
					}
					fmt.Println()
				}

			} else {
				for _, word := range line {
					fmt.Print(word + "\t")
				}
				fmt.Println()
			}
		}
	}
}
