package remote

import (
	"strings"

	"github.com/gateway-fm/cdk-erigon-lib/gointerfaces/types"
)

func NodeInfoReplyLess(i, j *types.NodeInfoReply) int {
	if cmp := strings.Compare(i.Name, j.Name); cmp != 0 {
		return cmp
	}
	return strings.Compare(i.Enode, j.Enode)
}
