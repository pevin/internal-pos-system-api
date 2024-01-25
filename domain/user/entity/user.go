package entity

type User struct {
	Username         string `dynamodbav:"username" json:"cognito:username"`
	GivenName        string `dynamodbav:"given_name" json:"given_name"`
	FamilyName       string `dynamodbav:"family_name" json:"family_name"`
	Email            string `dynamodbav:"email" json:"email"`
	PhoneNumber      string `dynamodbav:"phone_number" json:"phone_number"`
	Type             string `dynamodbav:"-" json:"custom:type"`
	EmployeeNumber   string `dynamodbav:"-" json:"custom:employee_number"`
	CompanyID        string `dynamodbav:"-" json:"custom:company_id"`
	ConcessionaireID string `dynamodbav:"-" json:"custom:concessionaire_id"`
}

func (u User) ToPartialUser() PartialUser {
	return PartialUser{
		Username: u.Username,
		Name:     u.GivenName + " " + u.FamilyName,
	}
}

type PartialUser struct {
	Username string `dynamodbav:"username" json:"username,omitempty"`
	Name     string `dynamodbav:"name" json:"name,omitempty"`
}
