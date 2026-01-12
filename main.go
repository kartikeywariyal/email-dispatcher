package main

import (
	"sync"
)

var i int = 1

type Recipient struct {
	Name  string
	Email string
}

func main() {
	filePath := "./dummy_users.csv"
	var wg sync.WaitGroup
	RecipientChannel := make(chan Recipient)
	go loadRecipients(filePath, RecipientChannel)
	workerCount := 5
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go processRecipients(i, RecipientChannel, &wg)
	}

	wg.Wait()
}
