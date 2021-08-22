package utils

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}
