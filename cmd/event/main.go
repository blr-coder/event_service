package main

import "log"

func main() {
	if err := runEventApp(); err != nil {
		log.Fatalln(err)
	}
}
