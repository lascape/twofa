package generate

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

//Generate 计算Time-based One-time Password 数字
func Generate(secret string) (Authenticator, error) {
	t := time.Now().Unix()
	
	count := uint64(t) / 30
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		return Authenticator{}, err
	}

	codeInt := hotp(key, count, 6)
	code, err := codeAlignment(codeInt)
	if err != nil {
		return Authenticator{}, err
	}

	return Authenticator{
		Secret: secret,
		Expire: int(30 - (t % 30)),
		Code:   code,
	}, nil
}

func codeAlignment(code int) (int, error) {
	intFormat := fmt.Sprintf("%%0%dd", 6)
	codeStr := fmt.Sprintf(intFormat, code)
	atoi, err := strconv.Atoi(codeStr)
	if err != nil {
		return 0, err
	}
	return atoi, nil
}

func hotp(key []byte, counter uint64, digits int) int {
	h := hmac.New(sha1.New, key)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)
	//取十进制的余数
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}
