package golang_reflect

import (
	"log"
	"reflect"
)


func CheckTypeOf() {
	var ff float64
	ff = 9.887
	r := reflect.TypeOf(ff)
	log.Println(r.String())

	var s1 *student
	s1 = new(student)
	s1.SetName("test")
	r = reflect.TypeOf(s1)
	log.Println(r.String())

	var s2 student
	r = reflect.TypeOf(s2)
	log.Println(r.String())

	var sli []int
	r = reflect.TypeOf(sli)
	log.Println(r.String())
}
