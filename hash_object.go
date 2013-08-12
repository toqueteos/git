package git

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func HashObject(name string, write bool) (string, error) {
	// var content, header, store string
	var content string

	// Try to read file
	_, err := os.Stat(name)
	if err != nil || os.IsNotExist(err) { // It can be PathError
		content = name
	} else {
		body, _ := ioutil.ReadFile(name)
		content = string(body)
	}
	// BUG(toqueteos): Hardcodeed `core.autocrlf=true`. Should use git config.
	content = crlf_to_git(content)

	// Either it's a file or just plain text from stdin compute SHA-1 checksum
	header := fmt.Sprintf("%s %d\x00", Blob, len(content))
	store := header + content
	hash := sha1.New()
	_, err = io.Copy(hash, bytes.NewBufferString(store))
	if err != nil {
		return "", err
	}

	// Get hex digest
	checksum := fmt.Sprintf("%x", hash.Sum(nil))

	// Create directory
	if write {
		os.MkdirAll(filepath.Join(".git", "objects", checksum[:2]), 0666)
		// Create file inside that directory
		object, err := os.Create(filepath.Join(".git", "objects", checksum[:2], checksum[2:]))
		if err != nil {
			return "", err
		}
		defer object.Close()

		w := zlib.NewWriter(object)
		defer w.Close()

		_, err = io.Copy(w, bytes.NewBufferString(store))
		if err != nil {
			return "", err
		}
	}

	return checksum, nil
}
