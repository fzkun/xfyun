package xfyun

import (
	naturalLanguageCfg "github.com/fzkun/xfyun/config"
	"github.com/fzkun/xfyun/context"
	"github.com/fzkun/xfyun/natural_language"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// XFYun 科大讯飞sdk
type XFYun struct {
	ctx *context.Context
}

func NewXFYun(cfg *naturalLanguageCfg.Config) *XFYun {
	return &XFYun{ctx: &context.Context{
		Config: cfg,
	}}
}

// GetNaturalLanguage 获取自然语言处理实例
func (xf *XFYun) GetNaturalLanguage() *natural_language.NaturalLanguage {
	return natural_language.NewNaturalLanguage(xf.ctx)
}
