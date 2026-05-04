package schema

import (
	"encoding/json"
	"time"
)


// M2MResponse matches your strict JSON schema
type ResponseV1 struct {
	Metadata Metadata `json:"metadata"`
	Content  Content  `json:"content"`
}

type Metadata struct {
	Version   string    `json:"version"`
	RequestID string    `json:"request_id"`
	Timestamp time.Time `json:"timestamp"` // Serializes to RFC3339 (ISO 8601)
}

type Content struct {
	Introduction string   `json:"introduction"`
	Elaboration  string   `json:"elaboration"`
	Summary      string   `json:"summary"`
	Keywords     []string `json:"keywords"`
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
