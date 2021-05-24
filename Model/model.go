package Model

import (
	"net/http"
	"time"
)

type Login struct {
	Email string
	Pwd   string
}

type Reg struct {
	Name       string
	Email      string
	Pwd        string
	Country    string
	Language   string
	Sex        int8
	Profession string
}

type NewEntry struct {
	Title   string
	Author  uint
	Content string
	Tags    []string
	Cat     string
	Info    string
	Draft   bool
}

// For Show

type Tag struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Cat struct {
	ID   int32  `json:"id"`
	Path string `json:"path"`
}

type Author struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type History struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
	Info      string    `json:"info"`
}

type Draft struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AllEntry struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []Tag     `json:"tags"`
}

type SingleEntry struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
	Title     string    `json:"title"`
	Author    Author    `json:"author"`
	Tags      []Tag     `json:"tags"`
	Category  Cat       `json:"cat"`
	Content   string    `json:"content"`
	History   []History `json:"history"`
	Info      string
}

type data struct {
	ID         string      `json:"id"`
	Lang       string      `json:"lang"`
	TotalItems int         `json:"totalItems"`
	Items      interface{} `json:"items"`
}

type SuccessRes struct {
	ApiVersion string            `json:"apiVersion"`
	Params     map[string]string `json:"params"`
	Data       data              `json:"data"`
}

type errs struct {
	Reason string
}
type err struct {
	Code    http.ConnState
	Message string
	Errors  []errs
}

type ErrorRes struct {
	ApiVersion string `json:"apiVersion"`
	Error      err
}
