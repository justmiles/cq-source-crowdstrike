package crowdstrike

import "github.com/go-openapi/strfmt"

type Incident struct {
	// assigned to
	AssignedTo string `json:"assigned_to,omitempty"`

	// assigned to name
	AssignedToName string `json:"assigned_to_name,omitempty"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// created
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created"`

	// description
	Description string `json:"description,omitempty"`

	// end
	// Required: true
	// Format: date-time
	End *strfmt.DateTime `json:"end"`

	// fine score
	// Required: true
	FineScore *int32 `json:"fine_score"`

	// host ids
	// Required: true
	HostIds []string `json:"host_ids"`

	// incident id
	// Required: true
	IncidentID *string `json:"incident_id"`

	// incident type
	IncidentType int64 `json:"incident_type,omitempty"`

	// lm host ids
	LmHostIds []string `json:"lm_host_ids"`

	// lm hosts capped
	LmHostsCapped bool `json:"lm_hosts_capped,omitempty"`

	// modified timestamp
	// Format: date-time
	ModifiedTimestamp strfmt.DateTime `json:"modified_timestamp,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// objectives
	Objectives []string `json:"objectives"`

	// start
	// Required: true
	// Format: date-time
	Start *strfmt.DateTime `json:"start"`

	// state
	// Required: true
	State *string `json:"state"`

	// status
	Status int32 `json:"status,omitempty"`

	// tactics
	Tactics []string `json:"tactics"`

	// tags
	Tags []string `json:"tags"`

	// techniques
	Techniques []string `json:"techniques"`

	// users
	Users []string `json:"users"`

	// visibility
	Visibility int32 `json:"visibility,omitempty"`
}
