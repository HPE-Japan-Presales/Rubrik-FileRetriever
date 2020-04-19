package util

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logging struct {
	LogDir      string
	LogFileName string
}

var writer io.Writer = os.Stdout

func (l *Logging) CreateLogFile(log_file_name string) {
	path := l.LogDir + "/" + log_file_name
	_, err := os.Stat(path)
	if err != nil {
		_, err := os.Create(path)
		if err != nil {
			l.Fatalf("Could not create log file: %s", path)
		}
		l.Infof("Created log file: %s", path)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Fatalf("Failed to open log file: %s", path)
	}

	l.LogFileName = log_file_name
	writer = io.MultiWriter(file, os.Stdout)
}

func (l *Logging) Info(v ...interface{}) {
	log.New(writer, "[Info] ", log.LstdFlags).Print(fmt.Sprint(v...))
}

func (l *Logging) Infoln(v ...interface{}) {
	log.New(writer, "[Info] ", log.LstdFlags).Println(fmt.Sprint(v...))
}

func (l *Logging) Infof(format string, v ...interface{}) {
	log.New(writer, "[Info] ", log.LstdFlags).Printf(fmt.Sprintf(format, v...))
}

func (l *Logging) Error(v ...interface{}) {
	log.New(writer, "[Error] ", log.LstdFlags).Print(fmt.Sprint(v...))
}

func (l *Logging) Errorln(v ...interface{}) {
	log.New(writer, "[Error] ", log.LstdFlags).Println(fmt.Sprint(v...))
}

func (l *Logging) Errorf(format string, v ...interface{}) {
	log.New(writer, "[Error] ", log.LstdFlags).Printf(fmt.Sprintf(format, v...))
}

func (l *Logging) Warn(v ...interface{}) {
	log.New(writer, "[Warning] ", log.LstdFlags).Print(fmt.Sprint(v...))
}

func (l *Logging) Warnln(v ...interface{}) {
	log.New(writer, "[Warning] ", log.LstdFlags).Println(fmt.Sprint(v...))
}

func (l *Logging) Warnf(format string, v ...interface{}) {
	log.New(writer, "[Warning] ", log.LstdFlags).Printf(fmt.Sprintf(format, v...))
}

func (l *Logging) Fatal(v ...interface{}) {
	log.New(writer, "[Fatal] ", log.LstdFlags).Fatal(fmt.Sprint(v...))
}

func (l *Logging) Fatalln(v ...interface{}) {
	log.New(writer, "[Fatal] ", log.LstdFlags).Fatalln(fmt.Sprint(v...))
}

func (l *Logging) Fatalf(format string, v ...interface{}) {
	log.New(writer, "[Fatal] ", log.LstdFlags).Fatalf(fmt.Sprintf(format, v...))
}
