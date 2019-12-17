package log

import (
	goLog "log"
	"os"
	"path/filepath"
	"time"
)

const logPrefix = ""
const logPath = "G:\\log\\flower"
const infoLogFileName = "info_20060102150405.log"
const warnLogFileName = "warn_20060102150405.log"
const errorLogFileName = "error_20060102150405.log"

var infoLogger *goLog.Logger
var warnLogger *goLog.Logger
var errorLogger *goLog.Logger

// 初始化log
func init() {
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		goLog.Fatalf("mk log dir error[%v]", err)
	}

	infoPath := filepath.Join(logPath, time.Now().Format(infoLogFileName))
	file, err := os.Create(infoPath)
	if err != nil {
		goLog.Fatalf("create log file error[%v]", err)
	}
	infoLogger = goLog.New(file, logPrefix, goLog.Ldate|goLog.Ltime)

	warnPath := filepath.Join(logPath, time.Now().Format(warnLogFileName))
	file, err = os.Create(warnPath)
	if err != nil {
		goLog.Fatalf("create log file error[%v]", err)
	}
	warnLogger = goLog.New(file, logPrefix, goLog.Ldate|goLog.Ltime)

	errorPath := filepath.Join(logPath, time.Now().Format(errorLogFileName))
	file, err = os.Create(errorPath)
	if err != nil {
		goLog.Fatalf("create log file error[%v]", err)
	}
	errorLogger = goLog.New(file, logPrefix, goLog.Ldate|goLog.Ltime)

}

func InfoF(format string, v ...interface{}) {
	infoLogger.Printf(format, v)
}

func WarnF(format string, v ...interface{}) {
	warnLogger.Printf(format, v)
}

func ErrorF(format string, v ...interface{}) {
	errorLogger.Fatalf(format, v)
}
