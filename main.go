package main

import (
	"os"
)

/*
simple key - value interface

	get, set, delete a single key
	list a rnage of keys in sorted order
*/
type KV interface {
	Get(key []byte) (val []byte, ok bool)
	Set(key []byte) (val []byte)
	Del(key []byte)

	FindGreaterThan(key []byte) Iterator
}

type Iterator interface {
	HasNext() bool
	Next() (key []byte, val []byte)
}

// typical way to save and write files
func SaveData(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close() // ensure file is closed after execution

	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	return fp.Sync() // fsync is the syscall that makes all previously written data durable is used to request and confirma durability
}

func main() {

}
