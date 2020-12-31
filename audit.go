package gonduit

import (
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
)

const AuditQueryMethod = "audit.query"

// AuditQuery is the method name on API.
func (c *Conn) AuditQuery(req requests.AuditQueryRequest) ([]responses.AuditQueryResponse, error) {
	var resp []responses.AuditQueryResponse
	if err := c.Call(AuditQueryMethod, &req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
