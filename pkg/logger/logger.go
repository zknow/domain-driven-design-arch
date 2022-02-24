package logger

import (
	"io"
	"os"
	"path"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var (
	fileDir  = "/data/log/"
	fileName = "log.txt"
	MultiIO  = false
)

func init() {
	if MultiIO {
		f := getFileioWriter()
		multiIO := io.MultiWriter(f, os.Stdout)
		log.SetOutput(multiIO)
	}

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func getFileioWriter() io.Writer {
	filepath := getLogFilePath()

	ensureDir(filepath)

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func getLogFilePath() string {
	cmdDir, _ := os.Getwd()
	return path.Join(cmdDir, fileDir, fileName)
}

func ensureDir(path string) {
	dirName := filepath.Dir(path)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}
}
