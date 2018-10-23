package golang_reflect

import (
	"log"
	"reflect"
)

func CheckValueOf() {
	var ff float64
	ff = 9.887
	ff_r := reflect.ValueOf(ff)
	log.Println(ff_r)

	var s1 *student
	s1 = new(student)
	s1.SetName("test")
	s1_r := reflect.ValueOf(s1)
	log.Println(s1_r)

	var s2 student
	s2.SetName("test")
	s2_r := reflect.ValueOf(s2)
	log.Println(s2_r)

	var sli []int
	sli = append(sli, 1, 2, 3, 4, 5)
	sli_r := reflect.ValueOf(sli)
	log.Println(sli_r)
}
