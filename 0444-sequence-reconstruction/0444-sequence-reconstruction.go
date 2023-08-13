func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func sequenceReconstruction(nums []int, sequences [][]int) bool {
    children := make(map[int]map[int]struct{})
    parents := make(map[int]map[int]struct{})
    ordered := []int{}
    
    for _, seq := range sequences {
        if _, ok := parents[seq[0]]; !ok {
            parents[seq[0]] = make(map[int]struct{})
        }
        for i := 1; i < len(seq); i++ {

            if _, ok := children[seq[i-1]]; ok {
                children[seq[i-1]][seq[i]] = struct{}{}
            } else {
                children[seq[i-1]] = make(map[int]struct{})
                children[seq[i-1]][seq[i]] = struct{}{}
            }
            
            if _, ok := parents[seq[i]]; ok {
                parents[seq[i]][seq[i-1]] = struct{}{}
            } else {
                parents[seq[i]] = make(map[int]struct{})
                parents[seq[i]][seq[i-1]] = struct{}{}
            }
        }
    }
    
    queue := []int{}
    for key, val := range parents {
        if len(val) == 0 {
            queue = append(queue, key)
        }
    }
    if len(queue) > 1 {
        return false
    }
    
    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]
        ordered = append(ordered, item)
        for child, _ := range children[item] {
            delete(parents[child], item)
            if len(parents[child]) == 0 {
                queue = append(queue, child)
            }
        }
        if len(queue) > 1 {
            return false
        }
    }
    return equal(nums, ordered)
}