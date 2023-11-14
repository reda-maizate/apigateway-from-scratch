package domain

type Note struct {
	Title   string
	Content string
}

func (n *Note) GetTitle() string {
	return n.Title
}

func (n *Note) GetContent() string {
	return n.Content
}

func (n *Note) New(title, content string) *Note {
	return &Note{
		Title:   title,
		Content: content,
	}
}
