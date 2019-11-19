package dbgenerators

// GenerateStudentsData - generate test data for table 'students'
func GenerateStudentsData(rowsNum int) [][]string {
	firstNames := []string{
		"Boris", "Egor", "Fyodor", "Gleb", "Ivan",
		"Konstantin", "Maksim", "Mikhail", "Oksana", "Olga",
		"Sergei", "Vadim", "Yevgeni", "Yuri", "Anya",
		"Luba", "Olga", "Svetlana", "Tatiana", "Yula",
	}

	lastNames := []string{
		"Smith", "Johnson", "Williams", "Jones", "Brown",
		"Davis", "Miller", "Wilson", "Moore", "Taylor",
		"Anderson", "Thomas", "Jackson", "White", "Harris",
		"Martin", "Thompson", "Garcia", "Martinez", "Robinson",
		"Clark", "Rodriguez", "Lewis", "Lee", "Walker",
	}

	result := make([][]string, rowsNum+1)

	result[0] = []string{"first_name", "last_name", "phone", "email", "birthdate"}

	for i := 1; i <= rowsNum; i++ {
		firstName, _ := getRandomValue(firstNames)
		lastName, _ := getRandomValue(lastNames)
		phone := generatePhone("89", 11)
		email := generateEmail(firstName, lastName)
		date := generateDate(1970, 2002)

		result[i] = []string{firstName, lastName, phone, email, date}
	}

	return result
}
