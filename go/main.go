package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("answers.json")
	if err != nil {
		log.Fatal(err)
	}

	var answers map[string]int
	err = json.Unmarshal(file, &answers)

	files, err := filepath.Glob("problems/*.go")
	if err != nil {
		log.Fatal(err)
	}

	for _, filename := range files {
		cmd := exec.Command("go", "run", filename)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		num := regexp.MustCompile(`\d+`)
		i := num.Find([]byte(filename))

		attempt, _ := strconv.Atoi(strings.TrimSpace(out.String()))
		expected := answers[string(i)]

		if attempt == expected {
			fmt.Printf("%v ✓\n", filename)
		} else {
			fmt.Printf("** FAIL ** %v: attempt=%v, expected=%v\n", filename, attempt, expected)

		}
	}
}
