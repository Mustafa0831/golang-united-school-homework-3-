package homework

func reverse(input []int64) (result []int64) {
	// int64AsIntValues := make([]int, len(input))

	// for i, val := range input {
	// 	int64AsIntValues[i] = int(val)
	// }
	// sort.Ints(int64AsIntValues)
	// for i, val := range int64AsIntValues {
	// 	input[i] = int64(val)
	// }
	// for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
	// 	input[i], input[j] = input[j], input[i]
	// }
	// copy(result, input)
	// //Place your code here
	// return
	for i,j:=0,len(input)-1; i<j;i,j=i+1, j-1{
		input[i],input[j]=input[j],input[i]
	}
	for _,v:=range input{
		result=append(result,v)

	}
	//Place your code here
	return
}
