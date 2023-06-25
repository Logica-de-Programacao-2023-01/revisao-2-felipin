package q1

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1 map[string]Student, studentData2 map[string]Student) map[string]Student {
	mapaResultante := make(map[string]Student)

	for chave, valor := range studentData1 {
		mapaResultante[chave] = valor
	}
	for chave, valor := range studentData2 {
		mapaResultante[chave] = valor
	}

	return mapaResultante
}
