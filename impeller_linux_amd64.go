package impeller

import _ "embed"

//go:embed internal/interop-linux_amd64/lib/libimpeller.so
var sharedObject []byte
