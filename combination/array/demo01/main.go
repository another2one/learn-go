package demo01

type Currency int8

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

var Symbol = [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
