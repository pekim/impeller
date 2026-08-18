package impeller

import "unsafe"

func cString(str string) *byte {
	data := []byte(str + "\x00")
	return &data[0]
}

func goString(data *byte) string {
	for i, b := range unsafe.Slice(data, 1024) {
		if b == 0 {
			return unsafe.String(data, i)
		}
	}

	return ""
}
