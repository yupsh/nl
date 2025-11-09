package command_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/gloo-foo/testable/assertion"
	"github.com/gloo-foo/testable/run"
	command "github.com/yupsh/nl"
)

func TestNl_Basic(t *testing.T) {
	result := run.Command(command.Nl()).
		WithStdinLines("a", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 3)
	assertion.Contains(t, result.Stdout, "1")
	assertion.Contains(t, result.Stdout, "2")
}

func TestNl_EmptyLines(t *testing.T) {
	result := run.Command(command.Nl()).
		WithStdinLines("a", "", "b").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 3)
}

func TestNl_NumberAll(t *testing.T) {
	result := run.Command(command.Nl(command.BodyNumbering(command.NumberAll))).
		WithStdinLines("a", "", "b").Run()
	assertion.NoError(t, result.Err)
	assertion.Contains(t, result.Stdout, "1")
	assertion.Contains(t, result.Stdout, "2")
	assertion.Contains(t, result.Stdout, "3")
}

func TestNl_CustomFormat(t *testing.T) {
	result := run.Command(command.Nl(command.NumberFormat("%3d"))).
		WithStdinLines("a", "b").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 2)
}

func TestNl_CustomSeparator(t *testing.T) {
	result := run.Command(command.Nl(command.NumberSeparator(" | "))).
		WithStdinLines("a").Run()
	assertion.NoError(t, result.Err)
	assertion.Contains(t, result.Stdout, " | ")
}

func TestNl_StartNumber(t *testing.T) {
	result := run.Command(command.Nl(command.StartNumber(10))).
		WithStdinLines("a", "b").Run()
	assertion.NoError(t, result.Err)
	assertion.Contains(t, result.Stdout, "10")
	assertion.Contains(t, result.Stdout, "11")
}

func TestNl_Increment(t *testing.T) {
	result := run.Command(command.Nl(command.Increment(5))).
		WithStdinLines("a", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Contains(t, result.Stdout, "1")
	assertion.Contains(t, result.Stdout, "6")
	assertion.Contains(t, result.Stdout, "11")
}

func TestNl_EmptyInput(t *testing.T) {
	result := run.Quick(command.Nl())
	assertion.NoError(t, result.Err)
	assertion.Empty(t, result.Stdout)
}

func TestNl_InputError(t *testing.T) {
	result := run.Command(command.Nl()).
		WithStdinError(errors.New("read failed")).Run()
	assertion.ErrorContains(t, result.Err, "read failed")
}

func TestNl_NumberNone(t *testing.T) {
	result := run.Command(command.Nl(command.BodyNumbering(command.NumberNone))).
		WithStdinLines("a", "b").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 2)
}

func TestNl_HeaderNumbering(t *testing.T) {
	result := run.Command(command.Nl(command.HeaderNumbering(command.NumberAll))).
		WithStdinLines("a").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestNl_FooterNumbering(t *testing.T) {
	result := run.Command(command.Nl(command.FooterNumbering(command.NumberAll))).
		WithStdinLines("a").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestNl_NoRenumber(t *testing.T) {
	result := run.Command(command.Nl(command.NoRenumber)).
		WithStdinLines("a").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestNl_CustomBodyNumberingStyle(t *testing.T) {
	// Test with a custom/unknown style (default case)
	result := run.Command(command.Nl(command.BodyNumbering("x"))).
		WithStdinLines("a", "  ").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 2)
}

func TestNl_EmptyFormatString(t *testing.T) {
	// Test with empty format to trigger the fallback
	result := run.Command(command.Nl(command.NumberFormat(""))).
		WithStdinLines("a").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestNl_NumberNone_WithContent(t *testing.T) {
	// Ensure the "not numbered" path is fully exercised
	result := run.Command(command.Nl(command.BodyNumbering(command.NumberNone))).
		WithStdinLines("line1", "line2", "line3").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 3)
	// Lines should have spacing but no actual numbers
	_ = strings.Contains // use strings to avoid unused import
}

func TestNl_TableDriven(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{"three lines", []string{"a", "b", "c"}},
		{"with empty", []string{"a", "", "b"}},
		{"single", []string{"x"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := run.Command(command.Nl()).
				WithStdinLines(tt.input...).Run()
			assertion.NoError(t, result.Err)
			assertion.Count(t, result.Stdout, len(tt.input))
		})
	}
}

