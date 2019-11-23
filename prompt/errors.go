package prompt

import (
	"log"
)

func loged(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
