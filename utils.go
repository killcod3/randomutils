package randomutils

import (
	crand "crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	digits           = "0123456789"
	lowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"
	upperCaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomStr(pattern string) (string, error) {
	var result strings.Builder
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case '?':
			i++
			if i < len(pattern) {
				switch pattern[i] {
				case 'd':
					result.WriteByte(digits[rand.Intn(len(digits))])
				case 'l':
					result.WriteByte(lowerCaseLetters[rand.Intn(len(lowerCaseLetters))])
				case 'u':
					result.WriteByte(upperCaseLetters[rand.Intn(len(upperCaseLetters))])
				default:
					return "", errors.New("Invalid pattern identifier")
				}
			} else {
				result.WriteByte('?')
			}
		default:
			result.WriteByte(pattern[i])
		}
	}
	return result.String(), nil
}

func RandNumInRange(min, max int) (int, error) {
	if min > max {
		return 0, errors.New("Invalid range: min must be less than or equal to max")
	}

	return rand.Intn(max-min+1) + min, nil
}

func GetRandHex(length int) string {
	if length <= 0 {
		return ""
	}

	const hexDigits = "0123456789abcdef"
	var result strings.Builder

	for i := 0; i < length; i++ {
		result.WriteByte(hexDigits[rand.Intn(len(hexDigits))])
	}

	return result.String()
}

func GetRandBytes(length int) []byte {
	if length <= 0 {
		return []byte{}
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = byte(rand.Intn(256))
	}

	return result
}

func GetRandInt(length int) (int, error) {
	if length <= 0 {
		return 0, errors.New("Length must be greater than 0")
	}

	min := int(math.Pow10(length - 1))
	max := int(math.Pow10(length)) - 1
	return rand.Intn(max-min+1) + min, nil
}

func GetUUIDv1() (string, error) {
	timestamp := uint64(time.Now().UnixNano() / 100)
	node := make([]byte, 6)
	if _, err := rand.Read(node); err != nil {
		return "", err
	}
	node[0] |= 0x01
	var sequence uint16
	if err := binary.Read(crand.Reader, binary.BigEndian, &sequence); err != nil {
		return "", err
	}
	uuidBytes := make([]byte, 16)
	binary.BigEndian.PutUint32(uuidBytes[0:], uint32(timestamp))
	binary.BigEndian.PutUint16(uuidBytes[4:], uint16(timestamp>>32))
	binary.BigEndian.PutUint16(uuidBytes[6:], uint16(timestamp>>48|0x1000))
	binary.BigEndian.PutUint16(uuidBytes[8:], sequence|0x8000)
	copy(uuidBytes[10:], node)
	uuidStr := hex.EncodeToString(uuidBytes)
	uuidStr = fmt.Sprintf("%s-%s-%s-%s-%s", uuidStr[:8], uuidStr[8:12], uuidStr[12:16], uuidStr[16:20], uuidStr[20:])
	return uuidStr, nil
}

func GetUUIDv4() (string, error) {
	uuidBytes := make([]byte, 16)
	if _, err := rand.Read(uuidBytes); err != nil {
		return "", err
	}

	uuidBytes[6] = (uuidBytes[6] & 0x0F) | 0x40 // Set version to 4
	uuidBytes[8] = (uuidBytes[8] & 0x3F) | 0x80 // Set variant to 1

	uuidStr := hex.EncodeToString(uuidBytes)
	uuidStr = fmt.Sprintf("%s-%s-%s-%s-%s", uuidStr[:8], uuidStr[8:12], uuidStr[12:16], uuidStr[16:20], uuidStr[20:])
	return uuidStr, nil
}

func GetUUIDv5(namespace, name string) (string, error) {
	namespaceUUID, err := parseUUID(namespace)
	if err != nil {
		return "", err
	}

	hash := sha1.New()
	hash.Write(namespaceUUID)
	hash.Write([]byte(name))
	hashedBytes := hash.Sum(nil)

	uuidBytes := make([]byte, 16)
	copy(uuidBytes, hashedBytes[:16])

	uuidBytes[6] = (uuidBytes[6] & 0x0F) | 0x50 // Set version to 5
	uuidBytes[8] = (uuidBytes[8] & 0x3F) | 0x80 // Set variant to 1

	uuidStr := hex.EncodeToString(uuidBytes)
	uuidStr = fmt.Sprintf("%s-%s-%s-%s-%s", uuidStr[:8], uuidStr[8:12], uuidStr[12:16], uuidStr[16:20], uuidStr[20:])
	return uuidStr, nil
}

func parseUUID(uuidStr string) ([]byte, error) {
	if len(uuidStr) != 36 {
		return nil, errors.New("Invalid UUID string length")
	}

	uuidStr = uuidStr[0:8] + uuidStr[9:13] + uuidStr[14:18] + uuidStr[19:23] + uuidStr[24:]
	uuidBytes, err := hex.DecodeString(uuidStr)
	if err != nil {
		return nil, err
	}

	return uuidBytes, nil
}
