package go_prompt

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"golang.org/x/sys/unix"
)

func getCurrentPath() string {
	u, err := user.Current()
	if err != nil {
		// TODO : Log error
		return "[error]"
	}
	home := u.HomeDir
	path, err := os.Getwd()
	if err != nil {
		// TODO : Log error
		return "[error]"
	}
	if strings.HasPrefix(path, home) {
		path = strings.Replace(path, home, "~", 1)
	}
	return path
}

func getFreeSpace(format string) string {
	var stat unix.Statfs_t
	wd, _ := os.Getwd()
	err := unix.Statfs(wd, &stat)
	if err != nil {
		// TODO: Log error
		return "[error]"
	}
	// Available blocks * size per block = available space in bytes
	free := stat.Bavail * uint64(stat.Bsize)
	var freeSpace string
	if format == SectionBytesFormatIEC  {
		freeSpace = byteCountIEC(free)
	} else {
		freeSpace = byteCountSI(free)
	}

	return freeSpace
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
