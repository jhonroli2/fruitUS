package domain

type User struct {
    ID       string
    TenantID string
    Name     string
    Email    string
    Password string
}

func (u *User) Validate() error {
    // Validation logic for User
    return nil
}