//go:build js
// +build js

package qtls

// Check the cpu flags for each platform that has optimized GCM implementations.
// Worst case, these variables will just all be false.
var (
	hasGCMAsmAMD64 = false
	hasGCMAsmARM64 = false
	hasGCMAsmS390X = false

	hasGCMAsm = false
)
