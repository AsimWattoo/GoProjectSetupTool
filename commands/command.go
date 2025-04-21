package commands

type Command struct {
	name    string
	args    []string
	execute func(args []string) error
}
