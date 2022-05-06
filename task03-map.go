package homework

import "sort"

func sortMapValues(input map[int]string) []string {
	result :=make([]string,len(input))
	in :=make([]int,len(input))
	i:=0
	for  k := range input {
		in[i]=k
		i++
	}
	sort.Ints(in)
	for i:=range in{
		result[i]=input[in[i]]
	}
	//Place your code here
	return result
}
