package domain

type Message struct {
	Content string
	Role    string
	Name    string
}

type ChatRequest struct {
	Messages []Message
	Model    string
}
