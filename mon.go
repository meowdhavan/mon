package mon

type Command struct {
	Names      []string
	AboutShort string
	AboutLong  string
	Run        func() error

	subcommands map[string]*Command
	flags       []flag
	posArgs     []posArg
	errors      []error

	parent *Command
}

func (c *Command) Execute() {
	c.parseFlags()
}