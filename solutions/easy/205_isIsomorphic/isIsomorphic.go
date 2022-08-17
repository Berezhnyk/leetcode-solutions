package isisomorphic

func isIsomorphic(s string, t string) bool {
	return replaceWithMap(s, t) == t && replaceWithMap(t, s) == s
}

func replaceWithMap(a string, b string) string {
	l := len(a)
	m := make(map[byte]byte, l)

	for i, c := range []byte(a) {
		m[c] = b[i]
	}

	r := make([]byte, l)
	for i, c := range []byte(a) {
		r[i] = m[c]
	}
	return string(r)
}
