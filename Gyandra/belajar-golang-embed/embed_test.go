package belajar_golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
	"io/ioutil"
	"io/fs"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.jpg
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

