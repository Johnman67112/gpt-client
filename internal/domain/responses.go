package domain

type Function struct {
	Name      string
	Arguments string
}

type ToolCall struct {
	Id       int64
	Type     string
	Function Function
}

type TopLogProb struct {
	Token   string
	LogProb int64
	Bytes   []byte
}

type Content struct {
	Token       string
	LogProb     int64
	Bytes       []byte
	TopLogProbs []TopLogProb
}

type LogProb struct {
	Content Content
}

type ResponseMessage struct {
	Content   string
	ToolCalls []ToolCall
	Role      string
}

type Choice struct {
	FinishReason string
	Index        int64
	Message      ResponseMessage
	LogProbs     []LogProb
}

type Usage struct {
	CompletionTokens int64
	PromptTokens     int64
	TotalTokens      int64
}

type ChatResponse struct {
	Id                string
	Choices           []Choice
	Created           int64
	Model             string
	SystemFirgerprint string
	Object            string
	Usage             Usage
}
