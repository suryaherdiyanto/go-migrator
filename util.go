package gomigrator

func IfNe(a bool, b string) string {
	if a {
		return b
	}
	return ""
}
