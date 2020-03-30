package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

// init() is called before main()
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./gosub <dir_to_pbs_files>")
		os.Exit(-1)
	}
}

// Source: https://flaviocopes.com/go-list-files/
// NOTE: subdirectories will also be visited
func visit(files *[]string, ext string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if filepath.Ext(path) == ext {
			*files = append(*files, path)
		}
		return nil
	}
}

func submit (file string) int {
	log.Printf("Submitting %s\n", file)
	cmd := exec.Command("qsub", file)
	_, err := cmd.CombinedOutput()
	if err != nil {
		if match, _ := regexp.MatchString("queue limit", err.Error()); match {
			log.Println("Found queue limit for submitting job!")
			log.Println("Waiting for 5 minutes..")
			time.Sleep(5 * time.Minute)
			log.Println("Calling back to submit...")
			return submit(file)
		} else {
			log.Fatalf("Submitting %s failed with error:\n[%s]\n", file, err)
		}
	}

	return 0
}


func main() {
	var files []string

	err := filepath.Walk(os.Args[1], visit(&files, ".pbs"))
	if err != nil {
		log.Panic(err)
	}

	//for _, file := range files {
	//	fmt.Println(file)
	//}

	if len(files) == 0 {
		log.Fatalf("No pbs files found in directory %s!", os.Args[1])
	}

	// Submit PBS one by one
	for _, file := range files {
		submit(file)
	}
}

