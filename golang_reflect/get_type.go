package main

import (
	"log"
	"reflect"
)

type student struct {
	Id_   int64  `json:"id"`
	Num_  string `json:"num_"`
	Name_ string `json:"name_"`
	Sex_  int    `json:"sex_"`
}

func (s *student) SetName(name string) {
	s.Name_ = name
}
func main() {
	var ff float64
	ff = 9.887
	r := reflect.TypeOf(ff)
	log.Println(r.String())

	var ss *student
	ss = new(student)
	r = reflect.TypeOf(ss)
	log.Println(r.String())
}
