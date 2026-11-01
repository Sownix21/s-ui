package util

import (
	"encoding/base64"
	"strings"
)

func decodeBase64(str string) ([]byte, error) {
	str = strings.TrimSpace(str)
	var err error
	for _, enc := range []*base64.Encoding{base64.StdEncoding, base64.RawStdEncoding, base64.URLEncoding, base64.RawURLEncoding} {
		var decoded []byte
		if decoded, err = enc.DecodeString(str); err == nil {
			return decoded, nil
		}
	}
	return nil, err
}

// Function to return decoded bytes if a string is Base64 encoded
func StrOrBase64Encoded(str string) string {
	decoded, err := decodeBase64(str)
	if err == nil {
		return string(decoded)
	}
	return str
}

func B64StrToByte(str string) ([]byte, error) {
	return decodeBase64(str)
}

func ByteToB64Str(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
