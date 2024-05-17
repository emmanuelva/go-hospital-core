package utils_test

import (
	"github.com/emmanuelva/go-hospital-core/internal/utils"
	"testing"
)

func TestNormalizeText(t *testing.T) {
	expected := "test string"

	got, err := utils.NormalizeText("Tést StrInG")

	if err != nil {
		t.Errorf("NormalizeText() error = %v", err)
	}

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}
