package auth

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateJWT(t *testing.T) {
	t.Run("should create JWT", func(t *testing.T) {
		secret := []byte("secret")

		uid, err := uuid.NewUUID()
		if err != nil {
			t.Errorf("error creating uuid: %v", err)
		}

		token, err := CreateJWT(secret, uid)
		if err != nil {
			t.Errorf("error creating JWT: %v", err)
		}

		if token == "" {
			t.Error("expected non-empty JWT token")
		}

	})
}
