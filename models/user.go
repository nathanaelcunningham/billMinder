package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSettings struct {
	UserID             int64  `json:"user_id"`
	ReminderOccurrence string `json:"reminder_occurrence"`
	Timezone           string `json:"timezone"`
}
