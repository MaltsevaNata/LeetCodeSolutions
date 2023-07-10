func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapping := map[rune][]string{'2': {"a", "b", "c"}, '3': {"d", "e", "f"},
		'4': {"g", "h", "i"}, '6': {"m", "n", "o"},
		'5': {"j", "k", "l"}, '7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"}, '9': {"w", "x", "y", "z"}}

	var totalCombinations = 1
	for _, digit := range digits {
		totalCombinations *= len(mapping[digit])
	}

	var combinations = make([]string, totalCombinations)
	for _, digit := range digits {
		repeat := totalCombinations / len(mapping[digit])
		k := 0
		for k < len(combinations) {
			for _, letter := range mapping[digit] {
				for j := 0; j < repeat; j++ {
					combinations[k] += letter
					k++
				}
			}
		}
		totalCombinations = repeat
		fmt.Println(combinations)
	}
	return combinations
}