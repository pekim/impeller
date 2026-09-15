package impeller

import _ "embed"

//go:embed internal/interop-darwin_amd64/lib/libimpeller.dylib
var sharedObject []byte
