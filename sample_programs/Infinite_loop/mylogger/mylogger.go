package mylogger

import "log"

func ListenForLogs(ch chan string) {
	for {
		msg := <-ch
		if msg == "" {
			continue
		}
		log.Println(msg)
	}
}


