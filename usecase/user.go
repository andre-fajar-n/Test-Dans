package usecase

import (
	"context"
	"dans/entity"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		userPostgre entity.UserPostgre
		Config      Config
	}

	Config struct {
		SecretKey string
	}
)

func NewUser(
	userPostgre entity.UserPostgre,
	cfg *viper.Viper,
) entity.UserUsecase {
	configUser := cfg.Sub("app")
	config := Config{
		SecretKey: configUser.GetString("secret_key"),
	}
	return &User{
		userPostgre,
		config,
	}
}

func (u *User) Register(ctx context.Context, req *entity.UserRegisterRequest) error {
	user, err := u.userPostgre.GetByUsername(ctx, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if user != nil {
		return errors.New("username already registered")
	}

	// hash password
	hashed, err := u.HashAndSalt(ctx, req.Password)
	if err != nil {
		return err
	}

	if err := u.userPostgre.Create(ctx, entity.User{
		Username: req.Username,
		Password: hashed,
	}); err != nil {
		return err
	}

	return nil
}

func (u *User) Login(ctx context.Context, req *entity.UserLoginRequest) (*entity.LoginResponse, error) {
	user, err := u.userPostgre.GetByUsername(ctx, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if user == nil || err == gorm.ErrRecordNotFound {
		return nil, errors.New("username not found")
	}

	// compare password
	if !u.ComparePasswords(ctx, user.Password, req.Password) {
		return nil, errors.New("invalid password")
	}

	expiredTime := time.Now().UTC()
	token, err := u.CreateToken(ctx, req.Username, expiredTime)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		Token:     token,
		ExpiredAt: expiredTime.Format(time.RFC3339),
	}, nil
}

func (u *User) CreateToken(ctx context.Context, username string, expiredTime time.Time) (string, error) {
	// Create the token
	token := jwt.New(jwt.GetSigningMethod(jwt.SigningMethodHS256.Name))

	now := time.Now().Local()
	token.Claims = jwt.MapClaims{
		"username": username,
		"iat":      now.Unix(),
		"exp":      expiredTime.Unix(),
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(u.Config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *User) HashAndSalt(ctx context.Context, password string) (string, error) {
	// hash and salt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *User) ComparePasswords(ctx context.Context, hashed, plain string) bool {
	// compare the hashed and plain passwords
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
