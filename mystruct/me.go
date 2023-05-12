package mystruct

import "encoding/json"

type MyState int

type Me struct {
	state MyState
	Name  string
}

func NewMe() Me {
	me := Me{
		state: 1,
		Name:  "kokoichi206",
	}

	return me
}

func Marshal() string {
	me := Me{
		state: 1,
		Name:  "kokoichi206",
	}

	// only exported fields are marshaled.
	b, err := json.Marshal(me)
	if err != nil {
		return ""
	}

	return string(b)
}
