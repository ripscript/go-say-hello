package go_say_helo

func GetHelloWorld(name string) string {
	return "Hello world, " + name
}

func SumInt(number ...int) int {
	var result int
	for _, v := range number {
		result += v
	}

	return result
}
