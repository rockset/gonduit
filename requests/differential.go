package requests

import "github.com/uber/gonduit/constants"

// DifferentialGetCommitMessageRequest represents a request to the
// differential.getcommitmessage call.
type DifferentialGetCommitMessageRequest struct {
	RevisionID uint64                                         `json:"revision_id"`
	Fields     []string                                       `json:"fields"`
	Edit       constants.DifferentialGetCommitMessageEditType `json:"edit"`
	Request
}

// DifferentialQueryRequest represents a request to the
// differential.query call.
type DifferentialQueryRequest struct {
	Authors          []string                           `json:"authors"`
	CCs              []string                           `json:"ccs"`
	Reviewers        []string                           `json:"reviewers"`
	Paths            [][]string                         `json:"paths"`
	CommitHashes     [][]string                         `json:"commitHashes"`
	Status           constants.DifferentialStatusLegacy `json:"status"`
	Order            constants.DifferentialQueryOrder   `json:"order"`
	Limit            uint64                             `json:"limit"`
	Offset           uint64                             `json:"offset"`
	IDs              []uint64                           `json:"ids"`
	PHIDs            []string                           `json:"phids"`
	Subscribers      []string                           `json:"subscribers"`
	ResponsibleUsers []string                           `json:"responsibleUsers"`
	Branches         []string                           `json:"branches"`
	Request
}

// DifferentialQueryDiffsRequest represents a request
// to the differential.querydiffs call.
type DifferentialQueryDiffsRequest struct {
	IDs         []uint64 `json:"ids"`
	RevisionIDs []uint64 `json:"revisionIDs"`
	Request
}

// DifferentialGetCommitPathsRequest represents a request to the
// differential.getcommitpaths call.
type DifferentialGetCommitPathsRequest struct {
	RevisionID uint64 `json:"revision_id"`
	Request
}