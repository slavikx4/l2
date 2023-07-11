package main

import (
	"errors"
	"log"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.


Паттерн "Фасад" (Facade pattern) является поведенческим паттерном проектирования, который предоставляет простой интерфейс для взаимодействия с более сложной системой, скрывая детали ее работы. По сути, фасад обеспечивает уровень абстракции, который делает систему более легкой в использовании.

Применимость паттерна "Фасад" возникает, когда у вас есть сложная подсистема или набор классов, и вы хотите предоставить унифицированный интерфейс, скрывающий сложности и детали внутренней реализации. Это может быть полезно, когда вы хотите упростить взаимодействие с системой и избежать прямых зависимостей от сложных компонентов.

Плюсы использования паттерна "Фасад":

Упрощение использования: Фасад предоставляет простой интерфейс, который скрывает сложность и детали внутренней реализации системы. Это делает систему более легкой в использовании для клиентского кода.

Сокрытие сложности: Фасад изолирует клиентский код от сложных компонентов системы, предоставляя только необходимые методы и операции. Это помогает упростить взаимодействие с системой и скрыть ее сложность.

Гибкость и поддержка изменений: Фасад позволяет изменять внутреннюю реализацию системы без воздействия на клиентский код. Это обеспечивает гибкость и легкость внесения изменений в систему.

Минусы использования паттерна "Фасад":

Ограниченная функциональность: Фасад предоставляет унифицированный интерфейс для доступа к системе, что может привести к ограничению функциональности. Если клиентскому коду понадобится более сложное взаимодействие с системой, возможно, потребуется обходить фасад и работать напрямую с компонентами системы.

Дополнительный уровень абстракции: Использование фасада добавляет дополнительный уровень абстракции к системе, что может затруднить понимание и отладку кода.

Реальные примеры использования паттерна "Фасад" в Go (Golang) могут быть связаны с различными API или библиотеками. Например, вы можете создать фасад для работы с базой данных, который предоставляет простой интерфейс для выполнения запросов и скрывает сложности подключения к базе данных и манипулирования данными. Фасад может обрабатывать взаимодействие с различными типами баз данных и предоставлять единый унифицированный интерфейс для клиентского кода.
*/

// main - клиент
func main() {

	// создание фасада кошелька
	walletFacade := newWalletFacade("asd", 1234)

	// добавляем монеты
	if err := walletFacade.AddMoneyToWallet("asd", 1234, 100000); err != nil {
		log.Println(err)
	}
}

// фасад кошелька - абстракция, для более простого способа использовать функционала кошелька
type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	ladger       *ladger
	notification *notification
}

// конструктор фасада кошелька
func newWalletFacade(accountID string, code uint16) *walletFacade {
	return &walletFacade{
		account:      newAccount(accountID),
		wallet:       &wallet{},
		securityCode: newSecurityCode(code),
		ladger:       &ladger{},
		notification: &notification{},
	}
}

// функция добавляющая сколько-то монет в кошелёк
func (w *walletFacade) AddMoneyToWallet(accountID string, code uint16, amount int64) error {

	// проверка аккаунта
	if err := w.account.checkAccountID(accountID); err != nil {
		return err
	}

	// провекра секретного кода
	if err := w.securityCode.checkSecurityCode(code); err != nil {
		return err
	}

	// добавление монет
	w.wallet.addMoney(amount)

	// создание записи в бухгалтерскую книгу
	w.ladger.createRecord()

	// отправка уведомления
	w.notification.sendNotification()

	return nil
}

// аккаунт клиента
type account struct {
	accountID string
}

// конструктор аккаунта
func newAccount(accountID string) *account {
	return &account{
		accountID: accountID,
	}
}

// проверека ID аккаунта
func (a *account) checkAccountID(accountID string) error {

	// проверка сходится ли ID
	if accountID != a.accountID {
		return errors.New("error check amountID")
	}

	return nil
}

// кошелёк аккаунта
type wallet struct {
	amount int64
}

// функция добавление монет в кошелёк
func (w *wallet) addMoney(amount int64) {
	w.amount += amount
}

// секретный код аккаунта
type securityCode struct {
	code uint16
}

// конструктор секретного кода
func newSecurityCode(code uint16) *securityCode {
	return &securityCode{
		code: code,
	}
}

// проверка секретного кода
func (c *securityCode) checkSecurityCode(code uint16) error {

	if code != c.code {
		return errors.New("error check security code")
	}

	return nil
}

// записывальщик в бухгалтерскую книгу
type ladger struct{}

func (l *ladger) createRecord() {
	// записывание
}

// отправщик уведомления
type notification struct{}

func (n *notification) sendNotification() {
	// отправка
}
