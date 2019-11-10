package controller

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	upToken := getUploadToken()
	print(upToken)
}

