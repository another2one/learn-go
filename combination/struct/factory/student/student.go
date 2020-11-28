package student

type student struct{
	name string
	age int
}

func Create(name string) *student {
	var stu student
	stu.name = name
	stu.construct()
	return &stu
}


func (s *student) construct(){
	s.name += "666"
}

func (s *student) GetName() string{
	return s.name
}