package golang_reflect

import (
	"log"
	"reflect"
)

func CheckStructOP() {
	var s1 = Student{
		Sex_:  1,
		Name_: "Test",
		Id_:   123,
	}
	//s1.SetName("test")

	getType := reflect.TypeOf(s1)
	log.Println("type:", getType.String())

	getValue := reflect.ValueOf(s1)
	log.Println("field:", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)              // 获取字段名称
		value := getValue.Field(i).Interface() // 获取字段的值
		tag := field.Tag.Get("json")
		log.Printf("field_name:%s, field_type:%v, field_tag:%s, field_value:%v\n", field.Name, field.Type, tag, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		log.Printf("method_name:%s, method_type:%v\n", m.Name, m.Type)
	}
}
