package chores

type Chore struct {
	ChoreName        string
	ChoreDescription string
	ChoreId          string
	SmoreValue       int64
}

type ChoreHandler struct {
	Database map[string]Chore
}