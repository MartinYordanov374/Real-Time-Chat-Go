package HelperFunctions

import(
	"regexp"
	"strings"
	"net/mail"
	"golang.org/x/crypto/bcrypt"
	"RealTimeChatApp/Backend/GlobalVariables"
	"RealTimeChatApp/Backend/Mongo"
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func SetSessionCookie(GinContext *gin.Context, SessionID string){
	// TODO: Make the cookie last as long as the session, i.e., create a global variable for this
	GinContext.SetCookie("SessionID", SessionID, 60, "/", "localhost", false, false)
}

func ValidateUsername(Username string) bool{
	TrimmedUsername := strings.TrimSpace(Username)
	if len(TrimmedUsername) >= 2{
		UsernameRegex, _ := regexp.Compile("^[a-zA-Z]{2,}$")

		ValidUsername := UsernameRegex.MatchString(TrimmedUsername)
		if ValidUsername {
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}

func ValidatePassword(Password string) bool{
	// TODO: Move all regexes to a seperate file
	TrimmedPassword := strings.TrimSpace(Password)
	if len(TrimmedPassword) >= 15{
		AtLeastOneLowerCaseRegex:= regexp.MustCompile(`[a-z]`)
		AtLeastOneUpperCaseRegex:= regexp.MustCompile(`[A-Z]`)
		AtLeastOneDigitRegex := regexp.MustCompile(`[\\d]`)
		SpecialRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)

		PasswordContainsLowerCase := AtLeastOneLowerCaseRegex.MatchString(TrimmedPassword)
		PasswordContainsUpperCase := AtLeastOneUpperCaseRegex.MatchString(TrimmedPassword)
		PasswordContainsDigit := AtLeastOneDigitRegex.MatchString(TrimmedPassword)
		PasswordHasSpecialCharacter := SpecialRegex.MatchString(TrimmedPassword)


		if (PasswordContainsLowerCase && PasswordContainsUpperCase && PasswordContainsDigit && PasswordHasSpecialCharacter){
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}

func ValidateEmail(Email string) bool{
	TrimmedEmai := strings.TrimSpace(Email)
	_, err := mail.ParseAddress(TrimmedEmai)
	if err != nil{
		return false
	}else{
		return true
	}
}

func HashPassword(Password string) string{
	// TODO: The password shoihuld be using a cost value from an env file.
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil{
		panic(err)
	}

	return string(HashedPassword)
}

func UsernameExists(Username string) bool{
	// TODO: Rename to UserExists
	var user MongoConfig.User
	TrimmedUsername := strings.TrimSpace(Username)
	filter := bson.M{"username": TrimmedUsername}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil{
		log.Println(err)
		return false
	}else{
		return true
	}
}

func EmailExists(Email string) bool{
	var user MongoConfig.User
	TrimmedEmail := strings.TrimSpace(Email)
	filter := bson.M{"email": TrimmedEmail}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil{
		log.Println(err)
		return false
	}else{
		return true
	}

	return false
}
