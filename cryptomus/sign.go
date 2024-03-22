package cryptomus

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func (c *Cryptomus) sign(rawData []byte) string {
	data := base64.StdEncoding.EncodeToString(rawData)
	hash := md5.Sum([]byte(data + c.apiKey))

	return hex.EncodeToString(hash[:])
}
