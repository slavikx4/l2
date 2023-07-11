package main

import "time"

// user структура содержащая информацию о пользователе
type user struct {

	// инд. номер пользователя
	UserID int `json:"user_id"`

	// события пользователя
	events // type events map[time.Time][]event
}

// event структура содержащая информацию о событии
type event struct {

	// название события
	Name string `json:"name"`

	// дата события
	Data time.Time `json:"data"`
}

// events множество событий
type events map[time.Time][]event

// tokenRequest стуктура для передачи в сервис
type tokenRequest struct {

	// инд. номер пользователя
	UserID int `json:"user_id"`

	// название события
	NameEvent string `json:"name"`

	// дата события
	Data string `json:"data"`

	// при update_event новое имя
	NewName string `json:"new_name"`

	// при update_event новая дата
	NewData string `json:"new_data"`
}

// tokenGetResponse структура в случае успешного завершения GET запроса
type tokenGetResponse struct {

	// результат - список событий
	Result []event `json:"Result"`

	// код ответа
	Code int `json:"-"`
}

// tokenPostResponse структура в случае успешного завершения POST запроса
type tokenPostResponse struct {

	// результат - сообщение о проделанном действии
	Result string `json:"Result"`

	// код ответа
	Code int `json:"-"`
}

// errorHttp структура для наполнения информации об ошибки
type errorHttp struct {

	// сообщение ошибки
	Error string `json:"Error"`

	// код ошибки
	Code int `json:"-"`
}
