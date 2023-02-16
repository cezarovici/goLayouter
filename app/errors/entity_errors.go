package errors

import (
	"fmt"
	"time"
)

type ErrEntity struct {
	Caller     string
	MethodName string
	Issue      error

	WithTime bool
}

const areaErrEntity = "Entity"

func (e *ErrEntity) Error() string {
	var res [4]string

	if e.WithTime {
		res[0] = fmt.Sprintf("\nArea: %s [%d]", areaErrEntity, time.Now().Unix())
	} else {
		res[0] = fmt.Sprintf("\nArea: %s", areaErrEntity)
	}

	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Method Name: %s", e.MethodName)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}
