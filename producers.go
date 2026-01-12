package main

import (
	"encoding/csv"
	"os"
)

func loadRecipients(filePath string, RecipientChannel chan Recipient) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return
	}

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return
	}

	defer func() {
		close(RecipientChannel)
		f.Close()
	}()
	for _, record := range records {

		RecipientChannel <- Recipient{
			Name:  record[0],
			Email: record[1],
		}

	}

}
