package model

type History struct {
	Items []string `json:"items"`
}

func (history *History) Exist(name string) bool {
	var exists bool
	for _, item := range history.Items {
		if item == name {
			exists = true
		}
	}

	return exists
}
