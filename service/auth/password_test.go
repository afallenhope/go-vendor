package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	t.Run("should hash the password", func(t *testing.T) {
		hash, err := HashPassword("testpassword")
		if err != nil {
			t.Errorf("error hashing password: %v", err)
		}

		if hash == "" {
			t.Error("expected hash to not be empty")
		}

		if hash == "testpassword" {
			t.Error("expected hash to be different from plain text")
		}
	})
}

func TestComparePassword(t *testing.T) {
	t.Run("should compare passwords", func(t *testing.T) {
		hash, err := HashPassword("testpassword")
		if err != nil {
			t.Errorf("error hashing password: %v", err)
		}

		if !ComparePasswords(hash, []byte("testpassword")) {
			t.Errorf("expected password to match hash")
		}

		if ComparePasswords(hash, []byte("invalidpassword")) {
			t.Errorf("expected password to be invalid")
		}
	})
}
