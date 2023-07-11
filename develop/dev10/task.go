package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// парсинг флагов при запуске программы
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	args := flag.Args()

	// проверка что были введены все параметры
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet [--timeout=<duration>] <host> <port>")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]

	// установление соединения с указанием таймаута
	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		os.Exit(1)
	}

	defer func() { _ = conn.Close() }()

	// запуск обработчика завершения работы
	go handleSignal(conn)

	// запуск обработчика соединения
	if err := handleConnection(conn); err != nil {
		panic(err)
	}

}

// handleConnection функция обработчик соединения
func handleConnection(conn net.Conn) error {

	for {
		// считываем сообщение из стандартного ввода
		reader := bufio.NewReader(os.Stdin)
		request, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		// отправляем сообщение на сервер
		if _, err := fmt.Fprintf(conn, request+"\n"); err != nil {
			return err
		}

		// прослушиваем ответ
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return err
		}

		// печатаем ответ в стандартный вывод
		fmt.Println(response)
	}
}

// handleSignal функция обработчик обработчик сигнала для заакрытия соединения
func handleSignal(conn net.Conn) {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	// Закрытие сокета при получении сигнала
	func() { _ = conn.Close() }()
	os.Exit(0)

}
