func stringToMins(time string) int {
    arr := strings.Split(time, ":")
    hr, _ := strconv.Atoi(arr[0])
    min, _ := strconv.Atoi(arr[1])
    return hr * 60 + min
}

func haveConflict(event1 []string, event2 []string) bool {
    start1, end1 := stringToMins(event1[0]), stringToMins(event1[1])
    start2, end2 := stringToMins(event2[0]), stringToMins(event2[1])
    
    if start2 <= end1 && end2 >= start1 {
        return true
    }
    if start1 <= end2 && end1 >= start2 {
        return true
    }
    return false
}