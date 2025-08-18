package helpers

import (
	"fmt"
	"time"
)

func GenerateSessionID() string {
	return fmt.Sprintf("session_%d", time.Now().UnixNano())
}
