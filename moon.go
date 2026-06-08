package moon

import "os"

func Execute(rootCmd *Command) {
	showHelp := false

	rootCmd.AddBoolFlag(&showHelp, []string{"help"}, "h", "Show help message", false)

	parser := newParser(rootCmd, os.Args)
	parser.parseFlags()

	cmd := parser.currentCmd

	if showHelp {
		printer := newPrinter(os.Stdout)
		printer.printHelp(cmd)
		os.Exit(0)
	}

	if len(parser.errors) > 0 {
		// TODO
	}

	if len(parser.warnings) > 0 {
		// TODO
	}

	cmd.Run()
}
