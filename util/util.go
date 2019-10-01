package util

import (
	"bytes"
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

func UnprefixKey(key, prefix []byte) ([]byte, error) {
	if len(prefix)+1 > len(key) {
		return nil, fmt.Errorf("prefix %s longer than key %s", string(prefix), string(key))
	}
	return key[len(prefix)+1:], nil
}

// HasPrefix checks if the given key was prefixed with the given prefix using the PrefixKey function.
func HasPrefix(key, prefix []byte) bool {
	if len(prefix) == 0 {
		return false
	}
	p := append(prefix, byte(0))
	return bytes.HasPrefix(key, p)
}

// Returns the bytes that mark the end of the key range for the given prefix.
func PrefixRangeEnd(prefix []byte) []byte {
	if prefix == nil {
		return nil
	}

	end := make([]byte, len(prefix))
	copy(end, prefix)

	for {
		if end[len(end)-1] != byte(255) {
			end[len(end)-1]++
			break
		} else if len(end) == 1 {
			end = nil
			break
		}
		end = end[:len(end)-1]
	}
	return end
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
