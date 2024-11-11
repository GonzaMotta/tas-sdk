package types

import (
	"time"
)

type Audit struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time `bson:"created_at"`
}

type HelperAudit struct {
	ID string `bson:"_id,omitempty"`
}

type Trace struct {
	ID          string      `bson:"_id,omitempty" json:"id"`
	Audit       HelperAudit `bson:"audit" json:"audit"`
	ServiceName string      `bson:"service_name" json:"service_name"`
	Data        string      `bson:"data" json:"data"`
	TypeValue   string      `bson:"TypeValue" json:"TypeValue"`
}

// Responses
type CreateAuditResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type CreateTraceResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
