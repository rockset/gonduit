package requests

import "github.com/uber/gonduit/entities"

// AuditQueryRequest represents a request to audit.query API method.
type AuditQueryRequest struct {
	AuditorPHIDs []string `json:"auditorPHIDs,omitempty"`
	CommitPHIDs  []string `json:"commitPHIDs,omitempty"`
	Status       string   `json:"status,omitempty"`

	*entities.Cursor
	Request
}
