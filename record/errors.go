package record

import "errors"

var (
	//errRecordExists This user already has a recordr
	errRecordExists = errors.New("This user already has a Record")

	//errProblemLoadingRecord This user already has a record
	errProblemLoadingRecord = errors.New("There was a problem loading one or more records in this list")
)
