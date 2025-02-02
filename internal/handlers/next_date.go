package handlers

import (
	repetitionrule "final_task/internal/repetitionRule"
	"net/http"
	"time"
)

func NextDate(w http.ResponseWriter, r *http.Request) {
	now, err := time.Parse("20060102", r.FormValue("now"))
	if err != nil {
		http.Error(w, "invalid 'now' format", http.StatusBadRequest)
		return
	}

	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	repeatDate, err := repetitionrule.RepetitionRule(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte(repeatDate))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
