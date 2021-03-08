package main

import (
	"goProgram/concurrency/channels"
	httpBasic "goProgram/httpbasic1"
	"goProgram/logReader"
)

func main() {
	channels.Wait()
	logReader.ReadFileLog()
	httpBasic.StartServer()
}
