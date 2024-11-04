package validator

import "transformer/pkg/types"


type (
    v1UserInfo = types.V1UserInformation
    v2UserInfo = types.V2UserInformation
)

type Validator interface {
	ValidateV1UserInformation(v1UserInformation *v1UserInfo) bool
	ValidateV2UserInformation(v2UserInformation *v2UserInfo) bool
    // ValidateUserID(id int64) bool
}
