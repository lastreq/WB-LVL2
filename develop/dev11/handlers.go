package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var e Storage

func hello(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

}

type EventsResponse struct {
	Data []Event `json:"result"`
}

// NewEventsResponse конструктор EventsResponse
func NewEventsResponse(data []Event) []byte {
	resp := EventsResponse{data}
	jsonResp, _ := json.Marshal(resp)
	return jsonResp
}

type ErrorResponse struct {
	Message string `json:"error"`
}

// NewErrorResponse  конструктор ErrorResponse
func NewErrorResponse(err error) []byte {
	resp := ErrorResponse{err.Error()}
	jsonResp, _ := json.Marshal(resp)
	return jsonResp
}

type EditResponse struct {
	Message string `json:"result"`
}

// NewEditResponse  конструктор EditResponse
func NewEditResponse(msg string) []byte {
	resp := EditResponse{msg}
	jsonResp, _ := json.Marshal(resp)
	return jsonResp
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	text := r.Form.Get("text")
	dateStr := r.Form.Get("date")

	date, err := time.Parse("02-01-2006", dateStr)

	if dateStr == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.CreateEvent(text, date, 0)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(NewEditResponse(fmt.Sprintf("Ивент добавлен: %v  %v", text, dateStr)))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idStr := r.Form.Get("id")
	text := r.Form.Get("text")
	dateStr := r.Form.Get("date")

	date, err := time.Parse("02-01-2006", dateStr)
	id, err := strconv.Atoi(idStr)
	if id <= 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	eventExist, err := e.UpdateEvent(id, text, date)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	if !eventExist {
		w.Write(NewEditResponse(fmt.Sprintf("Ивент с id:%v не найден\n", idStr)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(NewEditResponse(fmt.Sprintf("Ивент с id:%v обновлен\n", idStr)))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idStr := r.Form.Get("id")
	id, err := strconv.Atoi(idStr)
	if id <= 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok, err := e.DeleteById(id)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}
	if ok == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewEditResponse(fmt.Sprintf("Ивент с id:%v не найден", idStr)))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(NewEditResponse(fmt.Sprintf("Ивент с id:%v удалён", idStr)))
	}

}

// dayHandler возвращает события дня
func dayHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dateStr := r.Form.Get("date")
	date, err := time.Parse("02-01-2006", dateStr)

	jsonResp, err := e.GetDay(date)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}

	w.Write(NewEventsResponse(jsonResp))
}

// weekHandler возвращает события недели, начиная с date
func weekHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dateStr := r.Form.Get("date")
	date, err := time.Parse("02-01-2006", dateStr)

	jsonResp, err := e.GetWeek(date)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}

	w.Write(NewEventsResponse(jsonResp))
}

// weekHandler возвращает события месяца, начиная с date
func monthHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dateStr := r.Form.Get("date")
	date, err := time.Parse("02-01-2006", dateStr)

	jsonResp, err := e.GetMonth(date)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}

	w.Write(NewEventsResponse(jsonResp))
}
