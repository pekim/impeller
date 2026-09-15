package impeller

import _ "embed"

//go:embed internal/interop-darwin_arm64/lib/libimpeller.dylib
var sharedObject []byte
