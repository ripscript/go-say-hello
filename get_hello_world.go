package go_say_helo

func GetHelloWorld() string {
	return "Hello worlds"
}

func SumInt(number ...int) int {
	var result int
	for _, v := range number {
		result += v
	}

	return result
}
