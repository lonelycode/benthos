//go:build !arm

package couchbase

import (
	// Bring in the internal plugin definitions.
	_ "github.com/TykTechnologies/benthos/v4/internal/impl/couchbase"
)
