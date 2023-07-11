package main

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"github.com/spf13/viper"

	"log"
	"net/http"
)

// DB база данных
var DB = make(map[int]*user)

// main старт программы
func main() {

	// инициализация конфига
	if err := initConfig(); err != nil {
		log.Fatalln("не удалось прочитать конфиг: ", err)
	}
	// получаем порт из кофнига
	port := viper.GetString("port")

	// инициализация роутера/маршрутизатора
	router := http.NewServeMux()

	// инициализация сервиса
	service := &service{}

	// инициализация общего хендлера
	gHandler := newGeneralHandler(service)

	// добавляем пути маршрутизатору, прогоняя хендлер через мидлвар
	router.Handle("/create_event", middleware(gHandler.handlerCreateEvent))
	router.Handle("/update_event", middleware(gHandler.handlerUpdateEvent))
	router.Handle("/delete_event", middleware(gHandler.handlerDeleteEvent))
	router.Handle("/events_for_day", middleware(gHandler.handlerEventsForDay))
	router.Handle("/events_for_week", middleware(gHandler.handlerEventsForWeek))
	router.Handle("/events_for_month", middleware(gHandler.handlerEventsForMonth))

	// запуск прослушивания и обслуживания сервера
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalln("не удалось прослушать и обслужить сервер: ", err)
	}
}

// initConfig функция инициализации конфига
func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("develop/dev11")
	return viper.ReadInConfig()
}
