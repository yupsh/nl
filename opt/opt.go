package opt

// Custom types for parameters
type BodyNumbering string
type HeaderNumbering string
type FooterNumbering string
type NumberSeparator string
type NumberFormat string
type StartNumber int
type Increment int

// Constants for numbering styles
const (
	NumberAll    = "a" // Number all lines
	NumberNone   = "n" // Number no lines
	NumberNonEmpty = "t" // Number non-empty lines only
)

// Boolean flag types with constants
type NoRenumberFlag bool
const (
	NoRenumber   NoRenumberFlag = true
	Renumber     NoRenumberFlag = false
)

// Flags represents the configuration options for the nl command
type Flags struct {
	BodyNumbering   BodyNumbering   // Body numbering style (default: "t")
	HeaderNumbering HeaderNumbering // Header numbering style (default: "n")
	FooterNumbering FooterNumbering // Footer numbering style (default: "n")
	NumberSeparator NumberSeparator // Separator between number and line (default: "\t")
	NumberFormat    NumberFormat    // Number format (default: "%6d")
	StartNumber     StartNumber     // Starting number (default: 1)
	Increment       Increment       // Increment (default: 1)
	NoRenumber      NoRenumberFlag  // Don't reset numbering for each file
}

// Configure methods for the opt system
func (b BodyNumbering) Configure(flags *Flags)     { flags.BodyNumbering = b }
func (h HeaderNumbering) Configure(flags *Flags)   { flags.HeaderNumbering = h }
func (f FooterNumbering) Configure(flags *Flags)   { flags.FooterNumbering = f }
func (n NumberSeparator) Configure(flags *Flags)   { flags.NumberSeparator = n }
func (n NumberFormat) Configure(flags *Flags)      { flags.NumberFormat = n }
func (s StartNumber) Configure(flags *Flags)       { flags.StartNumber = s }
func (i Increment) Configure(flags *Flags)         { flags.Increment = i }
func (n NoRenumberFlag) Configure(flags *Flags)    { flags.NoRenumber = n }
