package golang_reflect

import (
	"log"
)

type Student struct {
	Id_   int64  `json:"id"`
	Name_ string `json:"name"`
	Sex_  int    `json:"sex"`
}

func (s *Student) SetName(name string) {
	s.Name_ = name
}

func (s Student) ReflectFuncCall() {
	log.Println("被调用了")
}
