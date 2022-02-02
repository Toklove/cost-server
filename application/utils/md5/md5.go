package MD5

import (
	"crypto/md5"
	"fmt"
)

func Create(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	result := h.Sum(nil)
	return fmt.Sprintf("%x", result)
}
