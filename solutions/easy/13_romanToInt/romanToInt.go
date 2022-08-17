package romantoint

var romanMap map[string]int = map[string]int{
	"M":   1000,
	"D":   500,
	"C":   100,
	"L":   50,
	"X":   10,
	"V":   5,
	"I":   1,
	"IV":  4,
	"IX":  9,
	"XL":  40,
	"XC":  90,
	"CD":  400,
	"CM":  900,
	"MM":  2000,
	"CC":  200,
	"XX":  20,
	"II":  2,
	"MMM": 3000,
	"CCC": 300,
	"XXX": 30,
	"III": 3,
}

func romanToInt(s string) int {
	res := 0
	l := len(s)
	for l != 0 {
		if l >= 3 {
			if val, ok := romanMap[s[:3]]; ok {
				res += val
				s = s[3:]
				l = len(s)
				continue
			}
		}
		if l >= 2 {
			if val, ok := romanMap[s[:2]]; ok {
				res += val
				s = s[2:]
				l = len(s)
				continue
			}
		}
		if l >= 1 {
			if val, ok := romanMap[s[:1]]; ok {
				res += val
				s = s[1:]
				l = len(s)
				continue
			}
		}
	}

	return res
}
