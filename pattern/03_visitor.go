package main

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.

Паттерн "Посетитель" (Visitor pattern) является поведенческим паттерном проектирования, который позволяет добавлять новые операции к объектам без изменения их классов. Он достигается путем вынесения операций в отдельные классы посетителей, которые могут работать с различными объектами через единый интерфейс.

Применимость паттерна "Посетитель" возникает, когда у вас есть сложная структура объектов с различными типами, и вы хотите выполнить некоторые операции над этими объектами, не изменяя их классы. Это полезно, когда вам нужно добавить новое поведение или операции к объектам, но вы не хотите изменять их иерархию классов.

Плюсы использования паттерна "Посетитель":

Разделение операций и объектов: Паттерн "Посетитель" позволяет разделить операции от структуры объектов. Каждый посетитель реализует свою операцию, а объекты принимают посетителя и предоставляют ему доступ к своим данным. Это облегчает добавление новых операций без изменения классов объектов.

Добавление операций без изменения объектов: Посетитель позволяет добавлять новые операции к объектам, даже если их классы неизменны. Это делает паттерн особенно полезным, когда у вас есть большое количество различных операций или когда вы не хотите изменять иерархию классов.

Обработка сложной структуры объектов: Посетитель обрабатывает объекты, опираясь на их типы и структуру. Это позволяет легко реализовать операции, которые зависят от типа объекта или его связей с другими объектами.

Минусы использования паттерна "Посетитель":

Усложнение структуры кода: Использование посетителей может привести к усложнению структуры кода. Вам нужно создавать классы посетителей для каждой операции, и это может привести к увеличению числа классов и зависимостей между ними.

Нарушение инкапсуляции: Посетитель может получить доступ к приватным членам объектов, что может нарушить инкапсуляцию и создать потенциальные проблемы безопасности.

Реальные примеры использования паттерна "Посетитель" в программировании могут быть связаны с обработкой сложных структур данных или моделей. Например, визуализационные библиотеки могут использовать посетителя для обхода графических объектов и выполнения операций рендеринга или обработки событий.
*/

import (
	"fmt"
)

type shape interface {
	printType()
	accept(visitor)
}

type square struct {
	//
}

func (s *square) printType() {
	fmt.Println("square")
}

func (s *square) accept(v visitor) {
	v.acceptForSquare(s)
}

type circle struct {
	//
}

func (c *circle) printType() {
	fmt.Println("circle")
}

func (c *circle) accept(v visitor) {
	v.acceptForCircle(c)
}

type visitor interface {
	acceptForSquare(*square)
	acceptForCircle(*circle)
}

type areaCalculator struct {
	//
}

func (a *areaCalculator) acceptForSquare(s *square) {
	fmt.Println("area square")
}

func (a *areaCalculator) acceptForCircle(c *circle) {
	fmt.Println("area circle")
}

func main() {
	square := &square{}
	circle := &circle{}
	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
}
