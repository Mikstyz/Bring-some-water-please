package uniquecash

import (
	"fmt"

	"github.com/google/uuid"
)

func NewCash(userID int64) string {
	return fmt.Sprintf("%d-%s", userID, uuid.NewString())
}
