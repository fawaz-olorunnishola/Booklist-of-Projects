package piscinego

func StrRev(s string) string {
	var str2 string
	for i := 0; i < len(s); i++ {
		str2 += string(s[len(s)-i-1])
	}
	return str2
}
