package Model

func IsContain(items []string , item string) bool{
	for _,eachitem := range items{
		if eachitem == item{
			return true
		}
	}
	return false
}
