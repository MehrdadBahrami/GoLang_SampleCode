package ConstantsVar

import (
	"math"
	"strconv"
)

func DefineConstantsVariable() string {

	const a = "constant"
	const b = 500000000
	const c = 3e20 / b
	const d = int64(c)
	const e = float64(c)
	var f = math.Sin(e)

	return "a Value: " + a + "\n" +
		"b Value: " + strconv.Itoa(b) + "\n" +
		"c Value: " + strconv.Itoa(c) + "\n" +
		"d Value: " + strconv.FormatInt(d, 32) + "\n" +
		"e Value: " + strconv.FormatFloat(e, 'f', -1, 64) + "\n" +
		"f Value: " + strconv.FormatFloat(f, 'f', -1, 64)
}
