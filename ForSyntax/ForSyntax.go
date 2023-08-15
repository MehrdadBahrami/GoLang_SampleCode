package forsyntax

import "strconv"

func ForLoop() string {
	tmp := ""
	for i := 0; i < 3; i++ {
		tmp = "Data Number is: " + strconv.Itoa(i) + "\n"
	}
	return tmp
}
