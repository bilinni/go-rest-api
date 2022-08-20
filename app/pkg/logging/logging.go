package loging

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s;%d", filename, f.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	err := os.Mkdir("logs", 0644)
	if err != nil {
		panic(err)
	}
}
