package main

import (
	"fmt"
	"time"
)

type ReqI interface {
	Handler(url string) string
}

type Req struct {
}

func (r *Req) Handler(url string) string {
	fmt.Println("req:" + url)
	time.Sleep(100 * time.Millisecond)
	return ""
}

type LogDecorate struct {
	req ReqI
}

func (l *LogDecorate) Handler(url string) string {
	fmt.Println("before...")
	l.req.Handler(url)
	fmt.Println("after...")
	return ""
}

type MonitorDecorate struct {
	req ReqI
}

func (m *MonitorDecorate) Handler(url string) string {
	fmt.Println("monitor before..")
	m.req.Handler(url)
	fmt.Println("monitor after")
	return ""
}
func main() {
	req := Req{}
	l := &LogDecorate{req: &req}
	m := &MonitorDecorate{req: l}
	m.Handler("xxx.com")
}
