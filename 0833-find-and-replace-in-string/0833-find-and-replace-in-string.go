type Operation struct {
    Idx int
    Source []rune
    Target []rune
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
    ops := []*Operation{}
    for i, idx := range indices {
        ops = append(ops, &Operation{Idx: idx, Source: []rune(sources[i]),
                                     Target: []rune(targets[i])})
    }
    
    sort.Slice(ops, func(i, j int) bool {
        return ops[i].Idx < ops[j].Idx
    })
    
    result := []rune{}
    ss := []rune(s)
    lastReplaced := -1
    
    for _, op := range ops {
        repl := true
        for i, sym := range op.Source {
            if op.Idx+i > len(ss)-1 || ss[op.Idx+i] != sym {
                repl = false
                break
            }
        }
        if repl {
            for i := lastReplaced+1; i< op.Idx; i++ {
                result = append(result, ss[i])
            }
            result = append(result, op.Target...)
            lastReplaced = op.Idx + len(op.Source) - 1
        }
    }
    for i := lastReplaced+1; i< len(ss); i++ {
        result = append(result, ss[i])
    }
    return string(result)
}