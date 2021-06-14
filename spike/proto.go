package spike

import "context"

type Spike struct {
	ctx     context.Context
	methods map[string]string
}
