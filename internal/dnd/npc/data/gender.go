package data

func GetGender() string {
	genderNumber := random.Intn(2)

	if genderNumber == 1 {
		return "Male"
	} else {
		return "Female"
	}
}
