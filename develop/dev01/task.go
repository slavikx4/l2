package dev01

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"github.com/beevik/ntp"
	"io"
	"os"
)

// глобальные перменные для тестов
var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
)

// PrintTime функция печати времени
func PrintTime() {

	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, _ = fmt.Fprintln(stderr, err.Error())
		os.Exit(1)
	}
	_, _ = fmt.Fprintln(stdout, time)
}
