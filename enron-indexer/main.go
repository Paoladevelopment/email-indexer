package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/Paoladevelopment/enron-indexer/models"
	"github.com/Paoladevelopment/enron-indexer/utils"
	zincutilities "github.com/Paoladevelopment/enron-indexer/zincUtilities"
)

func main() {

	// Start profiling
	cpuFile, err := os.Create("indexer1.pprof")
	if err != nil {
		fmt.Printf("Ow, there was an error creating cpu profile file: %v\n", err)
	}
	defer cpuFile.Close()

	//Start cpu profile
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Printf("Error starting cpu profile: %v\n", err)
	}
	defer pprof.StopCPUProfile()

	emailsIndex := models.NewEmailIndex()
	zincutilities.CreateIndex(emailsIndex)
	path := "./enron_mail_20110402"
	utils.IndexEmails(path)
	fmt.Println("Emails indexed!!")

	//Start memory profile

	memFile, err := os.Create("mem_indexer1.prof")
	if err != nil {
		fmt.Printf("Ow, there was an error creating mem profile file: %v\n", err)
	}
	defer memFile.Close()

	err = pprof.WriteHeapProfile(memFile)

	if err != nil {
		fmt.Printf("Error writing memory profile: %v\n", err)
	}
}
