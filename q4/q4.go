package q4

type Student struct {
	ID      int
	Name    string
	Grades  []float64
	Average float64
}

func UpdateAverage(students map[int]*Student) {
	for _, student := range students {
		totalGrades := 0.0
		for _, grade := range student.Grades {
			totalGrades += grade
		}
		student.Average = totalGrades / float64(len(student.Grades))
	}
}
