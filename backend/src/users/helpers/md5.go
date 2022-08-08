package helpers

import (
	"crypto/md5"
	"fmt"
	"github.com/AlekSi/pointer"
)

func Hash(data *string) *string {
	result := []byte(*data)
	return pointer.ToString(fmt.Sprintf("%x", md5.Sum(result)))
}
