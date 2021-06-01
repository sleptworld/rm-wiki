package main

import (
	"github.com/sleptworld/test/Model"
)



func main() {


	//b := []Model.NewEntry{
	//	{
	//		Title:      "t",
	//		Content:    "t",
	//		Tags:       []string{"a","b","c"},
	//	},
	//	{
	//		Title: "tl",
	//		Content: "a",
	//		Tags: []string{"a","c","d"},
	//	},
	//}

	c := Model.Reg{Name: "hi"}

	a := Model.UserModel{}

	a.InitModel(&c)
	return
}
