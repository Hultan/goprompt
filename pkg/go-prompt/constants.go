/*Package go_prompt:

Lets you configure you terminal prompt by editing a JSON file

*/
package go_prompt

type SectionType string

const (
	// SectionTypeText shows a text item
	SectionTypeText SectionType = "text"
	// SectionTypeSeparator adds a separator item. If a colors are specified in the config file,
	// they will be used, otherwise the background color of the previous item will be used as
	// foreground color, and the background color of the following item will be used as background
	// color. If a previous item or a following item does not exist, an empty color will be used.
	SectionTypeSeparator = "separator"
	// SectionTypePWD shows the current path.
	SectionTypePWD = "pwd"
	// SectionTypeUserName shows the current username.
	SectionTypeUserName = "user-name"
	// SectionTypeComputerName shows the current computer name.
	SectionTypeComputerName = "computer-name"
	// SectionTypeDateTime shows the current date, time or both, formatted by the go formatting string in the
	// format field. An example formatting string would be "2006-01-02 15:04:05". Find more examples here:
	// https://pkg.go.dev/time#Time.Format.
	SectionTypeDateTime = "datetime"
	// SectionTypeGit shows the GIT status for the current repository (if the current directory has a
	// .git sub folder).
	SectionTypeGit
	// SectionTypeGoVersion shows the Go version for the current directory (if the current directory
	// contains a go.mod file).
	SectionTypeGoVersion
	// SectionTypeDrive shows the amount of disk space on the drive where the current users home folder is.
	SectionTypeDrive = "free-space"
)

type SectionStyle string

const (
	SectionStyleNone SectionStyle = "none"			// No style
	SectionStyleEmpty = ""							// No style
	SectionStyleBold = "bold"						// Bold font
	SectionStyleFaint = "faint"						// Faint font
	SectionStyleItalic = "italic"					// Italic font
	SectionStyleUnderLine = "underline"				// Underline font
)

type SectionColor string

const (
	SectionColorBlack SectionColor = "black"
	SectionColorRed = "red"
	SectionColorGreen = "green"
	SectionColorYellow = "yellow"
	SectionColorBlue = "blue"
	SectionColorMagenta = "magenta"
	SectionColorCyan = "cyan"
	SectionColorWhite = "white"
	SectionColorHiBlack = "hiblack"
	SectionColorHiRed = "hired"
	SectionColorHiGreen = "higreen"
	SectionColorHiYellow = "hiyellow"
	SectionColorHiBlue = "hiblue"
	SectionColorHiMagenta = "himagenta"
	SectionColorHiCyan = "hicyan"
	SectionColorHiWhite = "hiwhite"
)

type SectionBytesFormat string

// See https://en.wikipedia.org/wiki/Binary_prefix for more info.

const (
	SectionBytesFormatNone SectionBytesFormat = ""		// SI format, 1k = 1000
	SectionBytesFormatSI = "SI"							// SI format, 1k = 1000
	SectionBytesFormatIEC = "IEC"						// IEC format, 1k = 1024
)