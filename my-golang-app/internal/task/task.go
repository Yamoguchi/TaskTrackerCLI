package task

type Task struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}
