package golang_reflect

type student struct {
	id_   int64  `json:"id"`
	num_  string `json:"num_"`
	name_ string `json:"name_"`
	sex_  int    `json:"sex_"`
}

func (s *student) SetName(name string) {
	s.name_ = name
}
