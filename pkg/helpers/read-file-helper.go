package helpers

import "io/ioutil"

func MyReadFile() []byte {
	b, err := ioutil.ReadFile("save.txt")
	if err != nil {
		panic(err)
	}

	return b
}
