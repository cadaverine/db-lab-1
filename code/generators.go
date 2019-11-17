package main

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
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

func generateInstitutionsData(rowsNum int) [][]string {
	countries := []string{"Russia"}

	cities := []string{
		"Arkhangelsk", "Astrakhan", "Belgorod", "Bryansk", "Chelyabinsk",
		"Irkutsk", "Ivanovo", "Kaliningrad", "Kaluga", "Kemerovo",
		"Khabarovsk", "Kirov", "Kostroma", "Krasnodar", "Krasnoyarsk", "Kursk",
		"Leningrad", "Lipetsk", "Moscow", "Murmansk", "Novgorod",
		"Novosibirsk", "Omsk", "Orenburg", "Oryol", "Penza",
		"Perm", "Pskov", "Rostov", "Ryazan", "Saint Petersburg",
		"Samara", "Saratov", "Smolensk", "Sverdlovsk", "Tambov",
		"Tomsk", "Tula", "Tver", "Tyumen", "Ulyanovsk",
		"Vladimir", "Volgograd", "Vologda", "Voronezh", "Yaroslavl",
	}

	types := []string{
		"College", "Institute", "University", "High School",
	}

	result := make([][]string, rowsNum+1)

	result[0] = []string{"name", "address", "phone", "email"}

	for i := 1; i <= rowsNum; i++ {
		country, _ := getRandomValue(countries)
		city, _ := getRandomValue(cities)
		institutionType, _ := getRandomValue(types)

		name := city + " " + institutionType
		address := generateAddress(country, city)
		phone := generatePhone("89", 11)
		email := generateEmail(city, institutionType)

		result[i] = []string{name, address, phone, email}
	}

	return result
}
