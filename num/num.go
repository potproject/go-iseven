package num

func Parse(i int64) []string {
	n := i
	var s []string
	if n == 0 {
		return []string{"zero"}
	}
	if (n / 1000000000000000000) >= 1 {
		s = append(s, under1000num(n/1000000000000000000)...)
		s = append(s, "quintillion")
		n = n % 1000000000000000000
	}
	if (n / 1000000000000000) >= 1 {
		s = append(s, under1000num(n/1000000000000000)...)
		s = append(s, "quadrillion")
		n = n % 1000000000000000
	}
	if (n / 1000000000000) >= 1 {
		s = append(s, under1000num(n/1000000000000)...)
		s = append(s, "trillion")
		n = n % 1000000000000
	}
	if (n / 1000000000) >= 1 {
		s = append(s, under1000num(n/1000000000)...)
		s = append(s, "billion")
		n = n % 1000000000
	}
	if (n / 1000000) >= 1 {
		s = append(s, under1000num(n/1000000)...)
		s = append(s, "million")
		n = n % 1000000
	}
	if (n / 1000) >= 1 {
		s = append(s, under1000num(n/1000)...)
		s = append(s, "thousand")
		n = n % 1000
	}
	s = append(s, under1000num(n)...)
	return s
}

var a []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var b []string = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var c []string = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func under1000num(i int64) []string {
	n := i
	var s []string
	if n/100 >= 1 {
		t := (n / 100) - 1
		s = append(s, a[t], "hundred")
		n = n % 100
	}
	if n >= 20 {
		t := (n / 10) - 2
		n = n % 10
		if n > 0 {
			s = append(s, c[t]+"-"+a[n-1])
		} else {
			s = append(s, c[t])
		}
	} else if n >= 10 {
		s = append(s, b[n-10])
	} else if n >= 1 {
		s = append(s, a[n-1])
	}
	return s
}
