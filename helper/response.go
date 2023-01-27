package helper

type StudentResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	NIM   string `json:"nim"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}
