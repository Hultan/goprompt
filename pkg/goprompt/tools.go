package goprompt

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)


func GetSectionText(c *color.Color, template,data, separator string) string {
	return c.Sprintf("%s%s", strings.Replace(template,"$data$", data, -1) , separator)
}

// From https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/

// byteCountSI formats the number of bytes given as an SI formatted string (1k = 1000)
func byteCountSI(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

// byteCountIEC formats the number of bytes given as an IEC formatted string (1k = 1024)
func byteCountIEC(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
