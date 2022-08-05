package main

func Itoa(ival int) string {
	var buf []byte
	var r []byte
	var next int
	var right int

	for {
		if ival < 0 {
			ival = -1 * ival
			r = append(r, '-')
		}
		next = ival / 10
		right = ival - next*10
		ival = next
		buf = append(buf, byte('0'+right))
		if ival == 0 {
			break
		}
	}
	for j := 0; j < len(buf); j++ {
		r = append(r, buf[len(buf)-j-1])
	}
	return string(r)
}
