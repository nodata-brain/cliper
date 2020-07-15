package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/mitchellh/go-homedir"
)

type FcInfo struct {
	c   string
	s   string
	now string
}

func main() {
	var err error
	var f *os.File
	d := FcInfo{}

	d.now, err = now()
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err = newDir(d.now)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {

		time.Sleep(3 * time.Second)
		now, err := now()
		if err != nil {
			fmt.Println(err)
			return
		}
		if d.now != now {
			d.now = now
			f, err = newDir(now)
		}
		d.s, _ = clipboard.ReadAll()
		if d.c != d.s {
			d.c = d.s
			if strings.Trim(d.c, " ") != "" {
				fmt.Fprintln(f, d.c+"\n")
			}
		}
	}
}

func now() (string, error) {
	p, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return "", err
	}
	t := time.Now().In(p)
	return t.Format("20060102"), nil
}

func newDir(d string) (*os.File, error) {
	hd, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	dir := hd + "/.cliper/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
	f, err := os.OpenFile(dir+d+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return f, nil

}
