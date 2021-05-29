package Model

import (
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/tools"
	"net/http"
	"time"
)

func convertWrapper(from interface{},to interface{}) (interface{},[]string)  {
	label, f,err := tools.ReadLabel(from, to)
	if err != nil {
		return nil,f
	}
	return label,f
}

type UserModel struct {
	DB.Model
}

func (u *UserModel)InitModel(s interface{})  {
	tmp ,f := convertWrapper(s,&DB.User{})
	if e,ok := tmp.(DB.User);ok{
		for _,v := range f{
			if v == "Pwd"{
				pwd := tools.ReflectGetValue(s,v)
				DB.UserPretreatment(&e,pwd.(string))
			}
		}
		u.Init(&e)
	}
}


func (e *EntryModel) InitModel(s interface{})  {
	tmp ,f := convertWrapper(s,&DB.Entry{})
	if en,ok := tmp.(DB.Entry);ok{
		for _,v := range f{
			switch v {
			case "Tags":
				tmp := tools.ReflectGetValue(s,v).([]string)
				tags := DB.Tags2Entry(tmp)
				en.Tags = tags
			default:
				return
			}
		}
		e.Init(&en)
	}
}

type EntryModel struct {
	DB.Model
}


type Login struct {
	UserModel
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type Reg struct {
	UserModel
	Name       string `json:"name"`
	Email      string `json:"email"`
	Pwd        string `json:"pwd"`
	Country    string `json:"country"`
	Language   string `json:"language"`
	Sex        int8   `json:"sex"`
	Profession string `json:"profession"`
}

type UserUpdate struct {
	UserModel
	ID          uint
	Name        string
	Pwd         string
	Avatar      string
	Profession  string
	Description string
	Site        string
	Mechanism   string
}


type NewEntry struct {
	EntryModel
	Title   string
	Content string
	Tags    []string
	Cat     string
	Info    string
	Draft   bool
}

type UpdateEntry struct {
	EntryModel
	ID      uint
	Content string
	Tags    []string
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

// user

type userGroup struct {
	GroupName string `json:"groupName"`
	Level     int8
}

type AllUser struct {
	ID        uint
	CreatedAt time.Time `json:"createdAt"`
	Name      string
	Email     string
	Birthday  time.Time
	UserGroup userGroup `json:"userGroup"`
}

type SingleUser struct {
	ID          uint
	CreateAt    time.Time
	Name        string
	Email       string
	Birthday    time.Time
	UserGroup   userGroup
	Avatar      string
	Description string
	Site        string
	Country     string
	Language    string
	Entries     []AllEntry
	EditEntries []History
	Drafts      []Draft
	Mechanism   string
	Sex         int8
	Profession  string
}
