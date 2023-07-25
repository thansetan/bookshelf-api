package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type bookSimpleResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}
type bookFullResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Year       int        `json:"year,omitempty"`
	Author     string     `json:"author,omitempty"`
	Summary    string     `json:"summary,omitempty"`
	Publisher  string     `json:"publisher"`
	PageCount  int        `json:"pageCount,omitempty"`
	ReadPage   int        `json:"readPage,omitempty"`
	Finished   bool       `json:"finished"`
	Reading    bool       `json:"reading"`
	InsertedAt *time.Time `json:"insertedAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}

func responseWriter(w http.ResponseWriter, statuscode int, message string, data any) {
	var status string
	switch statuscode {
	case 200, 201:
		status = "success"
	default:
		status = "fail"
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
