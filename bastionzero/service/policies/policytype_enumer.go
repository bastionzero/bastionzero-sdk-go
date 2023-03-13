// Code generated by "enumer -type=PolicyType -json"; DO NOT EDIT.

package policies

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _PolicyTypeName = "TargetConnectOrganizationControlsSessionRecordingKubernetesProxyJustInTime"

var _PolicyTypeIndex = [...]uint8{0, 13, 33, 49, 59, 64, 74}

const _PolicyTypeLowerName = "targetconnectorganizationcontrolssessionrecordingkubernetesproxyjustintime"

func (i PolicyType) String() string {
	if i < 0 || i >= PolicyType(len(_PolicyTypeIndex)-1) {
		return fmt.Sprintf("PolicyType(%d)", i)
	}
	return _PolicyTypeName[_PolicyTypeIndex[i]:_PolicyTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PolicyTypeNoOp() {
	var x [1]struct{}
	_ = x[TargetConnect-(0)]
	_ = x[OrganizationControls-(1)]
	_ = x[SessionRecording-(2)]
	_ = x[Kubernetes-(3)]
	_ = x[Proxy-(4)]
	_ = x[JustInTime-(5)]
}

var _PolicyTypeValues = []PolicyType{TargetConnect, OrganizationControls, SessionRecording, Kubernetes, Proxy, JustInTime}

var _PolicyTypeNameToValueMap = map[string]PolicyType{
	_PolicyTypeName[0:13]:       TargetConnect,
	_PolicyTypeLowerName[0:13]:  TargetConnect,
	_PolicyTypeName[13:33]:      OrganizationControls,
	_PolicyTypeLowerName[13:33]: OrganizationControls,
	_PolicyTypeName[33:49]:      SessionRecording,
	_PolicyTypeLowerName[33:49]: SessionRecording,
	_PolicyTypeName[49:59]:      Kubernetes,
	_PolicyTypeLowerName[49:59]: Kubernetes,
	_PolicyTypeName[59:64]:      Proxy,
	_PolicyTypeLowerName[59:64]: Proxy,
	_PolicyTypeName[64:74]:      JustInTime,
	_PolicyTypeLowerName[64:74]: JustInTime,
}

var _PolicyTypeNames = []string{
	_PolicyTypeName[0:13],
	_PolicyTypeName[13:33],
	_PolicyTypeName[33:49],
	_PolicyTypeName[49:59],
	_PolicyTypeName[59:64],
	_PolicyTypeName[64:74],
}

// PolicyTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PolicyTypeString(s string) (PolicyType, error) {
	if val, ok := _PolicyTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PolicyTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PolicyType values", s)
}

// PolicyTypeValues returns all values of the enum
func PolicyTypeValues() []PolicyType {
	return _PolicyTypeValues
}

// PolicyTypeStrings returns a slice of all String values of the enum
func PolicyTypeStrings() []string {
	strs := make([]string, len(_PolicyTypeNames))
	copy(strs, _PolicyTypeNames)
	return strs
}

// IsAPolicyType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PolicyType) IsAPolicyType() bool {
	for _, v := range _PolicyTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for PolicyType
func (i PolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for PolicyType
func (i *PolicyType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PolicyType should be a string, got %s", data)
	}

	var err error
	*i, err = PolicyTypeString(s)
	return err
}
