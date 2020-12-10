package mock

import (
	"github.com/cvstom/go-zendesk/zendesk"
)

var _ zendesk.API = (*Client)(nil)
