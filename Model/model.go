package Model

import (
	"fmt"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/tools"
	"net/http"
	"reflect"
	"time"
)


type UserModel struct {
	DB.Model
}

type Model interface {
	InitModel(s interface{})
}

func (u *UserModel)InitModel(s interface{})  {
	reValue := reflect.ValueOf(s)
	if reValue.Kind() != reflect.Ptr {
		return
	}

	errHandler := func(value reflect.Value, key string, set *reflect.Value) {
		switch key {
		case "Pwd":
			pwd, _ := tools.PwdEncrypt(value.String(),Config.AesKey)
			set.Set(reflect.ValueOf(pwd))
		default:
			return
		}
	}

	switch reValue.Elem().Kind() {
	case reflect.Struct:
		a,_ := tools.ReflectReadStruct(s,&DB.User{}, errHandler)
		u.Init(&a)
	case reflect.Slice:
		a,_ := tools.ReflectSliceRead(s,&DB.User{}, errHandler)
		u.Init(&a)
	default:
		return
	}
}


func (e *EntryModel) InitModel(s interface{})  {
	reValue := reflect.ValueOf(s)
	if reValue.Kind() != reflect.Ptr {
		return
	}
	errHandler := func(value reflect.Value, key string, set *reflect.Value) {
		switch key {
		case "Tags":
			var originTags []string
			for tag := 0;tag < value.Len();tag++{
				t := value.Index(tag)
				originTags = append(originTags,t.String())
			}
			convertTags := DB.Tags2Entry(originTags)
			set.Set(reflect.ValueOf(convertTags))

		case "Cat":
			var cat DB.Cat
			_, err2 := DB.CatCheck(value.String(),&cat)
			if err2 != nil {
				return
			}
			set.Set(reflect.ValueOf(cat))
		default:
			return
		}
	}

	switch reValue.Elem().Kind() {
	case reflect.Struct:
		a,_ := tools.ReflectReadStruct(s,&DB.Entry{}, errHandler)
		fmt.Println(reflect.TypeOf(a).Kind())
		e.Init(&a)
	case reflect.Slice:
		a,_ := tools.ReflectSliceRead(s,&DB.Entry{}, errHandler)
		e.Init(a)
	default:
		return
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

type Data struct {
	ID         string      `json:"id"`
	Lang       string      `json:"lang"`
	TotalItems int         `json:"totalItems"`
	Items      interface{} `json:"items"`
}

type SuccessRes struct {
	ApiVersion string            `json:"apiVersion"`
	Params     map[string]string `json:"params"`
	Data       Data              `json:"data"`
}

type Errs struct {
	Reason string
}
type Err struct {
	Code    http.ConnState
	Message string
	Errors  []Errs
}

type ErrorRes struct {
	ApiVersion string `json:"apiVersion"`
	Error      Err
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
