package args

import "os"

type Arguments struct {
	Command string
	Noun    string
}

func getSystemArgumens() []string {
	return os.Args[1:]
}
