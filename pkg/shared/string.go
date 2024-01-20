package shared

func UpperCamelCaseToLowerCameCase(s string) string {
	if len(s) == 0 || s[0] < 'A' || s[0] > 'Z' {
		return s
	}

	return string(s[0]+32) + s[1:]
}
