package domain

type Permission struct {
	role     string `json:"role"`
	service  string `json:"service"`
	resource string `json:"resource"`
	action   string `json:"action"`
}
