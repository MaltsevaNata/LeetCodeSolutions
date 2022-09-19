func maximumTime(time string) string { //hh:mm
	unknown_val := "?"
	time_rune := []rune(time)

	for idx, val := range time {
		if string(val) == unknown_val {
			switch idx {
			case 0:
				if time_rune[1] != '?' && time_rune[1] > '3' && time_rune[1] > '0' {
					time_rune[idx] = '1'
				} else {
					time_rune[idx] = '2'
				}
				break

			case 1:
				if time_rune[0] == '2' {
					time_rune[idx] = '3'
				} else {
					time_rune[idx] = '9'
				}

				break

			case 3:
				time_rune[idx] = '5'
				break

			case 4:
				time_rune[idx] = '9'
				break

			default:
				panic("Wrong input")

			}

		} else {
			time_rune[idx] = val
		}
	}
	return string(time_rune)
}