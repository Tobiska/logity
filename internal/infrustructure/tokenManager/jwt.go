package tokenManager

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"logity/config"
	"logity/internal/domain/usecase/auth/dto"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
}

type TokenManager struct {
	ttlAccess, ttlRefresh             time.Duration
	secretAccessKey, secretRefreshKey string
	issuer                            string
}

func NewTokenManager(cfg *config.Config) *TokenManager {
	return &TokenManager{
		ttlAccess:        time.Duration(cfg.AccessTokenTTLInSec) * time.Second,
		ttlRefresh:       time.Duration(cfg.RefreshTokenTTLInSec) * time.Second,
		secretAccessKey:  cfg.SecretAccessKey,
		secretRefreshKey: cfg.SecretRefreshKey,
		issuer:           cfg.Host,
	}
}

func (m *TokenManager) NewJWT(userId string) (jwtToken dto.JWT, err error) {
	expiredAt := time.Now().Add(m.ttlAccess)
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"sub": userId,
		"exp": expiredAt.Unix(),
		"iss": m.issuer,
	})
	signedToken, err := unsignedToken.SignedString(m.secretAccessKey)
	if err != nil {
		return jwtToken, err
	}
	return dto.JWT{
		Token:     signedToken,
		ExpiredAt: expiredAt,
	}, nil
}

func (m *TokenManager) NewRefreshToken(userId string) (jwtToken dto.JWT, err error) {
	expiredAt := time.Now().Add(m.ttlRefresh)
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"sub": userId,
		"exp": expiredAt.Unix(),
		"iss": m.issuer,
	})
	signedToken, err := unsignedToken.SignedString(m.secretRefreshKey)
	if err != nil {
		return jwtToken, err
	}
	return dto.JWT{
		Token:     signedToken,
		ExpiredAt: expiredAt,
	}, nil
}

func (m *TokenManager) ParseToken(token string) (*dto.PayloadToken, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method isn't HMAC")
		}

		return m.secretAccessKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parse jwt token: %w", err)
	}

	claims, ok := jwtToken.Claims.(Claims)
	if !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("error parse claims jwt token")
	}

	return &dto.PayloadToken{
		UserId: claims.Subject,
	}, nil
}
