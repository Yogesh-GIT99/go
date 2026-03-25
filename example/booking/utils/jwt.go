package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretKey"

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))

}

func VerifyToken(token string) (int64, error) {
	Parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { // to get parsed token

		_, Ok := token.Method.(*jwt.SigningMethodHMAC) // recomended method by go to check the token created is using the method above mentioned.

		if !Ok {
			return nil, errors.New("Unexpected Signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	IsTokenValid := Parsedtoken.Valid // checking the token is valid or not ( figure out)

	// This means:

	//  Signature matched ✔
	//  Token structure correct ✔
	//  Claims are valid ✔

	// What .Valid actually checks

	// 	Internally:
	// 	Signature verification ✔
	// 	Expiry (exp) ✔
	// 	Not before (nbf) ✔ (if present)

	if !IsTokenValid {
		return 0, errors.New("Invalid Token!")
	}

	claims, Ok := Parsedtoken.Claims.(jwt.MapClaims) // The parsed token is a type of map claims or not ( that we used in above func).
	// Returns Claim values and bool to check weather it matches the type or not.

	if !Ok {
		return 0, errors.New("Invalid Token Claims!")
	}

	// email := Claims["email"].(string)
	userId := claims["userId"].(int64) // way to retrive claim value

	return userId, nil // now after all the check it is proved that the token is valid.

}
