package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func console(out chan<- string, notify <-chan bool) {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please type your command!")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString fail, err = ", err)
			close(out)
		}
		out <- s

		time.Sleep(time.Millisecond * 100)

		select {
		case exit := <-notify:
			if exit {
				fmt.Println("exit console")
				return
			}

		default:
			//fmt.Println("console default")
		}
	}
}
