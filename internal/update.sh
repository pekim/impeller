#!/bin/bash

FLUTTER_SHA=4d3c9c58afe01bf68fae7b002918b014e2da90c7
TEMP_DIR=temp

mkdir -p $TEMP_DIR

PLATFORMS=(
  #GOOS_GOARCH    ARCHIVE_PATH  LIBRARY_FILENAME
  "darwin_arm64   darwin-arm64  libimpeller.dylib"
  "darwin_amd64   darwin-x64    libimpeller.dylib"
  "linux_arm64    linux-arm64   libimpeller.so"
  "linux_amd64    linux-x64     libimpeller.so"
)
for PLATFORM in "${PLATFORMS[@]}"
do
  read -a PARTS <<< "$PLATFORM"
  GOOS_GOARCH="${PARTS[0]}"
  ARCHIVE_PATH="${PARTS[1]}"
  LIBRARY_FILENAME="${PARTS[2]}"

  ZIP_FILE=$TEMP_DIR/impeller_sdk_$GOOS_GOARCH.zip
  DEST_DIR=internal/interop-$GOOS_GOARCH

  echo "platform : $GOOS_GOARCH"

  # Download an archive with prebuilt artifacts.
  rm -fr $DEST_DIR
  mkdir -p $DEST_DIR
  echo "  download"
  wget \
    --no-verbose \
    --output-document $ZIP_FILE \
    https://storage.googleapis.com/flutter_infra_release/flutter/$FLUTTER_SHA/$ARCHIVE_PATH/impeller_sdk.zip

  # Extract the archive, which includes the C header file and the pre-built library.
  echo "  extract"
  unzip -q $ZIP_FILE -d $DEST_DIR
  # Remove large, unwanted files
  rm -r $DEST_DIR/bin
  rm $DEST_DIR/lib/libimpeller.a

  # Create an GOOS/GOARCH-specific file that embeds the relevant shared library.
cat <<EOF > impeller_$GOOS_GOARCH.go
package impeller

import _ "embed"

//go:embed $DEST_DIR/lib/$LIBRARY_FILENAME
var sharedObject []byte
EOF

  echo
done


# Create a Go file with a constant for the Flutter SHA
cat <<EOF > fluttersha.go
package impeller

const FlutterSHA = "$FLUTTER_SHA"
EOF

rm -r temp
