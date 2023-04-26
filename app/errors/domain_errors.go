package apperrors

import (
	"fmt"
)

type DomainError struct {
	Issue      error
	Caller     string
	MethodName string
}

const areaErrDomain = "Domain"

func (e *DomainError) Error() string {
	var res [4]string

	res[0] = fmt.Sprintf("\nArea: %s", areaErrDomain)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Method Name: %s", e.MethodName)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}
