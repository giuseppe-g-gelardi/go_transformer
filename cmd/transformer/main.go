package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"transformer/pkg/mapper"
	"transformer/pkg/types"
	"transformer/pkg/validator"

	"github.com/charmbracelet/log"
)

type (
	v1UserInfo = types.V1UserInformation
	v2UserInfo = types.V2UserInformation
)

var (
	v validator.ModelValidator
	m mapper.Mapper
)

func main() {
	MockStream()
}

func writeRecordToFile(record v2UserInfo) error {
	file, err := os.OpenFile("./mock_data/output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Error("Error opening file", err)
		return errors.New("error opening file")
	}
	defer file.Close()

	// recordJSON, err := json.Marshal(record)
	recordJSON, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		log.Errorf("Error marshalling record: %v", err)
		return errors.New("error marshalling record")
	}

	if _, err := file.Write(append(recordJSON, '\n')); err != nil {
		log.Error("Error writing record to file", err)
		return errors.New("error writing record to file")
	}

	return nil
}

func MockStream() {
	path, interval := dataset(string(os.Args[1]))

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var streamData []v1UserInfo
	err = json.Unmarshal(data, &streamData)
	if err != nil {
		log.Error("Error unmarshalling data", err)
		return
	}

	simulateKinesisStream(streamData, interval)
}

func simulateKinesisStream(records []v1UserInfo, interval time.Duration) {
	log.Info("Validating user information")
	// for { // this will make it loop for ever!
	for i, record := range records {
		validateAndMap(record, i)

		time.Sleep(interval)
	}
	// }
}

func validateAndMap(record v1UserInfo, i int) {
	log.Infof("Processing record: #%d", i)
	log.Printf("Old record: %+v", record)
	valid := v.ValidateV1UserInformation(&record)
	if valid {
		data, err := m.MapV2Schema(record)
		if err != nil {
			log.Errorf("Error mapping record #%d: %v", i, err)
			return
		}
		log.Printf("New record: %+v", data)
		log.Infof("Record #%d processed successfully\n\n", i)
		// writeRecordToFile(*data) // uncomment this line to write to file

	} else {
		log.Error(record, "Record is invalid and will be dropped: %+v", i, record)
	}
}

func dataset(size string) (string, time.Duration) {
	small := "./mock_data/small.json"
	medium := "./mock_data/medium.json"
	large := "./mock_data/large.json"

	switch size {
	case "small":
		return small, 500 * time.Millisecond
	case "medium":
		return medium, 100 * time.Millisecond
	case "large":
		return large, 10 * time.Millisecond
	default:
		return small, 500 * time.Millisecond
	}
}
