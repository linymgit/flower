package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

const md5_salt = "forrily"

func GetPasswordWithMd5(pw string) (md5Pw string) {
	m5 := md5.New()
	m5.Write([]byte(pw))
	m5.Write([]byte(md5_salt))
	st := m5.Sum(nil)
	md5Pw = hex.EncodeToString(st)
	return
}

func ValidePassword(pw string, md5Pw string) (ok bool) {
	s := GetPasswordWithMd5(pw)
	ok = strings.EqualFold(s, md5Pw)
	return
}
