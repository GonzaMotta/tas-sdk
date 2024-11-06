package sdkImed

import "log"

func HelloWord() string {
	return "Hello World"
}

func HelloYou(name string) {
	log.Printf("Hello %s", name)
}
