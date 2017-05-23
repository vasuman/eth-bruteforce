package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

var (
	ksPath = flag.String("keystore-dir", "", "path to keystore dir")
	pwFile = flag.String("password-file", "", "line separated file containing possible passwords")
)

const (
	commentPrefix       = "#"
	escapeCommentPrefix = "\\#"
	lineSep             = "\n"
	addrKey             = "address"
)

type keyFile struct {
	json []byte
	addr string
}

func loadPasswords() []string {
	data, err := ioutil.ReadFile(*pwFile)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), lineSep)
}

func loadKeys() []*keyFile {
	ksFiles, err := ioutil.ReadDir(*ksPath)
	if err != nil {
		panic(err)
	}
	var keys []*keyFile
	for _, ksFile := range ksFiles {
		data, err := ioutil.ReadFile(filepath.Join(*ksPath, ksFile.Name()))
		if err != nil {
			panic(err)
		}
		var m interface{}
		err = json.Unmarshal(data, &m)
		if err != nil {
			panic(err)
		}
		f := m.(map[string]interface{})
		addr := f[addrKey].(string)
		keys = append(keys, &keyFile{json: data, addr: addr})
	}
	return keys
}

func main() {
	flag.Parse()
	passwords := loadPasswords()
	keys := loadKeys()
	for _, key := range keys {
		found := false
		for _, password := range passwords {
			if len(password) == 0 || strings.HasPrefix(password, commentPrefix) {
				continue
			}
			if strings.HasPrefix(password, escapeCommentPrefix) {
				password = password[1:]
			}
			_, err := keystore.DecryptKey(key.json, password)
			if err == nil {
				fmt.Printf("for address %s found password '%s'\n", key.addr, password)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("failed to find password for address %s\n", key.addr)
		}
	}
}
