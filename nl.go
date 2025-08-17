package nl

import (
	"context"
	"fmt"
	"io"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"
	localopt "github.com/yupsh/nl/opt"
)

// Flags represents the configuration options for the nl command
type Flags = localopt.Flags

// Command implementation using StandardCommand abstraction
type command struct {
	yup.StandardCommand[Flags]
}

// Nl creates a new nl command with the given parameters
func Nl(parameters ...any) yup.Command {
	args := opt.Args[string, Flags](parameters...)
	cmd := command{
		StandardCommand: yup.StandardCommand[Flags]{
			Positional: args.Positional,
			Flags:      args.Flags,
			Name:       "nl",
		},
	}
	// Set defaults
	if cmd.Flags.BodyNumbering == "" {
		cmd.Flags.BodyNumbering = localopt.NumberNonEmpty
	}
	if cmd.Flags.HeaderNumbering == "" {
		cmd.Flags.HeaderNumbering = localopt.NumberNone
	}
	if cmd.Flags.FooterNumbering == "" {
		cmd.Flags.FooterNumbering = localopt.NumberNone
	}
	if cmd.Flags.NumberSeparator == "" {
		cmd.Flags.NumberSeparator = "\t"
	}
	if cmd.Flags.NumberFormat == "" {
		cmd.Flags.NumberFormat = "%6d"
	}
	if cmd.Flags.StartNumber == 0 {
		cmd.Flags.StartNumber = 1
	}
	if cmd.Flags.Increment == 0 {
		cmd.Flags.Increment = 1
	}
	return cmd
}

func (c command) Execute(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
	lineNumber := int(c.Flags.StartNumber)

	return c.ProcessFiles(ctx, input, output, stderr,
		func(ctx context.Context, source yup.InputSource, output io.Writer) error {
			if !bool(c.Flags.NoRenumber) && source.Filename != "stdin" {
				lineNumber = int(c.Flags.StartNumber) // Reset for each file
			}
			return c.processReader(ctx, source.Reader, output, &lineNumber)
		},
	)
}

func (c command) processReader(ctx context.Context, reader io.Reader, output io.Writer, lineNumber *int) error {
	section := "body" // Simple implementation - always treat as body

	// Use ProcessLinesSimple to eliminate manual scanner management and context checking
	return yup.ProcessLinesSimple(ctx, reader, output,
		func(ctx context.Context, lineNum int, line string, output io.Writer) error {
			if c.shouldNumber(line, section) {
				format := string(c.Flags.NumberFormat)
				separator := string(c.Flags.NumberSeparator)
				fmt.Fprintf(output, format+separator+"%s\n", *lineNumber, line)
				*lineNumber += int(c.Flags.Increment)
			} else {
				// Output line without number (with appropriate spacing)
				padding := strings.Repeat(" ", 6) // Default width
				separator := string(c.Flags.NumberSeparator)
				fmt.Fprintf(output, padding+separator+"%s\n", line)
			}
			return nil
		},
	)
}

func (c command) shouldNumber(line, section string) bool {
	var style string

	switch section {
	case "header":
		style = string(c.Flags.HeaderNumbering)
	case "footer":
		style = string(c.Flags.FooterNumbering)
	default: // body
		style = string(c.Flags.BodyNumbering)
	}

	switch style {
	case localopt.NumberAll:
		return true
	case localopt.NumberNone:
		return false
	case localopt.NumberNonEmpty:
		return strings.TrimSpace(line) != ""
	default:
		return strings.TrimSpace(line) != ""
	}
}

func (c command) String() string {
	return fmt.Sprintf("nl %v", c.Positional)
}
