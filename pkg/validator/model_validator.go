package validator

import "github.com/charmbracelet/log"

type ModelValidator struct{}

func (m *ModelValidator) ValidateV1UserInformation(user *v1UserInfo) bool {
	if len(user.ID) == 0 {
		log.Debug("ID is empty, dropping record")
		return false
	}
	if len(user.Name) == 0 {
		log.Debug("Name is empty, dropping record")
		return false
	}
	if len(user.Email) == 0 {
		log.Debug("Email is empty, dropping record")
		return false
	}

	return true
}

func (m *ModelValidator) ValidateV2UserInformation(user *v2UserInfo) bool {
    if len(user.ID) == 0 {
        log.Debug("ID is empty, dropping record")
        return false
    }
	return true
}
