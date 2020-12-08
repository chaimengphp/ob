package Tools

func In_array(value string,list []string) bool {
	for _,val := range list {
		if val == value {
			return true
		}
	}
	return false
}
