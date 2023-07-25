package main

import "time"

type book struct {
	ID         string     `json:"bookId"`
	Name       string     `json:"name"`
	Year       int        `json:"year"`
	Author     string     `json:"author"`
	Summary    string     `json:"summary"`
	Publisher  string     `json:"publisher"`
	PageCount  int        `json:"pageCount"`
	ReadPage   int        `json:"readPage"`
	Finished   bool       `json:"finished"`
	Reading    bool       `json:"reading"`
	InsertedAt *time.Time `json:"insertedAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

var bookshelf []*book
