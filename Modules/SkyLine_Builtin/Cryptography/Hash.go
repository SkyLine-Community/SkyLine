package SkyLine_Crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

var Hasher = map[string]func(string) string{
	"MD5": func(s string) string {
		return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	},
	"SHA1": func(s string) string {
		return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
	},
	"SHA224": func(s string) string {
		return fmt.Sprintf("%x", sha256.Sum224([]byte(s)))
	},
	"SHA256": func(s string) string {
		return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
	},
	"SHA384": func(s string) string {
		return fmt.Sprintf("%x", sha512.Sum384([]byte(s)))
	},
	"SHA512": func(s string) string {
		return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
	},
}
