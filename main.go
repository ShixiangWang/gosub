package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var version = "0.3.1"

// init() is called before main()
func init() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ./gosub <dir_to_pbs_files>\t(version:%s)\n", version)
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
		//if match, _ := regexp.MatchString("queue limit", err.Error()); match {
		//	log.Println("Found queue limit for submitting job!")
		    log.Printf("Submitting %s failed with error:\n", file)
		    log.Println(err)
			log.Println("Waiting for 5 minutes..")
			time.Sleep(5 * time.Minute)
			log.Println("Calling back to submit...")
			return submit(file)
		//} else {
		//	log.Fatalf("Submitting %s failed with error:\n[%s]\n", file, err)
		//}
	}

	c := fmt.Sprintf("echo %s >> ./success_submitted_list.txt", file)
	cmd = exec.Command("sh", "-c", c)
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Print file list error: please contact Shixiang.")
	}
	return 0
}


func main() {
	var files []string

	fmt.Printf("gosub version: %s\n", version)
	fmt.Println("Starting...")
	fmt.Println("Submitted file list will be")
	fmt.Println("  save to success_submitted_list.txt")
	fmt.Println("====================================")
	// Remove previous file
	cmd := exec.Command("sh", "-c", "rm ./success_submitted_list.txt")
	_, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Remove file error: please contact Shixiang.")
	}

	err = filepath.Walk(os.Args[1], visit(&files, ".pbs"))
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

	fmt.Println("====================================")
	fmt.Println("End.")
}

