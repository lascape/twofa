package generate

import (
	"encoding/base32"
	"strings"

	"github.com/google/uuid"
)

func Secret() string {
	uid := strings.Replace(uuid.New().String(), "-", "", -1)[0:10]
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString([]byte(uid))
}
