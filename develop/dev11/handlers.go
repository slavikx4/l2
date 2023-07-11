package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// generalHandler структура объединяющая все хендлеры
type generalHandler struct {
	*handlerCreateEvent
	*handlerUpdateEvent
	*handlerDeleteEvent
	*handlerEventsForDay
	*handlerEventsForWeek
	*handlerEventsForMonth
}

// newGeneralHandler конструктор generalHandler
func newGeneralHandler(s iService) *generalHandler {
	return &generalHandler{
		handlerCreateEvent:    &handlerCreateEvent{s},
		handlerUpdateEvent:    &handlerUpdateEvent{s},
		handlerDeleteEvent:    &handlerDeleteEvent{s},
		handlerEventsForDay:   &handlerEventsForDay{s},
		handlerEventsForWeek:  &handlerEventsForWeek{s},
		handlerEventsForMonth: &handlerEventsForMonth{s},
	}
}

// handlerCreateEvent хендлер для обработки пути "create_event"
type handlerCreateEvent struct {
	service iService
}

// (h *handlerCreateEvent) ServeHTTP метод реализующий интерфейс http.Handler
func (h *handlerCreateEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// вызов функции парсинга POST запроса
	data, err := parsePostRequest(r)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// вызов метода сервиса
	dataResponse, err := h.service.createEvent(data)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// возвращаем результат успешного завершения запроса
	sendJSONResponse(w, dataResponse.Code, dataResponse)
}

// handlerUpdateEvent хендлер для обработки пути "update_event"
type handlerUpdateEvent struct {
	service iService
}

// (h *handlerUpdateEvent) ServeHTTP метод реализующий интерфейс http.Handler
func (h *handlerUpdateEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// вызов функции парсинга POST запроса
	data, err := parsePostRequest(r)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// вызов метода сервиса
	dataResponse, err := h.service.updateEvent(data)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// возвращаем результат успешного завершения запроса
	sendJSONResponse(w, dataResponse.Code, dataResponse)
}

// handlerDeleteEvent хендлер для обработки пути "delete_event"
type handlerDeleteEvent struct {
	service iService
}

// (h *handlerDeleteEvent)  ServeHTTP метод реализующий интерфейс http.Handler
func (h *handlerDeleteEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// вызов функции парсинга POST запроса
	data, err := parsePostRequest(r)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}
	// вызов метода сервиса
	dataResponse, err := h.service.deleteEvent(data)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// возвращаем результат успешного завершения запроса
	sendJSONResponse(w, dataResponse.Code, dataResponse)
}

// handlerEventsForDay хендлер для обработки пути "events_for_day"
type handlerEventsForDay struct {
	service iService
}

// (h *handlerEventsForDay) ServeHTTP метод реализующий интерфейс http.Handler
func (h *handlerEventsForDay) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// вызов функции парсинга GET запроса
	tempRequest, err := parseGetRequest(r)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// вызов метода сервиса
	dataResponse, err := h.service.getEventsForDay(tempRequest)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// возвращаем результат успешного завершения запроса
	sendJSONResponse(w, dataResponse.Code, dataResponse)
}

// handlerEventsForWeek хендлер для обработки пути "events_for_week"
type handlerEventsForWeek struct {
	service iService
}

// (h *handlerEventsForWeek) ServeHTTP метод реализующий интерфейс http.Handler
func (h *handlerEventsForWeek) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// вызов функции парсинга GET запроса
	tempRequest, err := parseGetRequest(r)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// вызов метода сервиса
	dataResponse, err := h.service.getEventsForWeek(tempRequest)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// возвращаем результат успешного завершения запроса
	sendJSONResponse(w, dataResponse.Code, dataResponse)
}

// handlerEventsForMonth хендлер для обработки пути "events_for_month"
type handlerEventsForMonth struct {
	service iService
}

// (h *handlerEventsForMonth) ServeHTTP метод реализующий интерфейс http.Handler
func (h *handlerEventsForMonth) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// вызов функции парсинга GET запроса
	tempRequest, err := parseGetRequest(r)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// вызов метода сервиса
	dataResponse, err := h.service.getEventsForMonth(tempRequest)
	if err != nil {
		// возвращаем неуспешное завершение запроса
		sendJSONResponse(w, err.Code, err)
		return
	}

	// возвращаем результат успешного завершения запроса
	sendJSONResponse(w, dataResponse.Code, dataResponse)
}

// parseGetRequest функция парсящая Get http.Request в tokenRequest
func parseGetRequest(r *http.Request) (*tokenRequest, *errorHttp) {

	if r.Method != http.MethodGet {
		return nil, &errorHttp{Error: "неподходящий метод запроса", Code: http.StatusBadRequest}
	}

	parsedUrl, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, &errorHttp{Error: "неудалось распарсить url", Code: http.StatusInternalServerError}
	}

	queryParams := parsedUrl.Query()

	userID, err := strconv.Atoi(queryParams.Get("user_id"))
	if err != nil {
		return nil, &errorHttp{Error: "некорректно введён user_id", Code: http.StatusBadRequest}
	}
	data := queryParams.Get("data")

	tempRequest := &tokenRequest{UserID: userID, Data: data}

	return tempRequest, nil
}

// parsePostRequest функция парсящая POST http.Request
func parsePostRequest(r *http.Request) (*tokenRequest, *errorHttp) {

	if r.Method != http.MethodPost {

		return nil, &errorHttp{Error: "не подходящий метод запроса", Code: http.StatusBadRequest}
	}

	if err := r.ParseForm(); err != nil {

		return nil, &errorHttp{Error: "не удалось распарсить запрос", Code: http.StatusInternalServerError}
	}

	userID, err := isValidUserID(r.Form.Get("user_id"))
	if err != nil {
		return nil, err
	}

	tRequest := &tokenRequest{
		UserID:    userID,
		NameEvent: r.Form.Get("name"),
		Data:      r.Form.Get("data"),
		NewName:   r.Form.Get("new_name"),
		NewData:   r.Form.Get("new_data"),
	}

	if tRequest.NameEvent == "" || tRequest.Data == "" {

		return nil, &errorHttp{Error: "не введены данные запроса", Code: http.StatusBadRequest}
	}

	return tRequest, nil
}

// sendJSONResponse функция записывающая http.Response
func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {

	// Преобразование данных в формат JSON
	jsonData, err := interfaceToJSON(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Установка заголовков HTTP для указания типа контента как JSON
	w.Header().Set("Content-Type", "application/json")

	// Установка кода состояния HTTP
	w.WriteHeader(statusCode)

	// Запись JSON-документа в тело ответа
	if _, err = w.Write(jsonData); err != nil {
		log.Println("не удалось записать ответ: " + err.Error())
	}
}

// interfaceToJSON функция сериализирующая интерфейс к json
func interfaceToJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return jsonData, nil
}

// isValidUserID функция конвертирует string в int
func isValidUserID(stringUserID string) (int, *errorHttp) {

	if userID, err := strconv.Atoi(stringUserID); err != nil {

		return 0, &errorHttp{Error: "не валидный user_id", Code: http.StatusBadRequest}

	} else {
		return userID, nil
	}

}
