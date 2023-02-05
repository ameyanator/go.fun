package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"goinpractice.com/GoInPractice/main/technique56/file"
)

func main() {
	content := `Lorem ipsum dolor sit amet, consectetur` +
		`adipiscing elit. Donec a diam lectus.Sed sit` +
		`amet ipsum mauris. Maecenascongue ligula ac` +
		`quam viverra nec consectetur ante hendrerit.`

	body := bytes.NewReader([]byte(content))
	store, err := fileStore()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Storing content...")
	err = store.Save("foo/bar", body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Retriving Content...")
	c, err := store.Load("foo/bar")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	o, err := ioutil.ReadAll(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(o))
}

func fileStore() (file.File, error) {
	return &file.LocalFile{Base: "."}, nil
}
