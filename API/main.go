package main

import "API/serve"

var vouchers map[string]string

func init() {
	vouchers = make(map[string]string)
	vouchers["ILOVEFOODPANDA"] = "5 percent off all restaurants"
	vouchers["IWONAGAME"] = "10 percent off all restaurants"
	vouchers["THISISWONDERFUL"] = "15 percent off all restaurants"
	vouchers["PLAYAGAIN"] = "20 percent off all restaurants"
}

func main() {
	serve.Start(vouchers)
}
