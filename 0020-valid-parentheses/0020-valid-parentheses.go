func isValid(s string) bool {
    dict := map[string]int{
		"(": -1,
		")": 1, 
		"{": -100,
		"}": 100,
		"[": -1000,
		"]": 1000,
	}
	

	list1 := []int{}
	count := 0
	for _, element := range s{
		el := string(element)
		if dict[el] < 0{
			//fmt.Println(dict[el], el, dict1[el])
			list1 = append(list1, dict[el])
			count += dict[el]
		}else{
			if len(list1) > 0 && dict[el]*(-1) == list1[len(list1) -1] {
				list1 = list1[:len(list1) - 1]
				count += dict[el]
			}else{return false}
		}
		
	}
	return count == 0
}