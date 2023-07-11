package main

import (
	"log"
	"net/http"
	"time"
)

// iService интерфейс поведения сервиса
type iService interface {
	createEvent(request *tokenRequest) (*tokenPostResponse, *errorHttp)
	updateEvent(request *tokenRequest) (*tokenPostResponse, *errorHttp)
	deleteEvent(request *tokenRequest) (*tokenPostResponse, *errorHttp)
	getEventsForDay(request *tokenRequest) (*tokenGetResponse, *errorHttp)
	getEventsForWeek(request *tokenRequest) (*tokenGetResponse, *errorHttp)
	getEventsForMonth(request *tokenRequest) (*tokenGetResponse, *errorHttp)
}

// service структура реализующая интерфейс iService
type service struct{}

// (s *service) createEvent метод создающий новое событие
func (s *service) createEvent(request *tokenRequest) (*tokenPostResponse, *errorHttp) {

	// парсинг даты(так как в tokenRequest дата записана как строка,
	// а немного ниже поиск в мапе будет происходить по time.Time
	var t time.Time
	if err := parseTime(&t, request.Data); err != nil {
		return nil, &errorHttp{Error: "не правильно задана дата", Code: http.StatusBadRequest}
	}

	// проверка существует ли пользователь с данным инд. номером
	if u, ok := DB[request.UserID]; ok {

		// если да то добавляем событие
		u.events[t] = append(u.events[t], event{request.NameEvent, t})

		// если нет, то
	} else {
		// создаём новго пользователя
		u := &user{
			UserID: request.UserID,
			events: make(map[time.Time][]event),
		}

		// добоавляем только что созданног пользователя в бд
		DB[request.UserID] = u

		// доавляем событие в события пользователя
		u.events[t] = append(u.events[t], event{request.NameEvent, t})
	}

	// возвращаем ответ
	return &tokenPostResponse{Result: "данные установлены", Code: http.StatusOK}, nil
}

// (s *service) updateEvent метод обновляющий событие
func (s *service) updateEvent(request *tokenRequest) (*tokenPostResponse, *errorHttp) {

	// удалить старое событие
	if _, err := s.deleteEvent(request); err != nil {
		return nil, err
	}

	// создаём новый запрос, чтобы передать его в параметры createEvent
	newRequest := &tokenRequest{
		UserID:    request.UserID,
		NameEvent: request.NewName,
		Data:      request.NewData,
	}

	// создание нового события
	if _, err := s.createEvent(newRequest); err != nil {
		return nil, err
	}

	// возврщаем ответ
	return &tokenPostResponse{Result: "событие обновлено", Code: http.StatusOK}, nil
}

// (s *service) deleteEvent функция удаляющая событие
func (s *service) deleteEvent(request *tokenRequest) (*tokenPostResponse, *errorHttp) {

	// парсинг даты(так как в tokenRequest дата записана как строка,
	// а немного ниже поиск в мапе будет происходить по time.Time
	var t time.Time
	if err := parseTime(&t, request.Data); err != nil {
		return nil, &errorHttp{Error: "не правильно задана дата", Code: http.StatusBadRequest}
	}

	// флаг - обозначающий произошло ли удаление
	var eventDeleted bool

	// проверка существует ли пользователь с данным инд. номером
	if u, ok := DB[request.UserID]; ok {

		// проверка, существует ли события в указаную дату у пользователя
		if events, ok := u.events[t]; ok {

			// проходит по всем событиям  в указанную дату(поиск события по названию)
			for i, event := range events {

				// проверка совпадает ли название из запроса и очередного события
				if event.Name == request.NameEvent {

					// если совпадает, то удаляет элемент из списка
					u.events[t] = removeElementSliceEvents(events, i)

					// флаг становится верным
					eventDeleted = true

					// завершение цикла, т.к. дальше бессмысленно
					break
				}
			}
		}
	}

	// проверка флага
	if eventDeleted {

		// если событие было удалено, то возвращаем ответ
		return &tokenPostResponse{Result: "событие удалено", Code: http.StatusOK}, nil

	} else {

		// событие не было удалено, то возвращаем ошибку
		return nil, &errorHttp{Error: "не нашлось такого события", Code: http.StatusBadRequest}
	}
}

// (s *service) getEventsForDay метод показывающий список событий за день
func (s *service) getEventsForDay(request *tokenRequest) (*tokenGetResponse, *errorHttp) {

	// парсинг даты(так как в tokenRequest дата записана как строка,
	// а немного ниже поиск в мапе будет происходить по time.Time
	var t time.Time
	if err := parseTime(&t, request.Data); err != nil {
		return nil, &errorHttp{Error: "не правильно задана дата", Code: http.StatusBadRequest}
	}

	// проверка существует ли пользователь с данным инд. номером
	if u, ok := DB[request.UserID]; ok {

		// проверка существует ли список событий в заданный день
		if events, ok := u.events[t]; ok {

			// возвращаем ответ
			return &tokenGetResponse{Result: events, Code: http.StatusOK}, nil
		}
	}

	// если ответ не был вернут, то возвращаем ошибку
	return nil, &errorHttp{Error: "нет данных с задаными параметрами", Code: http.StatusBadRequest}
}

// (s *service) getEventsForWeek метод возвращающий список событий за неделю
func (s *service) getEventsForWeek(request *tokenRequest) (*tokenGetResponse, *errorHttp) {

	// парсинг даты(так как в tokenRequest дата записана как строка,
	// а немного ниже поиск в мапе будет происходить по time.Time
	var t time.Time
	if err := parseTime(&t, request.Data); err != nil {
		return nil, &errorHttp{Error: "не правильно задана дата", Code: http.StatusBadRequest}
	}

	// проверка существует ли пользователь с данным инд. номером
	if u, ok := DB[request.UserID]; ok {

		// промежуточный список со всеми событиями, которые входят в рамки недели
		eventsResult := make([]event, 0)

		// задаём время окончания недели
		timeWeek := t.AddDate(0, 0, 7)

		// проходим по каждому дню недели
		for tempData := t; tempData.Before(timeWeek); {

			// если события есть в очередной день
			if events, ok := u.events[tempData]; ok {

				// записываем в результативный список
				eventsResult = append(eventsResult, events...)
			}

			// увеличиваем очередной день на 1 день
			tempData = tempData.AddDate(0, 0, 1)
		}

		// возвращаем ответ
		return &tokenGetResponse{Result: eventsResult, Code: http.StatusOK}, nil
	}

	// если ответ не был вернут, возвращаем ошибку
	return nil, &errorHttp{Error: "нет данных с задаными параметрами", Code: http.StatusBadRequest}
}

// (s *service) getEventsForMonth метод возвращающий список событий за месяц
func (s *service) getEventsForMonth(request *tokenRequest) (*tokenGetResponse, *errorHttp) {

	// парсинг даты(так как в tokenRequest дата записана как строка,
	// а немного ниже поиск в мапе будет происходить по time.Time
	var t time.Time
	if err := parseTime(&t, request.Data); err != nil {
		return nil, &errorHttp{Error: "не правильно задана дата", Code: http.StatusBadRequest}
	}

	// проверка существует ли пользователь с данным инд. номером
	if u, ok := DB[request.UserID]; ok {

		// результативный список событий , содержащий события входящие в рамки месяца
		eventsResult := make([]event, 0)

		// время через месяц от заданного
		timeMonth := t.AddDate(0, 1, 0)

		// проход по каждому дню месяца
		for tempData := t; tempData.Before(timeMonth); {

			// если в очередной день есть события
			if events, ok := u.events[tempData]; ok {

				// записываем события в результативный список
				eventsResult = append(eventsResult, events...)
			}

			// увелечение очередного дня на 1 день
			tempData = tempData.AddDate(0, 0, 1)
		}

		// возвращаем ответ
		return &tokenGetResponse{Result: eventsResult, Code: http.StatusOK}, nil
	}

	// если ответ не был вернут, возвращаем ошибку
	return nil, &errorHttp{Error: "нет данных с задаными параметрами", Code: http.StatusBadRequest}
}

// parseTime функция парсящаю строку во время
func parseTime(t *time.Time, str string) error {

	// парсинг строки во время определённого формата
	tempT, err := time.Parse("2006-01-02", str)
	if err != nil {
		log.Println(str, err)
		return err
	}

	// передаём значение
	*t = tempT

	return nil
}

// removeElementSliceEvents функция удаляющая event по индексу из []event
func removeElementSliceEvents(slice []event, index int) []event {
	return append(slice[:index], slice[index+1:]...)
}
