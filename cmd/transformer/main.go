package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	h "transformer/internal/helpers"
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
	path, interval := h.Dataset(string(os.Args[1]))

	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovered in f: %v", r)
		}
	}()

	err := mockStream(path, interval)
	if err != nil {
		log.Errorf("Error in mock stream: %v", err)
	}
}

func mockStream(path string, interval time.Duration) error {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file", err)
		return errors.New("error reading file in mock stream")
	}

	var streamData []v1UserInfo
	err = json.Unmarshal(data, &streamData)
	if err != nil {
		log.Error("Error unmarshalling data", err)
		return errors.New("error unmarshalling data in mock stream")
	}

	err = simulateKinesisStream(streamData, interval)
	if err != nil {
		panic("Error in simulate kinesis stream")
	}

	return nil
}

func simulateKinesisStream(records []v1UserInfo, interval time.Duration) error {
	log.Info("Validating user information")
	// for { // this will make it loop for ever!
	for i, record := range records {
		validateAndMap(record, i)

		time.Sleep(interval)
	}
	// }
	return nil
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
		h.WriteRecordToFile(*data) // uncomment this line to write to file

	} else {
		log.Error(record, "Record is invalid and will be dropped: %+v", i, record)
	}
}
