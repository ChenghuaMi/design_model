package main

import "fmt"

type Pay interface {
	PayPage(price int64) string
}
type PayType int

const (
	AlipayType PayType = 1
	WxpayType  PayType = 2
)

type Alipay struct {
}

func (a *Alipay) PayPage(price int64) string {
	return "alipay"
}

type WxPay struct {
}

func (a *WxPay) PayPage(price int64) string {
	return "wx"
}
func NewPayPage(payType PayType) Pay {
	switch payType {
	case AlipayType:
		return &Alipay{}
	case WxpayType:
		return &WxPay{}
	}
	return nil
}
func main() {
	fmt.Println(NewPayPage(AlipayType).PayPage(1))
	fmt.Println(NewPayPage(WxpayType).PayPage(2))
}
