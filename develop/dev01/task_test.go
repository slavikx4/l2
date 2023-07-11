package dev01

import (
	"bytes"
	"testing"
)

func TestPrintTimeSTDOUT(t *testing.T) {

	// Создаем буфер для перенаправления вывода
	stdout = new(bytes.Buffer)

	// Запускаем функцию PrintTime()
	PrintTime()

	res := stdout.(*bytes.Buffer).String()

	if res == "" {
		t.Errorf("ожидалось: %s\n результат: %s", "какое-то время", res)
	}
}

func TestPrintTimeSTDERR(t *testing.T) {

	stderr = new(bytes.Buffer)

	// Запускаем функцию PrintTime()
	PrintTime()

	res := stderr.(*bytes.Buffer).String()

	if res != "" {
		t.Errorf("ожидалось: %s\nв STDERR было записано: %s", "*пустота*", res)
	}
}
