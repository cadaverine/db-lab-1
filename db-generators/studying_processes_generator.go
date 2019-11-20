package dbgenerators

import "strconv"

// GenerateProcessesData - generate test data for table 'stydying_processes'
func GenerateProcessesData(rowsNum, coursesNum, studentsNum int) [][]string {
	result := make([][]string, rowsNum+1)

	result[0] = []string{"course_id", "student_id", "receipt_date", "deadline", "score", "status"}

	for i := 1; i <= rowsNum; i++ {
		courseID := getRandomInt64(1, int64(coursesNum))
		studentID := getRandomInt64(1, int64(studentsNum))

		date := generateDate(2018, 2019)
		deadline := generateDate(2019, 2020)
		score := getRandomInt64(0, 100)
		status, _ := getRandomValue(statuses)

		result[i] = []string{
			strconv.FormatInt(courseID, 10),
			strconv.FormatInt(studentID, 10),
			date,
			deadline,
			strconv.FormatInt(score, 10),
			status,
		}
	}

	return result
}
