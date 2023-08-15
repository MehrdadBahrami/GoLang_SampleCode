package Variables

import "strconv"

func DefineVariables(a string, b int, c int, d bool, e int, f string) string {

	return "a Value: " + a + "\n" +
		"b Value: " + strconv.Itoa(b) + "\n" +
		"c Value: " + strconv.Itoa(c) + "\n" +
		"d Value: " + strconv.FormatBool(d) + "\n" +
		"e Value: " + strconv.Itoa(e) + "\n" +
		"f Value: " + f
}
