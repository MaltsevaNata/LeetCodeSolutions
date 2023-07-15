func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 1 { // [73,74,75,71,69,72,76,73]
		return []int{0}
	}

	output := make([]int, len(temperatures)) // [1,1,4,3,2,1,0,0]
    save := make(map[int][]int) // temperature: index
    save[temperatures[0]] = []int{0}

    for i := 1; i < len(temperatures); i++ {
        for key, idxs := range save {
            if temperatures[i] > key {
                for _, idx := range idxs {
                    output[idx] = i - idx
                }
                delete(save, key)
            }
        }
        if _, ok := save[temperatures[i]]; ok {
            save[temperatures[i]] = append(save[temperatures[i]], i)
        } else {
            save[temperatures[i]] = []int{i}
        }
	}
	return output
}