package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type ImportCmd struct {
}

func readLines(filename string, ch chan string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')

		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
		}

		ch <- line

		if err != nil && err != io.EOF {
			return err
		}
	}

	return err
}

func parseLine(c chan string) {
	for s := range c {
		s = strings.Replace(s, ");", "", 1)
		s = strings.Replace(s, "'", "", -1)
		parts := strings.Split(s, ",")

		email := strings.Trim(strings.ToLower(parts[2]), "")
		pass := parts[3]

		fmt.Printf("Email %s Password %s\n", email, pass)
	}
}

func (c ImportCmd) Run(ctx context.Context) error {
	fmt.Println("Hello from import command")

	res := make(chan string, 100)

	go parseLine(res)

	filename := "/media/anboo/f3cda4e8-5110-499f-80e0-dea951a38b2b/VK_100M.txt"
	err := readLines(filename, res)

	if err != nil {
		return err
	}

	return nil
}

func (c ImportCmd) Name() string {
	return "import"
}
