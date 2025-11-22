package Routes

import Db "Svelgok-API/Database"

type LoginRequestForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Different structures in case more data needs to be collected
// on sign up.
type RegisterRequestForm struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordChangeForm struct {
	Original string `json:"password"`
	New      string `json:"new-password"`
}

type UpdateGoalRequest struct {
	Title       string       `json:"title" binding:"required"`
	Description string       `json:"description"`
	SubGoals    []Db.SubGoal `json:"subGoals"`
}
