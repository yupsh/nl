package command

type BodyNumbering string
type HeaderNumbering string
type FooterNumbering string
type NumberSeparator string
type NumberFormat string
type StartNumber int
type Increment int

const (
	NumberAll      = "a"
	NumberNone     = "n"
	NumberNonEmpty = "t"
)

type NoRenumberFlag bool

const (
	NoRenumber NoRenumberFlag = true
	Renumber   NoRenumberFlag = false
)

type flags struct {
	BodyNumbering   BodyNumbering
	HeaderNumbering HeaderNumbering
	FooterNumbering FooterNumbering
	NumberSeparator NumberSeparator
	NumberFormat    NumberFormat
	StartNumber     StartNumber
	Increment       Increment
	NoRenumber      NoRenumberFlag
}

func (b BodyNumbering) Configure(flags *flags)   { flags.BodyNumbering = b }
func (h HeaderNumbering) Configure(flags *flags) { flags.HeaderNumbering = h }
func (f FooterNumbering) Configure(flags *flags) { flags.FooterNumbering = f }
func (n NumberSeparator) Configure(flags *flags) { flags.NumberSeparator = n }
func (n NumberFormat) Configure(flags *flags)    { flags.NumberFormat = n }
func (s StartNumber) Configure(flags *flags)     { flags.StartNumber = s }
func (i Increment) Configure(flags *flags)       { flags.Increment = i }
func (n NoRenumberFlag) Configure(flags *flags)  { flags.NoRenumber = n }
