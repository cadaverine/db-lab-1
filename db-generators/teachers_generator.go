package dbgenerators

// GenerateTeachersData - generate test data for table 'teachers'
func GenerateTeachersData(rowsNum int) [][]string {
	result := make([][]string, rowsNum+1)

	result[0] = []string{"first_name", "last_name", "phone", "email", "academic_degree"}

	for i := 1; i <= rowsNum; i++ {
		firstName, _ := getRandomValue(firstNames)
		lastName, _ := getRandomValue(lastNames)
		phone := generatePhone("89", 11)
		email := generateEmail(firstName, lastName)
		degree, _ := getRandomValue(academicDegrees)

		result[i] = []string{firstName, lastName, phone, email, degree}
	}

	return result
}
