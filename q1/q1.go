package q1

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1 map[string]Student, studentData2 map[string]Student) map[string]Student {
	mergedData := make(map[string]Student)

	for key, student := range studentData1 {
		mergedData[key] = student
	}

	for key, student := range studentData2 {
		if _, exists := mergedData[key]; exists {
			for subject, grade := range student.Subjects {
				mergedData[key].Subjects[subject] = grade
			}
		} else {
			mergedData[key] = student
		}
	}
	
	return nil
}
