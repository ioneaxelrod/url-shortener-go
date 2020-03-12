package shorten

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
)

func CreateKey() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		err := fmt.Errorf("error: %v", err)
		return "", err
	}
	smallStr := base64.RawURLEncoding.EncodeToString(uuid[:])
	return smallStr, err
}

//func URLToKey(url string) (uint32, error) {
//	// Create a numeric hash
//	hash := hash(url)
//}
//
//func hash(s string) uint32 {
//	h := fnv.New32a()
//	h.Write([]byte(s))
//	return h.Sum32()
//}