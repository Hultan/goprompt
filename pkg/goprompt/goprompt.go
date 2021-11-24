package goprompt

import (
	"fmt"
	"strings"

	config "color/pkg/goprompt-config"

	"github.com/fatih/color"
)

var removedSections []int

func GetPrompt() string {
	cfg := config.NewConfig()
	err := cfg.Load()
	if err != nil {
		// TODO: Log error
		// We can't continue if we don't have a config file.
		panic(err)
	}

	// 1. Create sections
	allSections := createSections(cfg)
	// 2. Get data from sections
	data := getDataFromSections(allSections)
	// 3. Remove empty sections (Git and Go sections can be empty).
	nonEmptySections := removeEmptySections(cfg, data, allSections)
	// 4. Get the prompt
	p := getPrompt(nonEmptySections)
	result := fmt.Sprintf("%s%s%s", cfg.Prefix, p, cfg.Suffix)

	// length := goTermText.Len(result)
	return result

	// return result
}

func createSections(cfg *config.Config) []section {
	var sections []section
	for index := range cfg.Sections {
		c := configSection{cfg, index}
		var s section
		switch SectionType(cfg.Sections[index].SectionType) {
		case SectionTypeText:
			s = textSection{c}
		case SectionTypePWD:
			s = pwdSection{c}
		case SectionTypeUserName:
			s = userNameSection{c}
		case SectionTypeComputerName:
			s = computerNameSection{c}
		case SectionTypeDateTime:
			s = dateTimeSection{c}
		case SectionTypeGit:
			s = gitSection{c}
		case SectionTypeGoVersion:
			s = goVersionSection{c}
		case SectionTypeDrive:
			s = driveSection{c}
		}
		sections = append(sections, s)
	}
	return sections
}

func getDataFromSections(sections []section) []string {
	var data []string
	for i := range sections {
		data = append(data, sections[i].GetData())
	}
	return data
}

func removeEmptySections(cfg *config.Config, data []string, sections []section) []section {
	for i := len(data) - 1; i >= 0; i-- {
		if cfg.Sections[i].RemoveIfEmpty && data[i] == "" {
			removedSections = append(removedSections, i)
			data = append(data[:i], data[i+1:]...)
			sections = append(sections[:i], sections[i+1:]...)
		}
	}
	return sections
}

func getPrompt(sections []section) string {
	result := ""
	for i := range sections {
		result += sections[i].GetSection()
	}
	return result
}

func getSectionSeparator(cfg *config.Config, index int) string {
	s := cfg.Sections[index]
	var fg, bg string

	// ForeGroundColor
	// * Primarily from the config file
	// * Secondarily from the previous item
	// * Otherwise empty
	if s.SeparatorFgColor != "" {
		fg = s.SeparatorFgColor
	} else {
		fg = cfg.Sections[index].BgColor
	}

	// BackGroundColor
	// * Primarily from the config file
	// * Secondarily from the following item
	// * Otherwise empty
	if s.SeparatorBgColor != "" {
		bg = s.SeparatorBgColor
	} else if index < len(cfg.Sections)-1 {
		// Loop through the remaining sections and ignore removed sections.
		for i := index + 1; i < len(cfg.Sections); i++ {
			if isSectionRemoved(i) {
				continue
			}
			bg = cfg.Sections[i].BgColor
			break
		}
	}

	c := createColor(fg, bg)
	c = addStyles(s.SeparatorStyles, c)
	return c.Sprintf("%s%s%s", s.SeparatorPrefix, s.Separator, s.SeparatorSuffix)
}

func isSectionRemoved(index int) bool {
	for i := 0; i < len(removedSections); i++ {
		if removedSections[i] == index {
			return true
		}
	}

	return false
}

func addStyles(styles string, c *color.Color) *color.Color {
	if SectionStyle(styles) == SectionStyleNone || SectionStyle(styles) == SectionStyleEmpty {
		return c
	}
	stylesList := strings.Split(styles, ",")
	for _, style := range stylesList {
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
