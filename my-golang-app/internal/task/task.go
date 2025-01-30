package task

type Task struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

func (t *Task) UpdateTask(newTask string) {
	t.Text = newTask
}
