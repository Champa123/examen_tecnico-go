package controller

type SummaryRequest struct {
	PathToFile string `json:"path_to_file"`
	UserMail   string `json:"user_mail"`
}
