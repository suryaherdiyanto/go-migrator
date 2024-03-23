package gomigrator

func IfNe(a interface{}, b interface{}, c string) string {
	if a != b {
		return c
	}
	return ""
}
