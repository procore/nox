package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/softpunks/chroma/quick"
	"github.com/softpunks/goin"
	"github.com/spf13/viper"
)

var tempBufferFileName = os.TempDir() + "/.nox-request-body.tmp.json"

func confirm() bool {
	if override {
		return true
	}

	prompt := promptui.Prompt{
		Label:     "This action is potentially destructive, proceed?",
		IsConfirm: true,
	}
	p, err := prompt.Run()
	if err != nil {
		fmt.Println("Stopping...")
		os.Exit(0)
	}
	if p == "Y" || p == "y" {
		return true
	}
	return false
}

func readFromFile() string {
	if body != "" {
		return body
	}

	if viper.GetBool("silent") {
		return ""
	}

	var body string
	var err error

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		body, err = readFromStdin()
	} else {
		body, err = goin.ReadFromFile(tempBufferFileName)
	}

	if err != nil {
		log.Fatal(err)
	}

	return body

}

func readFromStdin() (string, error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func isError(m string) bool {
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal([]byte(m), &objmap)

	if err != nil {
		log.Fatal(err)
	}
	if _, ok := objmap["error"]; ok {
		return true
	}
	return false

}

func printResponse(m string) {
	if viper.GetBool("silent") && !isError(m) {
		return
	} else if viper.GetBool("pretty") {
		colorTheme := viper.GetString("theme")
		if colorTheme == "" {
			colorTheme = "fruity"
		}
		err := quick.Highlight(os.Stdout, m, "json", "terminal256", colorTheme)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print(m)
	}
}
