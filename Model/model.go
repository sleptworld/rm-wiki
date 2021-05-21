package Model

type Login struct {
	Email string
	Pwd   string
}

type Reg struct {
	Name      string
	Email     string
	Pwd       string
	Country   string
	Language  string
	Sex       int8
	Profesion string
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
