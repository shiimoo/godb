package mlog

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestLogLv(t *testing.T) {
	log.Println(Debug.String(), 1)
	log.Println(Info.String(), 1)
	log.Println(Warn.String(), 1)
	log.Println(Error.String(), 1)
	log.Println(Fatal.String(), 1)
}

func TestLogMsg(t *testing.T) {

	logger := newLogger(context.Background(), "default")
	// logger.
	logger.Start()
	logger.Output(Info, []string{"server", "test"}, "log Test")
	logger.Outputf(Info, []string{"server", "test"}, "log Test %s ", Data{time.Now(), "time"})
	time.Sleep(10 * time.Second)
}

func TestLogMgr(t *testing.T) {
	root := context.Background()
	MgrInit(root)
	GetMgr().SetOutFunc("login", func(l *Log) error {
		log.Println(l)
		return nil
	})
	GetMgr().Log("love")
	time.Sleep(1 * time.Second)
}