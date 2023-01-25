package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/JackDaniells/activity-selector/pkg/logger"
)

func readFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	return csvReader.ReadAll()
}

func main() {
	logger.Debug().Log("Initializing application")

	// get filename from the parameters
	filename := ""
	if filename == "" && len(os.Args) < 2 {
		panic("You must pass name of the file to be processed as a parameter")
	}
	filename = strings.TrimSpace(os.Args[1])

	// read csv file
	logger.Debug().Log("Reading file...")
	file, err := readFile(filename)
	if err != nil {
		logger.Error().Log(err.Error())
		return
	}

	// generate a random number and select a activity based on that
	logger.Debug().Log("Selecting activity from the list...")
	rand.Seed(time.Now().UnixNano())
	row := file[rand.Intn(len(file))]
	logger.Info().Log(fmt.Sprintf("Activity to be executed: %s", row))

	logger.Debug().Log("Finishing application")
}
