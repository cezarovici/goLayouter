package apperrors

import (
	"fmt"
	"time"
)

type ServiceError struct {
	Caller              string
	MethodName          string
	Issue               error
	WithTime            bool
	NanosecondsDuration int64
}

const areaErrService = "Service"

func (e *ServiceError) Error() string {
	var res [4]string

	if e.WithTime {
		res[0] = fmt.Sprintf("\nArea: %s [%d] - duration nanoseconds: %d", areaErrService,
			time.Now().Unix(), e.NanosecondsDuration)
	} else {
		res[0] = fmt.Sprintf("\nArea: %s", areaErrService)
	}

	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Method Name: %s", e.MethodName)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}

type DomainError struct {
	Caller     string
	MethodName string
	Issue      error

	WithTime bool
}

const areaErrDomain = "Domain"

func (e *DomainError) Error() string {
	var res [4]string

	if e.WithTime {
		res[0] = fmt.Sprintf("\nArea: %s [%d]", areaErrDomain, time.Now().Unix())
	} else {
		res[0] = fmt.Sprintf("\nArea: %s", areaErrDomain)
	}

	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Method Name: %s", e.MethodName)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}

type RenderError struct {
	Caller     string
	MethodName string
	Issue      error

	NanosecondsDuration int64
	WithTime            bool
}

const areaErrRender = "Render"

func (e *RenderError) Error() string {
	var res [4]string

	if e.WithTime {
		res[0] = fmt.Sprintf("\nArea: %s [%d] - duration nanoseconds: %d",
			areaErrRender, time.Now().Unix(), e.NanosecondsDuration)
	} else {
		res[0] = fmt.Sprintf("\nArea: %s", areaErrRender)
	}

	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Method Name: %s", e.MethodName)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}
