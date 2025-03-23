package logging

import (
	"github.com/natefinch/lumberjack"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

// Config general configuration struct
type Config struct {
	Directory string `yaml:"directory"`
	Filename  string `yaml:"filename"`
	GinLog    GinLog
}

type GinLog struct {
}

// New constructor for configuration
func New(conf Config) *slog.Logger {

	path := filepath.Join(conf.Directory, conf.Filename)

	lw := lumberjack.Logger{
		Filename: path,
		MaxSize:  20,
		MaxAge:   365,
	}

	mw := io.MultiWriter(&lw, os.Stdout)

	l := slog.New(
		slog.NewTextHandler(mw, nil),
	)

	return l
}
