package internal

import "time"

// Shared SDK Constants
const (
	BasePath         = "api/v1"
	TimeoutSeconds   = time.Second * 30
	TokenMaxAge      = int64((time.Hour * 23) / time.Millisecond)
	MQSendTimeout    = time.Millisecond * 5000
	MQReceiveTimeout = time.Millisecond * 5000
	SDKVersion       = "0.0.1"
)

// Environments are Rize infrastructure tiers
var Environments = []string{"sandbox", "integration", "production"}
