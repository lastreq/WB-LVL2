package main

import (
	"strconv"
	"time"
)

type Event struct {
	ID   int    `json:"ID"`
	Text string `json:"Text"`
}

type Storage struct {
	nextID   int
	Calendar map[int]map[time.Month]map[int][]Event //year month day []Event
}

//Создание ивента в указаной дате
func (store *Storage) CreateEvent(text string, date time.Time, Id int) error {
	var event Event
	if Id == 0 {
		store.nextID = store.nextID + 1
		event = Event{ID: store.nextID, Text: text}
	} else {
		event = Event{ID: Id, Text: text}
	}

	year, month, day := date.Date()

	if store.Calendar == nil {
		store.Calendar = make(map[int]map[time.Month]map[int][]Event)
	}
	if store.Calendar[year] == nil {
		store.Calendar[year] = make(map[time.Month]map[int][]Event)
	}
	if store.Calendar[year][month] == nil {
		store.Calendar[year][month] = make(map[int][]Event)
	}

	store.Calendar[year][month][day] = append(store.Calendar[year][month][day], event)

	return nil
}

//Проверка существования ивента
func (store *Storage) CheckEvent(id int) (Event, bool, []string) {
	var s []string
	for v1, years := range store.Calendar {
		for v2, months := range years {
			for v3, days := range months {
				for _, event := range days {
					if event.ID == id {
						var i int = int(v2)
						s = append(s, strconv.Itoa(v1), strconv.Itoa(i), strconv.Itoa(v3))
						return event, true, s
					}
				}
			}
		}
	}
	var t Event // пустой ивент
	return t, false, s
}

//Обновление ивента по id
func (store *Storage) UpdateEvent(id int, text string, date time.Time) (bool, error) {

	var el Event
	var ok bool
	var oldDate []string
	if el, ok, oldDate = store.CheckEvent(id); !ok {
		return false, nil
	}

	if text == "" {
		text = el.Text
	}

	oldYear, _ := strconv.Atoi(oldDate[0])
	oldMonth, _ := strconv.Atoi(oldDate[1])
	oldDay, _ := strconv.Atoi(oldDate[2])

	for idx, val := range store.Calendar[oldYear][time.Month(oldMonth)][oldDay] {
		if val.ID == id {
			store.Calendar[oldYear][time.Month(oldMonth)][oldDay] = append(store.Calendar[oldYear][time.Month(oldMonth)][oldDay][:idx], store.Calendar[oldYear][time.Month(oldMonth)][oldDay][idx+1:]...)

		}
	}

	store.CreateEvent(text, date, id)

	return true, nil
}

//DeleteById удаляет ивент по id.
func (store *Storage) DeleteById(id int) (bool, error) {
	var oldDate []string
	var ok bool
	if _, ok, oldDate = store.CheckEvent(id); !ok {
		return false, nil
	}
	oldYear, _ := strconv.Atoi(oldDate[0])
	oldMonth, _ := strconv.Atoi(oldDate[1])
	oldDay, _ := strconv.Atoi(oldDate[2])

	for idx, val := range store.Calendar[oldYear][time.Month(oldMonth)][oldDay] {
		if val.ID == id {
			store.Calendar[oldYear][time.Month(oldMonth)][oldDay] = append(store.Calendar[oldYear][time.Month(oldMonth)][oldDay][:idx], store.Calendar[oldYear][time.Month(oldMonth)][oldDay][idx+1:]...)

		}
	}

	return true, nil
}

//Возваращает []Event к данному дню
func (store *Storage) GetDay(date time.Time) ([]Event, error) {
	year, month, day := date.Date()

	return store.Calendar[year][month][day], nil
}

//Возваращает []Event к данной недели начиная с полученного дня
func (store *Storage) GetWeek(date time.Time) ([]Event, error) {

	var resp []Event
	for i := 0; i < 7; i++ {
		s, _ := store.GetDay(date)
		resp = append(resp, s...)
		date = date.AddDate(0, 0, 1)
	}
	return resp, nil
}

//Возваращает []Event к данному месяцу начиная с полученного дня
func (store *Storage) GetMonth(date time.Time) ([]Event, error) {
	year, month, _ := date.Date()
	var resp []Event

	for _, val := range store.Calendar[year][month] {

		//dateStr := fmt.Sprintf("%v-%v-%v", strconv.Itoa(idx), month, year)
		//newDate, _ := time.Parse("02-01-2006", dateStr)
		//s, _ := store.GetDay(newDate)
		resp = append(resp, val...)

	}

	return resp, nil
}
