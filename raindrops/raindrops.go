package raindrops

import "strconv"

func Convert(number int) string {
	factors := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	msg := ""
	for k, v := range factors {
		if number%k == 0 {
			msg += v
		}
	}

	if msg == "" {
		msg = strconv.Itoa(number)
	}
	return msg

}
