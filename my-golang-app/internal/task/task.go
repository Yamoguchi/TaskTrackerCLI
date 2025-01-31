package task

type Task struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

func (t *Task) UpdateTaskText(newTask string) {
	t.Text = newTask
}

func (t *Task) UpdateTaskStatus(newStatus string) {
	t.Status = newStatus
}
