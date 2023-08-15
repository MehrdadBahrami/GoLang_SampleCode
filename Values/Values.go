package Values

import "strconv"

func GetSampleValue() string {
	return "go" + "lang" + "\n" +
		"First Number is: 1+1 = = = = " + strconv.Itoa(1+1) + "\n" +
		"Secound Number is: 7.0/3.0 = = = = " + strconv.FormatFloat(7.0/3.0, 'f', -1, 64) + "\n" +
		strconv.FormatBool(true && false) + "\n" +
		strconv.FormatBool(true || false) + "\n" +
		strconv.FormatBool(!true) + "\n" +
		strconv.FormatBool(!false)
}
