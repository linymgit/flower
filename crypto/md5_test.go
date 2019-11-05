package crypto

import (
	"fmt"
	"testing"
)

func TestGetPasswordWithMd5(t *testing.T) {
	md5Pw := GetPasswordWithMd5("12345")
	fmt.Printf("%v", md5Pw)

}

func TestValidePassword(t *testing.T) {
	ok := ValidePassword("12345", "b3d1996418936fbd8c0b1241a0a7b225")
	fmt.Printf("%v", ok)
}
