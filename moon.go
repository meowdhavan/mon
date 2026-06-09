package moon

import "os"

type Moon struct {
	rootCmd *Command
	Printer Printer
}

func NewMoon(rootCmd *Command) *Moon {
	p := NewDefaultPrinter(os.Stdout, false)

	return &Moon{
		rootCmd: rootCmd,
		Printer: &p,
	}
}

func (m *Moon) Execute() {
	showHelp := false

	m.rootCmd.BoolFlag(&showHelp, "help", "h", "Show help message")

	parser := newParser(m.rootCmd, os.Args)
	parser.parse()

	cmd := parser.currentCmd

	if !parser.unrecognizedSubcommand && (showHelp || cmd.Run == nil) {
		m.Printer.printHelp(cmd)
		os.Exit(0)
	}

	if parser.unrecognizedSubcommand || len(parser.errors) > 0 {
		m.Printer.printFullUsage(cmd, &parser.errors, &parser.warnings)
		os.Exit(3)
	}

	cmd.Run()
}
