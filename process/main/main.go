package main

import (
	"CHN-Administrative-Divisions/process"
	"time"
)

func main() {
	var bar process.Bar
	bar.NewOption(0, 41536)
	for i := 0; i <= 41536; i += 20 {
		bar.Play(int64(i))
		time.Sleep(time.Millisecond * 100)
	}
	bar.Finish()
}
