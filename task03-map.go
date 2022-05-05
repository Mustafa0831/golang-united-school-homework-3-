package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	for _, v := range input {
		result = append(result, v)
	}
	sort.Strings(result)
	//Place your code here
	return
}
