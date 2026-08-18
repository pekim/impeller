package impeller

func cString(str string) *byte {
	data := []byte(str + "\x00")
	return &data[0]
}
