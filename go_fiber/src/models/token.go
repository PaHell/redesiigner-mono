package models

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/PaHell/redesiigner-mono/database"
	"github.com/PaHell/redesiigner-mono/util"
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/twinj/uuid"
)

type Token struct {
	// keys
	UserID       uint   `gorm:"ForeignKey" json:"user_id"`
	AccessToken  string `gorm:"index" json:"access_token"`
	RefreshToken string `gorm:"index" json:"refresh_token"`
	// expiration
	AccessExpires  time.Time `json:"access_expires"`
	RefreshExpires time.Time `json:"refresh_expires"`
	// relational
	User User `json:"user" gorm:"ForeignKey:UserID"`
}

type TokenModel struct{}

func (m TokenModel) OneByAccess(access string) (token *Token, err error) {
	err = database.InstanceGorm.Preload("User").First(&token, "access_token = ?", access).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_CANNOT_GET)
	}
	return token, nil
}

func (m TokenModel) OneByRefresh(refresh string) (token *Token, err error) {
	err = database.InstanceGorm.Preload("User").First(&token, "refresh_token = ?", refresh).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_CANNOT_GET)
	}
	return token, nil
}

func (m TokenModel) All() (list []Token, err error) {
	err = database.InstanceGorm.Preload("User").Find(&list).Error
	if err != nil {
		return list, errors.New(util.RESPONSE_MSG_CANNOT_GET)
	}
	return list, nil
}

func (m TokenModel) AllForUser(userID uint) (list []Token, err error) {
	err = database.InstanceGorm.Preload("User").Where("user_id = ?", userID).Find(&list).Error
	log.Print(len(list))
	if err != nil {
		return list, errors.New(util.RESPONSE_MSG_CANNOT_GET)
	}
	return list, nil
}

func (m TokenModel) Create(userID uint) (*Token, error) {
	var err error
	td := &Token{
		UserID:         userID,
		AccessExpires:  time.Now().Add(time.Minute * 15),
		RefreshExpires: time.Now().Add(time.Hour * 24 * 7),
	}
	// create jwt token
	atClaims := &jwt.MapClaims{
		"authorized":  true,
		"access_uuid": uuid.NewV4().String(),
		"user_id":     userID,
		"exp":         td.AccessExpires.Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	// create refresh token
	rtClaims := &jwt.MapClaims{
		"authorized":   true,
		"refresh_uuid": uuid.NewV4().String(),
		"user_id":      userID,
		"exp":          td.RefreshExpires.Unix(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	err = database.InstanceGorm.Create(&td).Error
	if err != nil {
		return nil, err
	}
	return m.OneByAccess(td.AccessToken)
}

func (m TokenModel) DeleteAllForUser(userID uint) (tokens []Token, err error) {
	err = database.InstanceGorm.Where("user_id = ?", userID).Delete(&tokens).Error
	return tokens, err
}

func (m TokenModel) DeleteByAccess(access string) (deleted *Token, err error) {
	err = database.InstanceGorm.First(&deleted, "AccessToken = ?", access).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_DOES_NOT_EXIST)
	}
	database.InstanceGorm.Delete(&deleted)
	return deleted, nil
}

/*

//VerifyToken ...
func (m TokenModel) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := m.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//TokenValid ...
func (m TokenModel) TokenValid(r *http.Request) error {
	token, err := m.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

//ExtractTokenMetadata ...
func (m TokenModel) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := m.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     uint(userID),
		}, nil
	}
	return nil, err
}

*/
