package responses

// AuditQueryResponse contains fields that are in server response to audit.query.
type AuditQueryResponse struct {
	ID          string   `json:"id"`
	CommitPHID  string   `json:"commitPHID"`
	AuditorPHID string   `json:"auditorPHID"`
	Status      string   `json:"status"`
	Reasons     []string `json:"reasons"`
}
