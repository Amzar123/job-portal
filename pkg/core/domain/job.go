package domain

type Job struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Company     string `json:"company"`
	JobType     string `json:"job_type"`
	Experience  string `json:"experience"`
	Salary      string `json:"salary"`
	PostDate    string `json:"post_date"`
	ExpireDate  string `json:"expire_date"`
	IsRemote    bool   `json:"is_remote"`
	IsApproved  bool   `json:"is_approved"`
	IsFilled    bool   `json:"is_filled"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
