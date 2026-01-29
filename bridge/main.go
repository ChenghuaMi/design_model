package main

import "fmt"

//打印机

type Printer interface {
	PrintFile(file string) //打印文件
}
type Hp struct {
}

func (h *Hp) PrintFile(file string) {
	fmt.Println("惠普打印机")
}

//电脑

type Computer interface {
	Print(file string) string   //打印
	SetPrinter(printer Printer) //设置打印机
}
type Mac struct {
	printer Printer
}

func (m *Mac) Print(file string) string {
	m.printer.PrintFile(file)
	return ""
}
func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

func main() {
	mac := &Mac{}
	hp := &Hp{}
	mac.SetPrinter(hp)
	mac.Print("xxxx")
}
