package core

import "errors"

func getType(te uint8) uint8 {
	return (te >> 4) << 4
}

func getEncoding(te uint8) uint8 {
	return te & 0b00001111
}