package HelperFunctions

import (
	"RealTimeChatApp/Backend/GlobalVariables"
	MongoConfig "RealTimeChatApp/Backend/Mongo"
	"context"
	"log"
	"net/mail"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func SetSessionCookie(GinContext *gin.Context, SessionID string) {
	GinContext.SetCookie("SessionID", SessionID, GlobalVariables.CookieExpirationSeconds, "/", "localhost", false, false)
}

func ValidateUsername(Username string) bool {
	TrimmedUsername := strings.TrimSpace(Username)
	if len(TrimmedUsername) >= 2 {
		UsernameRegex, _ := regexp.Compile("^[a-zA-Z]{2,}$")

		ValidUsername := UsernameRegex.MatchString(TrimmedUsername)
		if ValidUsername {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func ValidatePassword(Password string) bool {
	// TODO: Move all regexes to a seperate file
	TrimmedPassword := strings.TrimSpace(Password)
	if len(TrimmedPassword) >= 15 {
		AtLeastOneLowerCaseRegex := regexp.MustCompile(`[a-z]`)
		AtLeastOneUpperCaseRegex := regexp.MustCompile(`[A-Z]`)
		AtLeastOneDigitRegex := regexp.MustCompile(`[\\d]`)
		SpecialRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)

		PasswordContainsLowerCase := AtLeastOneLowerCaseRegex.MatchString(TrimmedPassword)
		PasswordContainsUpperCase := AtLeastOneUpperCaseRegex.MatchString(TrimmedPassword)
		PasswordContainsDigit := AtLeastOneDigitRegex.MatchString(TrimmedPassword)
		PasswordHasSpecialCharacter := SpecialRegex.MatchString(TrimmedPassword)

		if PasswordContainsLowerCase && PasswordContainsUpperCase && PasswordContainsDigit && PasswordHasSpecialCharacter {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func ValidateEmail(Email string) bool {
	TrimmedEmai := strings.TrimSpace(Email)
	_, err := mail.ParseAddress(TrimmedEmai)
	if err != nil {
		return false
	} else {
		return true
	}
}

func HashPassword(Password string) string {
	// TODO: The password shoihuld be using a cost value from an env file.
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(HashedPassword)
}

func UsernameExists(Username string) bool {
	// TODO: Rename to UserExists
	var user MongoConfig.User
	TrimmedUsername := strings.TrimSpace(Username)
	filter := bson.M{"username": TrimmedUsername}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func EmailExists(Email string) bool {
	var user MongoConfig.User
	TrimmedEmail := strings.TrimSpace(Email)
	filter := bson.M{"email": TrimmedEmail}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func UserExistsByID(UserID bson.ObjectID) bool {
	var user MongoConfig.User
	filter := bson.M{"_id": UserID}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetUserData(UserID bson.ObjectID) (string, error) {
	var user MongoConfig.User
	filter := bson.M{"_id": UserID}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}
