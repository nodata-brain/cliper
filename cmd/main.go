package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/mitchellh/go-homedir"
)

func main() {
	var c string

	hd, err := homedir.Dir()
	dir := hd + "/.cliper/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
	f, err := os.OpenFile(dir+"clip.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	for {
		time.Sleep(3 * time.Second)
		s, _ := clipboard.ReadAll()
		if c != s {
			c = s
			if strings.Trim(c, " ") != "" {
				fmt.Fprintln(f, c+"\n")
			}
		}
	}
}
