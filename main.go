package main

import "log"

func HelloWord() {
	log.Fatalf("Hello World")
}

func HelloYou(name string) {
	log.Printf("Hello %s", name)
}
