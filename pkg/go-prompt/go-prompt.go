package go_prompt

import (
	"os"
	"os/user"
	"strings"
	"time"

	config "color/pkg/go-prompt-config"

	"github.com/fatih/color"
)

func GetPrompt() string {
	cfg := config.NewConfig()
	err := cfg.Load()
	if err != nil {
		// TODO: Log error
		// We can't continue if we don't have a log file.
		panic(err)}

	return handleConfig(cfg)
}

func handleConfig(cfg *config.Config) string {
	result := ""

	for index := range cfg.Sections {
		result += handleSection(cfg, index)
	}

	return result
}

func handleSection(cfg *config.Config, index int) string {
	result := ""

	switch SectionType(cfg.Sections[index].SectionType) {
	case SectionTypeText:
		result += handleSectionTypeText(cfg, index)
	case SectionTypeSeparator:
		result += handleSectionTypeSeparator(cfg, index)
	case SectionTypePWD:
		result += handleSectionTypePWD(cfg, index)
	case SectionTypeUserName:
		result += handleSectionTypeUserName(cfg, index)
	case SectionTypeComputerName:
		result += handleSectionTypeComputerName(cfg, index)
	case SectionTypeDateTime:
		result += handleSectionTypeDateTime(cfg, index)
	case SectionTypeDrive:
		result += handleSectionTypeDrive(cfg, index)
	}

	return result
}

func handleSectionTypeText(cfg *config.Config, index int) string {
	s := cfg.Sections[index]
	c := createColor(s.ForeGroundColor, s.BackGroundColor)
	c = addStyles(s, c)
	return c.Sprintf("%s%s%s", s.Prefix, s.Text, s.Suffix)
}

func handleSectionTypeSeparator(cfg *config.Config, index int) string {
	// Get colors from previous and/or next sections.
	s := cfg.Sections[index]
	var fg, bg string

	// ForeGroundColor
	// * Primarily from the config file
	// * Secondarily from the previous item
	// * Otherwise empty
	if s.ForeGroundColor != "" {
		fg = s.ForeGroundColor
	} else if index > 0 {
		fg = cfg.Sections[index-1].BackGroundColor
	}

	// BackGroundColor
	// * Primarily from the config file
	// * Secondarily from the following item
	// * Otherwise empty
	if s.BackGroundColor != "" {
		bg = s.BackGroundColor
	} else if index < len(cfg.Sections)-1 {
		bg = cfg.Sections[index+1].BackGroundColor
	}

	c := createColor(fg, bg)
	c = addStyles(s, c)
	return c.Sprintf("%s%s%s", s.Prefix, s.Text, s.Suffix)
}

func handleSectionTypePWD(cfg *config.Config, index int) string {
	s := cfg.Sections[index]
	c := createColor(s.ForeGroundColor, s.BackGroundColor)
	c = addStyles(s, c)
	return c.Sprintf("%s%s%s", s.Prefix, getCurrentPath(), s.Suffix)
}

func handleSectionTypeUserName(cfg *config.Config, index int) string {
	s := cfg.Sections[index]
	c := createColor(s.ForeGroundColor, s.BackGroundColor)
	c = addStyles(s, c)
	u, _ := user.Current()
	return c.Sprintf("%s%s%s", s.Prefix, u.Username, s.Suffix)
}

func handleSectionTypeComputerName(cfg *config.Config, index int) string {
	s := cfg.Sections[index]
	c := createColor(s.ForeGroundColor, s.BackGroundColor)
	c = addStyles(s, c)
	host, _ := os.Hostname()
	return c.Sprintf("%s%s%s", s.Prefix, host, s.Suffix)
}

func handleSectionTypeDateTime(cfg *config.Config, index int) string {
	s := cfg.Sections[index]
	c := createColor(s.ForeGroundColor, s.BackGroundColor)
	c = addStyles(s, c)
	return c.Sprintf("%s%s%s", s.Prefix, time.Now().Format(s.Format), s.Suffix)
}

func handleSectionTypeDrive(cfg *config.Config, index int) string {
	s := cfg.Sections[index]
	c := createColor(s.ForeGroundColor, s.BackGroundColor)
	c = addStyles(s, c)

	return c.Sprintf("%s%s%s", s.Prefix, getFreeSpace(s.Format), s.Suffix)
}

func addStyles(cs config.ConfigSection, c *color.Color) *color.Color {
	if SectionStyle(cs.Styles) == SectionStyleNone || SectionStyle(cs.Styles) == SectionStyleEmpty {
		return c
	}
	styles := strings.Split(cs.Styles, ",")
	for _, style := range styles {
		if s, ok := toStyle(style); ok {
			c.Add(s)
		}
	}
	return c
}

func createColor(fg, bg string) *color.Color {
	c := color.New(color.Reset)
	c.EnableColor()
	if fg != "" {
		if col, ok := toFgColor(fg); ok {
			c.Add(col)
		}
	}
	if bg != "" {
		if col, ok := toBgColor(bg); ok {
			c.Add(col)
		}
	}
	return c
}

func toFgColor(col string) (color.Attribute, bool) {
	switch SectionColor(strings.ToLower(col)) {
	case SectionColorBlack:
		return color.FgBlack, true
	case SectionColorRed:
		return color.FgRed, true
	case SectionColorGreen:
		return color.FgGreen, true
	case SectionColorYellow:
		return color.FgYellow, true
	case SectionColorBlue:
		return color.FgBlue, true
	case SectionColorMagenta:
		return color.FgMagenta, true
	case SectionColorCyan:
		return color.FgCyan, true
	case SectionColorWhite:
		return color.FgWhite, true
	case SectionColorHiBlack:
		return color.FgHiBlack, true
	case SectionColorHiRed:
		return color.FgHiRed, true
	case SectionColorHiGreen:
		return color.FgHiGreen, true
	case SectionColorHiYellow:
		return color.FgHiYellow, true
	case SectionColorHiBlue:
		return color.FgHiBlue, true
	case SectionColorHiMagenta:
		return color.FgHiMagenta, true
	case SectionColorHiCyan:
		return color.FgHiCyan, true
	case SectionColorHiWhite:
		return color.FgHiWhite, true
	default:
		return color.FgBlack, false
	}
}

func toBgColor(col string) (color.Attribute, bool) {
	switch SectionColor(col) {
	case SectionColorBlack:
		return color.BgBlack, true
	case SectionColorRed:
		return color.BgRed, true
	case SectionColorGreen:
		return color.BgGreen, true
	case SectionColorYellow:
		return color.BgYellow, true
	case SectionColorBlue:
		return color.BgBlue, true
	case SectionColorMagenta:
		return color.BgMagenta, true
	case SectionColorCyan:
		return color.BgCyan, true
	case SectionColorWhite:
		return color.BgWhite, true
	case SectionColorHiBlack:
		return color.BgHiBlack, true
	case SectionColorHiRed:
		return color.BgHiRed, true
	case SectionColorHiGreen:
		return color.BgHiGreen, true
	case SectionColorHiYellow:
		return color.BgHiYellow, true
	case SectionColorHiBlue:
		return color.BgHiBlue, true
	case SectionColorHiMagenta:
		return color.BgHiMagenta, true
	case SectionColorHiCyan:
		return color.BgHiCyan, true
	case SectionColorHiWhite:
		return color.BgHiWhite, true
	default:
		return color.BgBlack, false
	}
}

func toStyle(style string) (color.Attribute, bool) {
	switch SectionStyle(style) {
	case SectionStyleBold:
		return color.Bold, true
	case SectionStyleFaint:
		return color.Faint, true
	case SectionStyleItalic:
		return color.Italic, true
	case SectionStyleUnderLine:
		return color.Underline, true
	default:
		return color.Bold, false
	}
}
