package moon

type Printer interface {
	printHelp(*Command)
	printWarnings(*[]error)
	printFullUsage(*Command, *[]error, *[]error)
}
