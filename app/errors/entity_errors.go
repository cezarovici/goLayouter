package apperrors

import (
	"fmt"
	"time"
)

type ErrService struct {
	Caller     string
	MethodName string
	Issue      error

	NanosecondsDuration int64
	WithTime            bool
}

const areaErrService = "Service"

func (e *ErrService) Error() string {
	var res [4]string

	if e.WithTime {
		res[0] = fmt.Sprintf("\nArea: %s [%d] - duration nanoseconds: %d", areaErrService, time.Now().Unix(), e.NanosecondsDuration)
	} else {
		res[0] = fmt.Sprintf("\nArea: %s", areaErrService)
	}

	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Method Name: %s", e.MethodName)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}

type ErrDomain struct {
	Caller     string
	MethodName string
	Issue      error

	WithTime bool
}

const areaErrDomain = "Domain"

func (e *ErrDomain) Error() string {
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

// Examples

//TODO remove

// func fooService() error {
// 	now := time.Now()

// 	return &ErrService{
// 		Caller:     "fooService",
// 		MethodName: "m1",
// 		Issue:      errors.New("issue 1"),

// 		WithTime:            true,
// 		NanosecondsDuration: int64(time.Since(now).Nanoseconds()),
// 	}
// }

// func fooDomain() error {
// 	return &ErrDomain{
// 		Caller:     "fooDomain",
// 		MethodName: "fooService",
// 		Issue:      fooService(),
// 	}
// }
