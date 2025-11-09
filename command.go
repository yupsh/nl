package command

import (
	"fmt"
	"strings"

	gloo "github.com/gloo-foo/framework"
)

type command gloo.Inputs[gloo.File, flags]

func Nl(parameters ...any) gloo.Command {
	cmd := command(gloo.Initialize[gloo.File, flags](parameters...))
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

func (p command) Executor() gloo.CommandExecutor {
	currentNumber := int(p.Flags.StartNumber)

	return gloo.StatefulLineTransform(func(lineNum int64, line string) (string, bool) {
		// Determine numbering style (using body numbering for simplicity)
		style := string(p.Flags.BodyNumbering)
		if style == "" {
			style = NumberNonEmpty
		}

		shouldNumber := false
		switch style {
		case NumberAll:
			shouldNumber = true
		case NumberNone:
			shouldNumber = false
		case NumberNonEmpty:
			shouldNumber = strings.TrimSpace(line) != ""
		default:
			shouldNumber = strings.TrimSpace(line) != ""
		}

		if shouldNumber {
			separator := string(p.Flags.NumberSeparator)
			format := string(p.Flags.NumberFormat)
			if format == "" {
				format = "%6d"
			}

			output := fmt.Sprintf(format, currentNumber) + separator + line
			currentNumber += int(p.Flags.Increment)
			return output, true
		}

		return "      " + string(p.Flags.NumberSeparator) + line, true
	}).Executor()
}
