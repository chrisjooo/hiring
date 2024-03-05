package usecase

import (
	"fmt"
	"time"

	"github.com/christianchrisjo/hiring/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// should be in config -- for simplicity, I put it here
var secretKey = []byte("hiring-secret-key")

func (u *Usecase) CreateUser(req models.CreateUserRequest) (models.User, error) {
	err := req.Validate()
	if err != nil {
		return models.User{}, err
	}

	user, err := u.postgres.CreateUser(models.User{
		UserID:      uuid.New(),
		Email:       req.Email,
		Password:    req.Password,
		Type:        req.Type,
		Description: req.Description,
		Name:        req.Name,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *Usecase) GetUserByEmail(email string) (models.User, error) {
	if email == "" {
		return models.User{}, models.ErrEmailRequired
	}
	user, err := u.postgres.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *Usecase) GetUserByID(id string) (models.User, error) {
	if id == "" {
		return models.User{}, models.ErrEmailRequired
	}
	user, err := u.postgres.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *Usecase) UpdateUser(req models.UpdateUserRequest) (models.User, error) {
	existing, err := u.postgres.GetUserByID(req.UserID.String())
	if err != nil {
		return models.User{}, err
	}

	err = req.Validate(existing)
	if err != nil {
		return models.User{}, err
	}

	// update data
	existing.Description = req.Description
	existing.Name = req.Name
	existing.UpdatedAt = time.Now()

	user, err := u.postgres.UpdateUser(existing)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// sign in by email
func (u *Usecase) SignInByEmail(email, password string) (string, error) {
	password = models.HashPassword(password)

	match, err := u.postgres.CheckUserCredsByEmail(email, password)
	if err != nil {
		return "", err
	}
	if !match {
		return "", models.ErrInvalidCreds
	}

	user, err := u.postgres.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	token, err := createToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func createToken(user models.User) (string, error) {
	claimToken := &models.ClaimToken{
		UserID:    user.UserID,
		Email:     user.Email,
		Type:      string(user.Type),
		Name:      user.Name,
		ExpiredAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken).SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func DecodeToken(token string) (claim models.ClaimToken) {
	jwt.ParseWithClaims(token, &claim, func(t *jwt.Token) (interface{}, error) {
		return nil, nil
	}, jwt.WithoutClaimsValidation())
	return claim
}
