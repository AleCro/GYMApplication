package Db

type PaginatedUsersResult struct {
	Users []*User `json:"users"`
	Total int64   `json:"total"`
	Limit int64   `json:"limit"`
	Page  int64   `json:"page"`
}

func (r *PaginatedUsersResult) RemovePasswords() {
	for _, user := range r.Users {
		user.Password = ""
	}
}
