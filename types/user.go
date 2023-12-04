package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	
)

const(
	bcryptCost = 12
	MinFirstNameLength = 2
	MinLastNameLength = 2
	MinPasswordLength = 7
)

type User struct{
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`
	Email string `json:"email" bson:"email"`
	EncryptedPassword string `json:"-" bson:"EncryptedPassword"`

}

type CreateUserParams struct{
	FirstName string `json:"firstName" `
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`

}

func (params CreateUserParams) ValidiateUsers() []string {
	var errors [] string
	if len(params.FirstName) < MinFirstNameLength{
		errors = append(errors,fmt.Sprintf("first Name length should be atleast %d characters", MinFirstNameLength))
	}

	if len(params.LastName) < MinLastNameLength{
		errors = append(errors,fmt.Sprintf("last Name length should be atleast %d characters",MinLastNameLength))
		
	}

	if len(params.Password) < MinPasswordLength{
		errors = append(errors, fmt.Sprintf("passwordlength should be atleast %d characters",MinPasswordLength))
	}
	if !isValidEmail(params.Email){
		errors = append(errors, fmt.Sprintf("email is invalid") )
	}
	return errors
	
}


func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
    encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
    if err != nil {
        return nil, err
    }
    return &User{
        FirstName:         params.FirstName,
        LastName:          params.LastName,
        Email:             params.Email,
        EncryptedPassword: string(encpw),
    }, nil
}
