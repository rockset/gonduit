package responses

import (
	"github.com/uber/gonduit/entities"
	"github.com/uber/gonduit/util"
)

// DiffusionQueryCommitsResponse represents a response of the
// diffusion.querycommits call.
type DiffusionQueryCommitsResponse struct {
	Data          map[string]entities.DiffusionCommit `json:"data"`
	IdentifierMap map[string]string                   `json:"identifierMap"`
	Cursor        entities.Cursor                     `json:"cursor"`
}

// DiffusionRepositorySearchResponse contains fields that are in server
// response to differential.revision.search.
type DiffusionRepositorySearchResponse struct {
	// Data contains search results.
	Data []*DiffusionRepositorySearchResponseItem `json:"data"`

	// Curson contains paging data.
	Cursor SearchCursor `json:"cursor,omitempty"`
}

// DiffusionRepositorySearchResponseItem contains information about a
// particular search result.
type DiffusionRepositorySearchResponseItem struct {
	ResponseObject
	Fields      DiffusionRepositorySearchResponseItemFields `json:"fields"`
	Attachments DiffusionRepositorySearchAttachments        `json:"attachments"`
	SearchCursor
}

// DiffusionRepositorySearchResponseItemFields is a collection of object
// fields.
type DiffusionRepositorySearchResponseItemFields struct {
	Name          string                         `json:"name"`
	VCS           string                         `json:"vcs"`
	Callsign      string                         `json:"callsign"`
	ShortName     string                         `json:"shortName"`
	Status        string                         `json:"status"`
	IsImporting   bool                           `json:"isImporting"`
	DefaultBranch string                         `json:"defaultBranch"`
	Description   DiffusionRepositoryDescription `json:"description"`
	DateCreated   util.UnixTimestamp             `json:"dateCreated"`
	DateModified  util.UnixTimestamp             `json:"dateModified"`
	SpacePHID     string                         `json:"spacePHID"`
}

// DiffusionRepositoryDescription holds the description of repository.
type DiffusionRepositoryDescription struct {
	Raw string `json:"raw"`
}

// DiffusionRepositorySearchAttachments holds possible attachments for the API
// method.
type DiffusionRepositorySearchAttachments struct {
	URIs     SearchAttachmentURIs     `json:"uris,omitempty"`
	Metrics  SearchAttachmentMetrics  `json:"metrics,omitempty"`
	Projects SearchAttachmentProjects `json:"projects,omitempty"`
}

type RepositoryURIItem struct {
	Fields RepositoryURIItemFields `json:"fields"`
}

// RepositoryURIItemFields is a collection of object fields.
type RepositoryURIItemFields struct {
	URI          RepositoryURI      `json:"uri"`
	Disabled     bool               `json:"disabled"`
	DateCreated  util.UnixTimestamp `json:"dateCreated"`
	DateModified util.UnixTimestamp `json:"dateModified"`
}

// RepositoryURI is VCS uri.
type RepositoryURI struct {
	Raw        string `json:"raw"`
	Display    string `json:"display"`
	Effective  string `json:"effective"`
	Normalized string `json:"normalized"`
}

// DiffusionCommitSearchResponse contains information about a particular search result.
type DiffusionCommitSearchResponse struct {
	// Data contains search results.
	Data []*DiffusionCommitSearchResponseItem `json:"data"`

	// Curson contains paging data.
	Cursor SearchCursor `json:"cursor,omitempty"`
}

// DiffusionCommitSearchResponseItem contains information about a particular search result.
type DiffusionCommitSearchResponseItem struct {
	ID     uint64                                  `json:"id"`
	Type   string                                  `json:"type"`
	PHID   string                                  `json:"phid"`
	Fields DiffusionCommitSearchResponseItemFields `json:"fields"`
}

// DiffusionCommitSearchResponseItemFields is a collection of object fields.
type DiffusionCommitSearchResponseItemFields struct {
	RepositoryPHID string                     `json:"repositoryPHID"`
	Message        string                     `json:"message"`
	IsUnreachable  bool                       `json:"isUnreachable"`
	IsImported     bool                       `json:"isImported"`
	Author         DiffusionCommitUser        `json:"author"`
	Committer      DiffusionCommitUser        `json:"committer"`
	AuditStatus    DiffusionCommitAuditStatus `json:"auditStatus"`
}

// DiffusionCommitUser is used both for the commit author and committer
type DiffusionCommitUser struct {
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	Raw          string             `json:"raw"`
	Epoch        util.UnixTimestamp `json:"epoch"`
	IdentityPHID string             `json:"identityPHID"`
	UserPHID     string             `json:"userPHID"`
}

// DiffusionCommitAuditStatus contains the commit audit status.
type DiffusionCommitAuditStatus struct {
	Value  string `json:"value"`
	Name   string `json:"name"`
	Closed bool   `json:"closed"`
}
