package enumer

import (
	"strings"

	"gitlab.com/eAuction/enumer/pkg/normalization"
)

var ownershipReplacer = strings.NewReplacer("<", "less", ">", "more")

func normalizeOwnershipDocumentE(s string) (string, error) {
	return normalization.NormalizeE(ownershipReplacer.Replace(s))
}
