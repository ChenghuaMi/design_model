package main

import "fmt"

//代理模式

type Log struct {
}

func (l *Log) Info(s string) {
	fmt.Println("sss:", s)
}

type ProxyLog struct {
	log *Log
}

func (p *ProxyLog) Info(s string) {
	fmt.Println("访问前")
	fmt.Println(p.log)
	//延迟初始化
	if p.log == nil {
		p.log = &Log{}
	}
	p.log.Info(s)
	fmt.Println("访问后")
}
func main() {
	p := &ProxyLog{}
	p.Info("xxx")
}
