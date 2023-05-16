package link

import "github.com/Wishrem/wuso/pkg/pool"

func init() {
	MsgPool = pool.NewPool(4)
}
