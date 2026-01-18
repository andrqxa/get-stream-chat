package proposal

type Status string

const (
	StatusSent     Status = "sent"
	StatusAccepted Status = "accepted"
	StatusRejected Status = "rejected"
)

type Proposal struct {
	ID        string
	ProjectID string
	FromUser  string
	ToUser    string
	Status    Status
}
