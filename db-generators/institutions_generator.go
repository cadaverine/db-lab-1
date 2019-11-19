package dbgenerators

// GenerateInstitutionsData - generate test data for
func GenerateInstitutionsData(rowsNum int) [][]string {
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
