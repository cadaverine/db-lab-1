package dbgenerators

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func getRandomValue(values []string) (string, error) {
	length := len(values)

	if length < 1 {
		return "", errors.New("Error: values must be more than 0")
	}

	return values[rand.Intn(length)], nil
}

func generateAddress(country, city string) string {
	return country + ", " + city + " city, " + strconv.Itoa(rand.Intn(49)+1)
}

func generatePhone(head string, length int) string {
	tailLength := length - len(head)
	result := head

	for i := 0; i < tailLength; i++ {
		result += strconv.Itoa(rand.Intn(10))
	}

	return result
}

func generateEmail(head, tail string) string {
	domains := []string{"mail", "gmail", "outlook", "yahoo", "ya"}
	zones := []string{"ru", "com", "org", "tech", "io"}

	domain, _ := getRandomValue(domains)
	zone, _ := getRandomValue(zones)

	name := strings.ReplaceAll(strings.ToLower(head+"_"+tail), " ", "")

	return name + "@" + domain + "." + zone
}

func generateDate(fromYear, toYear int) string {
	min := time.Date(fromYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(toYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02")
}
