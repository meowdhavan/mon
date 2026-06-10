package moon

import (
	"bytes"
	"testing"
)

type CustomWriter struct {
	buf bytes.Buffer
}

func (c *CustomWriter) Write(p []byte) (n int, err error) {
	return c.buf.Write(p)
}

func (c *CustomWriter) String() string {
	return c.buf.String()
}

func TestIntroLinePrint(t *testing.T) {
	c := Command{
		Name:       "app",
		AboutShort: "short about",
		AboutLong:  "Long About Section",
	}

	c.StringFlag(nil, "test-flag", "t", "Test Flag")

	w := CustomWriter{}

	p := DefaultPrinter{
		Writer: &w,
		SuppressWarnings: false,
	}

	p.printIntroLine(&c)

	got := w.String()
	want := "app - short about\n"

	if got != want {
		t.Errorf("Intro Line mismatch; got='%s', want='%s'", got, want)
	}
}

func TestFullHelpPrint(t *testing.T) {
	rootCmd := Command{
		Name:       "app",
		AboutShort: "short about for rootCmd",
		AboutLong:  "Long About Section",
	}

	rootCmd.StringFlag(nil, "test-flag", "t", "Test Flag")
	rootCmd.StringPosArg(nil, "TEST_ARG", "")

	subCmd := Command{
		Name:       "sub",
		AboutShort: "short about subCmd",
	}

	rootCmd.Subcommand(&subCmd)

	w := CustomWriter{}

	p := DefaultPrinter{
		Writer: &w,
		SuppressWarnings: false,
	}

	p.printHelp(&rootCmd)

	got := w.String()
	want := `app - short about for rootCmd

Long About Section

Usage:
app [FLAGS] <COMMAND>

Commands:
sub  short about subCmd

Flags:
--test-flag  -t   Test Flag
`

	if got != want {
		t.Errorf("Full Help mismatch; got='%s', want='%s'", got, want)
	}
}

func TestIndentPrint(t *testing.T) {
	rootCmd := Command{
		Name:       "app",
		AboutShort: "short about for rootCmd",
		AboutLong:  "Long About Section",
	}

	rootCmd.StringFlag(nil, "test-flag", "t", "Test Flag")
	rootCmd.StringPosArg(nil, "TEST_ARG", "")

	subCmd := Command{
		Name:       "sub",
		AboutShort: "short about subCmd",
	}

	rootCmd.Subcommand(&subCmd)

	w := CustomWriter{}

	p := DefaultPrinter{
		Writer: &w,
		SuppressWarnings: false,
		IndentLength: 4,
	}

	p.printHelp(&rootCmd)

	got := w.String()
	want := `app - short about for rootCmd

Long About Section

Usage:
    app [FLAGS] <COMMAND>

Commands:
    sub  short about subCmd

Flags:
    --test-flag  -t   Test Flag
`

	if got != want {
		t.Errorf("Full Help mismatch; got='%s', want='%s'", got, want)
	}
}
