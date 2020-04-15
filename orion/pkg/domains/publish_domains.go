package domains

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type UnpublishedDomains struct {
	Programs  []Program  `json:"programs"`
	Classes   []Class    `json:"classes"`
	Locations []Location `json:"locations"`
	Achieves  []Achieve  `json:"achieves"`
	Semesters []Semester `json:"semesters"`
	Sessions  []Session  `json:"sessions"`
}

type PublishErrorBody struct {
	RowId    uint   `json:"rowId,omitempty"`
	StringId string `json:"stringId,omitempty"`
	Error    error  `json:"error"`
}

func ConcatErrors(errorList []PublishErrorBody) error {
	switch len(errorList) {
	case 0:
		return nil
	case 1:
		return errors.New(getId(errorList[0]) + ": " + errorList[0].Error.Error())
	default:
		var errorString strings.Builder
		errorString.WriteString("one or more programs failed to publish:")

		for _, errorBody := range errorList {
			errorString.WriteString(" " + getId(errorBody) + ": " + errorBody.Error.Error())
		}

		return errors.New(errorString.String())
	}
}

func getId(errorBody PublishErrorBody) string {
	if errorBody.StringId == "" {
		return strconv.Itoa(int(errorBody.RowId))
	} else {
		return errorBody.StringId
	}
}

func Concatenate(initMessage string, errorList []PublishErrorBody, isString bool) error {
	var errorStrings []string
	errorStrings = append(errorStrings, initMessage)
	for _, errorBody := range errorList {
		if isString {
			errorStrings = append(errorStrings, errorBody.StringId+": "+errorBody.Error.Error())
		} else {
			errorStrings = append(errorStrings, fmt.Sprint(errorBody.RowId)+": "+errorBody.Error.Error())
		}
	}
	return errors.New(strings.Join(errorStrings, "\n"))
}
