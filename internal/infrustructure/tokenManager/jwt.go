package tokenManager

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"logity/config"
	"logity/internal/domain/usecase/auth/dto"
	"time"
)

type TokenManager struct {
	ttlAccess, ttlRefresh, ttlRealTimeServer                   time.Duration
	secretAccessKey, secretRefreshKey, secretRealTimeServerKey string
	issuer                                                     string
}

func NewTokenManager(cfg *config.Config) *TokenManager {
	return &TokenManager{
		ttlAccess:               time.Duration(cfg.AccessTokenTTLInSec) * time.Second,
		ttlRefresh:              time.Duration(cfg.RefreshTokenTTLInSec) * time.Second,
		ttlRealTimeServer:       time.Duration(cfg.Centrifugo.TokenTTLInSec) * time.Second,
		secretAccessKey:         cfg.SecretAccessKey,
		secretRefreshKey:        cfg.SecretRefreshKey,
		secretRealTimeServerKey: cfg.Centrifugo.SecretKey,
		issuer:                  cfg.App.Host,
	}
}

func (m *TokenManager) NewAccessToken(userId string) (jwtToken dto.JWT, err error) {
	expiredAt := time.Now().Add(m.ttlAccess)
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": userId,
		"exp": expiredAt.Unix(),
		"iss": m.issuer,
		"aud": m.issuer,
	})
	signedToken, err := unsignedToken.SignedString([]byte(m.secretAccessKey))
	if err != nil {
		return jwtToken, err
	}
	return dto.JWT{
		Token:     signedToken,
		ExpiredAt: expiredAt,
	}, nil
}

func (m *TokenManager) NewRealTimeToken(userId string) (jwtToken dto.JWT, err error) {
	expiredAt := time.Now().Add(m.ttlRealTimeServer)
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": userId,
		"exp": expiredAt.Unix(),
		"iss": m.issuer,
		"aud": "centrifugo",
	})
	signedToken, err := unsignedToken.SignedString([]byte(m.secretRealTimeServerKey))
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
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": userId,
		"exp": expiredAt.Unix(),
		"iss": m.issuer,
		"aud": m.issuer,
	})
	signedToken, err := unsignedToken.SignedString([]byte(m.secretRefreshKey))
	if err != nil {
		return jwtToken, err
	}
	return dto.JWT{
		Token:     signedToken,
		ExpiredAt: expiredAt,
	}, nil
}

func (m *TokenManager) ParseToken(token string) (*dto.PayloadToken, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method isn't HMAC")
		}

		return []byte(m.secretAccessKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parse jwt token: %w", err)
	}

	claims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
	if !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("error parse claims jwt token")
	}

	return &dto.PayloadToken{
		UserId: claims.Subject,
		Token: dto.JWT{
			Token:     jwtToken.Raw,
			ExpiredAt: claims.ExpiresAt.Time,
		},
	}, nil
}
