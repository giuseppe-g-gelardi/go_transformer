package mapper

import (
	"errors"
	"strconv"
	"strings"

	"transformer/pkg/types"
	"transformer/pkg/validator"

	"github.com/charmbracelet/log"
)

type (
	v1UserInfo = types.V1UserInformation
	v2UserInfo = types.V2UserInformation
)

var v validator.ModelValidator

type Mapper struct{}

func (t *Mapper) MapV2Schema(v1Data v1UserInfo) (*v2UserInfo, error) {
	var v2Data v2UserInfo
	var firstName, lastName string

	firstName, lastName, err := parseUserName(v1Data)
	if err != nil {
		log.Error("Error parsing name")
		return nil, errors.New("error: something went wrong parsing name")
	}

	street, city, state, zip, err := parseAddress(v1Data)
	if err != nil {
		log.Error("Error parsing address")
		return nil, errors.New("error: something went wrong parsing address")
	}

	log.Debug("Mapping V1 to V2")
    // there has to be a prettier way to do this
	v2Data.ID = v1Data.ID
	v2Data.AccountInformation.IsActive = v1Data.IsActive
	v2Data.AccountInformation.Registered = v1Data.Registered
	v2Data.AccountInformation.Balance = v1Data.Balance
	v2Data.UserInformation.FirstName = firstName
	v2Data.UserInformation.LastName = lastName
	v2Data.UserInformation.Age = v1Data.Age
	v2Data.UserInformation.Gender = v1Data.Gender
	v2Data.UserInformation.EyeColor = v1Data.EyeColor
	v2Data.UserInformation.Picture = v1Data.Picture
	v2Data.UserInformation.Company = v1Data.Company
	v2Data.ContactInformation.Email = v1Data.Email
	v2Data.ContactInformation.Phone = v1Data.Phone
	v2Data.ContactInformation.Address.Street = street
	v2Data.ContactInformation.Address.City = city
	v2Data.ContactInformation.Address.State = state
	v2Data.ContactInformation.Address.Zip = zip
	v2Data.Tags = v1Data.Tags
	v2Data.Profile = v1Data.About

	isV2Valid := v.ValidateV2UserInformation(&v2Data)
	if !isV2Valid {
		log.Error("V2 data is invalid")
		return nil, errors.New("error: something went wrong mapping to new schema")
	}

	return &v2Data, nil
}

func parseAddress(v1Data v1UserInfo) (string, string, string, int, error) {
	var street, city, state, zip string
	address := strings.Split(v1Data.Address, ",")
	if len(address) == 4 {
		street = address[0]
		city = address[1]
		state = address[2]
		zip = address[3]
	} else {
		log.Error("Address is not in the correct format")
	}

	zipInt, err := strconv.Atoi(strings.TrimSpace(zip))
	if err != nil {
		log.Error("Error converting zip to int")
		return "", "", "", 0, errors.New("error: something went wrong converting zip to int")
	}

	return street, city, state, zipInt, nil
}

func parseUserName(v1Data v1UserInfo) (string, string, error) {
	var firstName, lastName string
	splitName := strings.Split(v1Data.Name, " ")
	if len(splitName) == 2 {
		firstName = splitName[0]
		lastName = splitName[1]
	} else {
		log.Error("Name is not in the correct format")
	}

	return firstName, lastName, nil
}
