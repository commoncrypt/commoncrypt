package pushchallenge

type Model struct{}

type Status string

const (
	StatusPending   Status = "pending"
	StatusCompleted Status = "completed"
	StatusNotFound  Status = "notfound"
)
