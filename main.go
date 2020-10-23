package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var version = "1.1"

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

func submit(file string) int {
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

// IsFileExist check if a file exists
// Ref: <https://studygolang.com/topics/20>
func IsFileExist(fileName string) (error, bool) {
	_, err := os.Stat(fileName)
	if err == nil {
		return nil, true
	}
	if os.IsNotExist(err) {
		return nil, false
	}
	return err, false
}

// GenCallPBS check and generate a PBS file
func GenCallPBS(prefix string) string {
	number := 1
	fileName := fmt.Sprintf("./%s%d.pbs", prefix, number)
	_, exists := IsFileExist(fileName)
	for exists {
		log.Printf("File %s exists, trying to set another name.", fileName)
		number = number + 1
		fileName = fmt.Sprintf("./%s%d.pbs", prefix, number)
		_, exists = IsFileExist(fileName)
	}

	log.Printf("Generating file %s.", fileName)
	_, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Created file %s.", fileName)
	}
	return fileName
}

func main() {
	pPtr := flag.Bool("p", false, "enable parallel processing.")
	nodePtr := flag.Int("nodes", 1, "an int to specify node number to use. Only work when -p enabled.")
	ppnPtr := flag.Int("ppn", 1, "an int to specify cpu number per node. Only work when -p enabled.")
	jobPtr := flag.Int("jobs", 0, "run n jobs in parallel, at default will use nodes*ppn.")
	outPtr := flag.String("name", "pwork", "an file prefix for generating output PBS file. Only work when -p enabled.")

	flag.Parse()
	inPath := flag.Args()

	//fmt.Println(*pPtr, "-", *nodePtr, "-", *ppnPtr, "-", *outPtr, "-", inPath)
	nodes := *nodePtr
	ppns := *ppnPtr
	jobs := *jobPtr

	if len(inPath) != 1 {
		log.Fatalf("Only one directory path is allowed!")
	}

	log.Printf("gosub version: %s\n", version)
	log.Println("Submitted file list will be save to success_submitted_list.txt!")
	log.Println("====================================")

	// Remove previous file
	if _, exists := IsFileExist("./success_submitted_list.txt"); exists {
		log.Println("Previous file success_submitted_list.txt detected, removing it...")
		cmd := exec.Command("sh", "-c", "rm ./success_submitted_list.txt")
		_, err := cmd.CombinedOutput()

		if err != nil {
			log.Fatalf("Remove previous file error: please contact Shixiang.")
		}
	}

	// List all PBS files
	var files []string
	err := filepath.Walk(inPath[0], visit(&files, ".pbs"))
	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		log.Fatalf("No pbs files found in directory %s!", inPath[0])
	}

	if *pPtr {
		// Run parallel mode
		log.Println("Parallel mode is enabled.")
		log.Println("====================================")
		info := fmt.Sprintf("Use %d threads: %d CPUs per Node.", nodes*ppns, *ppnPtr)
		log.Println(info)
		pbs := GenCallPBS(*outPtr)

		// Use the first file as template
		// Generate header
		cmd1 := fmt.Sprintf("cat %s | grep '#PBS' | grep -v nodes | grep -v '#PBS -N' >> %s", files[0], pbs)
		cmd2 := "echo '#PBS -N gosub_parallel_work' >> " + pbs
		cmd3 := fmt.Sprintf("echo '#PBS -l nodes=%d:ppn=%d' >> %s", nodes, ppns, pbs)

		cmd := exec.Command("sh", "-c", cmd1)
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		cmd = exec.Command("sh", "-c", cmd2)
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		cmd = exec.Command("sh", "-c", cmd3)
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		// Write parallel computation commands to generated file
		// Doc: https://github.com/shenwei356/rush
		totalJobs := 0
		if jobs == 0 {
			totalJobs = nodes * ppns
		} else {
			totalJobs = jobs
		}
		fileStr := strings.Join(files, " ")
		log.Printf("Joined file list with spaces.")
		log.Println(fileStr)
		cmdP := fmt.Sprintf("echo \"echo %s | rush -D ' ' 'bash {}' -j %d\" >> %s", fileStr, totalJobs, pbs)
		cmd = exec.Command("sh", "-c", cmdP)
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("NOTE the 'rush' command should be available in PATH.")
		submit(pbs)

	} else {
		// Submit PBS one by one
		for _, file := range files {
			submit(file)
		}
	}

	log.Println("====================================")
	log.Println("End.")
}
