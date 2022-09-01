package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	arg := os.Args[1]
	inputDir := os.Args[2]
	outputDir := os.Args[3]
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total input files: %d\n", len(files))
	if err := os.RemoveAll(outputDir); err != nil {
		log.Fatalln(err)
	}
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		stdin, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", inputDir, file.Name()))
		var (
			stdout bytes.Buffer
			stderr bytes.Buffer
		)
		cmd := exec.Command("java", arg)
		cmd.Stdin = bytes.NewReader(stdin)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}
		if err := ioutil.WriteFile(fmt.Sprintf("%s/%s.out", outputDir, extractFilename(file.Name())), stdout.Bytes(), 0644); err != nil {
			log.Fatalln(err)
		}
	}
}

func extractFilename(name string) string {
	return name[:len(name)-3]
}
