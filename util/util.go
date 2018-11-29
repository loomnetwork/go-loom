package util

import (
	"fmt"
	"os"
	"syscall"
)

func PrefixKey(keys ...[]byte) []byte {
	size := len(keys) - 1
	for _, key := range keys {
		size += len(key)
	}
	buf := make([]byte, 0, size)

	for i, key := range keys {
		if i > 0 {
			buf = append(buf, 0)
		}
		buf = append(buf, key...)
	}
	return buf
}

func UnprefixKey(key []byte, prefixes ...[]byte) ([]byte, error) {
	for _, prefix := range prefixes {
		if len(prefix)+1 > len(key) {
			return []byte{}, fmt.Errorf("prefix %s longer than key %s", string(prefix), string(key))
		}
		key = key[len(prefix)+1:]
	}
	return key, nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func IgnoreErrNotExists(err error) error {
	if perr, ok := err.(*os.PathError); ok {
		// On Windows the error is actually syscall.ERROR_FILE_NOT_FOUND (set to syscall.Errno(2)),
		// and this is also the case in WSL (Windows Subsystem for Linux) except that in that case
		// syscall.ERROR_FILE_NOT_FOUND is undefined. To ensure this error is caught in both cases
		// have to check for syscall.Errno(2).
		if perr.Err == os.ErrNotExist || perr.Err == syscall.Errno(2) {
			return nil
		}
	}
	return err
}
