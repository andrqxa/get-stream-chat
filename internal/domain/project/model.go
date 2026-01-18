package project

type ProjectType string

const (
	TypeRemote ProjectType = "remote"
	TypeOnsite ProjectType = "onsite"
	TypeHybrid ProjectType = "hybrid"
)

type Project struct {
	ID      string
	OwnerID string
	Title   string
	Type    ProjectType
}
