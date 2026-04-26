package schema

import (
	"encoding/json"
	"time"
)

type ResponseV1 struct {
	SchemaVersion string    `json:"schema_version"`
	RequestId     string    `json:"request_id"`
	Timestamp     time.Time `json:"timestamp"`
	Request       string    `json:"request"`
	Introduction  string    `json:"introduction"`
	Description   string    `json:"description"`
	Summary       string    `json:"summary"`
	Keywords      []string  `json:"keywords"`
}

func Convert(responseV1Json string) ResponseV1 {
	j := responseV1Json
	var responseV1 ResponseV1
	if err := json.Unmarshal([]byte(j), &responseV1); err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", responseV1) // {Name:Alice Age:30}
	return responseV1
}
