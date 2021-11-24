package goprompt

import (
	"os"
	"os/user"
	"strings"
	"time"

	config "color/pkg/goprompt-config"

	"golang.org/x/sys/unix"

	gitStatusPrompt "github.com/hultan/gitstatusprompt"
	"github.com/hultan/gomod"
)

type section interface {
	GetData() string
	GetSection() string
}

type configSection struct {
	cfg     *config.Config
	index   int
}

//
// textSection
//

type textSection struct{ configSection }

func (ts textSection) GetData() string {
	return ts.cfg.Sections[ts.index].Text
}

func (ts textSection) GetSection() string {
	s := ts.cfg.Sections[ts.index]
	c := createColor(s.FgColor, s.BgColor)
	c = addStyles(s.Styles, c)

	return c.Sprintf("%s%s%s%s", s.Prefix, ts.GetData(), s.Suffix, getSectionSeparator(ts.cfg, ts.index))
}

//
// pwdSection
//

type pwdSection struct{ configSection }

func (ts pwdSection) GetData() string {
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

func (ts pwdSection) GetSection() string {
	s := ts.cfg.Sections[ts.index]
	c := createColor(s.FgColor, s.BgColor)
	c = addStyles(s.Styles, c)

	return c.Sprintf("%s%s%s%s", s.Prefix, ts.GetData(), s.Suffix, getSectionSeparator(ts.cfg, ts.index))
}

//
// userNameSection
//

type userNameSection struct{ configSection }

func (ts userNameSection) GetData() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	return u.Username
}

func (ts userNameSection) GetSection() string {
	s := ts.cfg.Sections[ts.index]
	c := createColor(s.FgColor, s.BgColor)
	c = addStyles(s.Styles, c)

	return c.Sprintf("%s%s%s%s", s.Prefix, ts.GetData(), s.Suffix, getSectionSeparator(ts.cfg, ts.index))
}

//
// computerNameSection
//

type computerNameSection struct{ configSection }

func (ts computerNameSection) GetData() string {
	host, err := os.Hostname()
	if err != nil {
		return ""
	}
	return host
}

func (ts computerNameSection) GetSection() string {
	s := ts.cfg.Sections[ts.index]
	c := createColor(s.FgColor, s.BgColor)
	c = addStyles(s.Styles, c)
	return c.Sprintf("%s%s%s%s", s.Prefix, ts.GetData(), s.Suffix, getSectionSeparator(ts.cfg, ts.index))
}

//
// dateTimeSection
//

type dateTimeSection struct{ configSection }

func (ts dateTimeSection) GetData() string {
	return time.Now().Format(ts.cfg.Sections[ts.index].Format)
}

func (ts dateTimeSection) GetSection() string {
	s := ts.cfg.Sections[ts.index]
	c := createColor(s.FgColor, s.BgColor)
	c = addStyles(s.Styles, c)
	return c.Sprintf("%s%s%s%s", s.Prefix, ts.GetData(), s.Suffix, getSectionSeparator(ts.cfg, ts.index))
}

//
// gitSection
//

type gitSection struct{ configSection }

func (ts gitSection) GetData() string {
	gs := gitStatusPrompt.GitStatusPrompt{}
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	status := gs.GetPrompt(path)
	if err != nil {
		return ""
	}
	return status
}

func (ts gitSection) GetSection() string {
	s := ts.cfg.Sections[ts.index]
	c := createColor(s.FgColor, s.BgColor)
	c = addStyles(s.Styles, c)
	return c.Sprintf("%s%s%s%s", s.Prefix, ts.GetData(), s.Suffix, getSectionSeparator(ts.cfg, ts.index))
}

//
// goVersionSection
//

type goVersionSection struct{ configSection }

func (ts goVersionSection) GetData() string {
	m := gomod.GoMod{}
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	info := m.GetInfo(path)
	if info == nil {
		return ""
	}
	return info.Version()
}

func (ts goVersionSection) GetSection() string {
	s := ts.cfg.Sections[ts.index]
	c := createColor(s.FgColor, s.BgColor)
	c = addStyles(s.Styles, c)
	return c.Sprintf("%s%s%s%s", s.Prefix, ts.GetData(), s.Suffix, getSectionSeparator(ts.cfg, ts.index))
}

//
// driveSection
//

type driveSection struct{ configSection }

func (ts driveSection) GetData() string {
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
	if ts.cfg.Sections[ts.index].Format == SectionBytesFormatIEC  {
		freeSpace = byteCountIEC(free)
	} else {
		freeSpace = byteCountSI(free)
	}

	return freeSpace
}

func (ts driveSection) GetSection() string {
	s := ts.cfg.Sections[ts.index]
	c := createColor(s.FgColor, s.BgColor)
	c = addStyles(s.Styles, c)
	return c.Sprintf("%s%s%s%s", s.Prefix, ts.GetData(), s.Suffix, getSectionSeparator(ts.cfg, ts.index))
}
