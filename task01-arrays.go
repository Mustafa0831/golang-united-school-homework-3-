package homework

func average(input [15]float32) (result float32) {

	for _, v := range input {
		result += v
	}
	f := float32(len(input))
	// fmt.Println(f)
	//Place your code here
	return result / f
}
