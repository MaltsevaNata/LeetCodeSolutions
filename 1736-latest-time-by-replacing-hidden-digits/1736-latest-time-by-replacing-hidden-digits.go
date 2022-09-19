func maximumTime(time string) string { //hh:mm
	unknown_val := "?"
	var result_string []string

	for idx, val := range time {
		if string(val) == unknown_val {
			switch idx {
			case 0:
				next_val_str := string(time[1])
				next_val_int, _ := strconv.Atoi(next_val_str)
				if next_val_str != "?" && next_val_int > 3 && next_val_int > 0 {
					result_string = append(result_string, "1")
				} else {
					result_string = append(result_string, "2")
				}
				break

			case 1:
				if result_string[0] == "2" {
					result_string = append(result_string, "3")
				} else {
					result_string = append(result_string, "9")
				}

				break

			case 3:
				result_string = append(result_string, "5")
				break

			case 4:
				result_string = append(result_string, "9")
				break

			default:
				panic("Wrong input")

			}

		} else {
			result_string = append(result_string, string(val))
		}
	}
	return strings.Join(result_string, "")
}