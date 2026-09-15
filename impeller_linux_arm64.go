package impeller

import _ "embed"

//go:embed internal/interop-linux_arm64/lib/libimpeller.so
var sharedObject []byte
