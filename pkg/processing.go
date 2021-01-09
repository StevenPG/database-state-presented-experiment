package pkg

import (
	"errors"
	"log"
)

// PutDatabaseItem performs verification of inbound requests against database state
func PutDatabaseItem(item *Item) error {

	// get all database items
	dbValueSlice := GetBoltContent()

	if len(dbValueSlice) != len(item.State) {
		log.Print("Validating inbound state, length of request and database different.\n")
		return errors.New("length of request and database different")
	}

	// unravel request
	for _, presentedValue := range item.State {
		// verify they match perfectly
		accurateState := contains(dbValueSlice, presentedValue.Value)
		if accurateState == false {
			log.Printf("Validating inbound state, failed verification:\n%s", item.State)
			return errors.New("failed verification")
		}
	}

	WriteItem(item)
	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
