package util

import (
	"strings"
	"testing"
)

func TestRandImage(t *testing.T) {
	url := RandImage()
	if !strings.Contains(url, "img.hacpai.com") {
		t.Errorf(url)
	}
}
