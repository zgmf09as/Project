package main

import (
	"fmt"
	"runtime"
	"time"
	"log"
)

type LogAnalysisService struct {
	id     int64
	cmdCh  chan GoCmd
}

func NewService() *LogAnalysisService {

	g := LogAnalysisService{
		id:     <- IdGenCh,
		cmdCh:  make(chan GoCmd, 10),
	}

	go g.Loop()
	return &g
}

func (s *LogAnalysisService) ID() int64 {
	return s.id
}

func (s *LogAnalysisService) SendGoCmd(Cmd string, Args interface{}, Rsp chan<- interface{}) {
	s.cmdCh <- GoCmd{
		Cmd: Cmd,
		Args: Args,
		Rsp: Rsp,
	}
}

func (s LogAnalysisService) String() string {
	return fmt.Sprintf("LogAnalysisService%v goroutine:%v IDs:%v",
		s.ID, runtime.NumGoroutine(), <-IdGenCh)
}

func (s *LogAnalysisService) Loop(){
	fps := 60
	timer60Ch := time.Tick(time.Duration(1000/fps) * time.Millisecond)
	timer1secCh := time.Tick(1 * time.Second)
loop:
	for {
		select {
		case cmd := <-s.cmdCh:
			//log.Println(cmd)
			switch cmd.Cmd {
			default:
				log.Printf("unknown cmd %v", cmd)
			case "quit":
				break loop
			}
		case <-timer60Ch:
			// do frame action
		case <-timer1secCh:
		}
	}
}