package main

import "log"

func strNotEmpty(s string, onerr string) {
	if len(s) < 1 {
		log.Fatal("\n\tError: an empty string!\n\t" + onerr)
	}
}
