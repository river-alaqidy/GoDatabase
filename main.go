package main

import (
	"fmt"
	"math/rand"
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

// save and write to files with atomicity
func SaveData(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer func() { // discard the temp file if it still exists
		fp.Close() // not expected to fail
		if err != nil {
			os.Remove(tmp)
		}
	}()

	if _, err = fp.Write(data); err != nil { // save to temporary file
		return err
	}
	if err = fp.Sync(); err != nil { // fsync
		return err
	}
	err = os.Rename(tmp, path)
	return fp.Sync() // fsync is the syscall that makes all previously written data durable is used to request and confirma durability
}

func main() {

}
