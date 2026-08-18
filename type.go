package impeller

type cint int32
type enum cint

type Bool uint8

func (bool Bool) Bool() bool {
	return bool != 0
}

func (bool *Bool) Set(value bool) {
	if value {
		*bool = 1
	} else {
		*bool = 0
	}
}
