package helpers

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"transformer/pkg/types"

	"github.com/charmbracelet/log"
)

type v2UserInfo = types.V2UserInformation



func WriteRecordToFile(record v2UserInfo) error {
	file, err := os.OpenFile("./mock_data/output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Error("Error opening file", err)
		return errors.New("error opening file")
	}
	defer file.Close()

	// recordJSON, err := json.Marshal(record) // single line for each record
	recordJSON, err := json.MarshalIndent(record, "", "  ") // pretty print
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

func Dataset(size string) (string, time.Duration) {
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
