package dbgenerators

import "strconv"

// GenerateCoursesData - generate test data for table 'courses'
func GenerateCoursesData(rowsNum, teachersNum, institutionsNum int) [][]string {

	result := make([][]string, rowsNum+1)

	result[0] = []string{"teacher_id", "institution_id", "name", "description", "duration", "registration_date"}

	for i := 1; i <= rowsNum; i++ {
		teacherID := getRandomInt64(1, int64(teachersNum))
		institutionID := getRandomInt64(1, int64(institutionsNum))

		name, _ := getRandomValue(coursesNames)
		description, _ := getRandomValue(coursesDescriptions)
		duration := getRandomInt64(7200, 360000)
		date := generateDate(2015, 2019)

		result[i] = []string{
			strconv.FormatInt(teacherID, 10),
			strconv.FormatInt(institutionID, 10),
			name,
			description,
			strconv.FormatInt(duration, 10),
			date,
		}
	}

	return result
}
