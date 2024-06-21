package request

type NewBotRequest struct {
	Name     string `json:"name"`
	StartNow bool   `json:"start_now"`
}
