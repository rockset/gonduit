package gonduit

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uber/gonduit/core"
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/test/server"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAuditQuery(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response, err := ioutil.ReadFile("testdata/audit_query.json")
	require.NoError(t, err)
	s.RegisterMethod(AuditQueryMethod, http.StatusOK, server.ResponseFromJSON(string(response)))

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.AuditQueryRequest{}
	resp, err := c.AuditQuery(req)
	require.NoError(t, err)
	require.Len(t, resp, 2)
	a := resp[0]
	assert.Equal(t, a.ID, "923")
	assert.Equal(t, a.CommitPHID, "PHID-CMIT-glb776fnqc4moh5crpp5")
	assert.Equal(t, a.AuditorPHID, "PHID-PROJ-t7ztapweky34p7dotggk")
	assert.Equal(t, a.Status, "requested")
	assert.Len(t, a.Reasons, 0)
}
