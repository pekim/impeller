#!/bin/bash

FLUTTER_SHA=4d3c9c58afe01bf68fae7b002918b014e2da90c7
PLATFORM_ARCH=linux-x64
DEST_DIR=internal/interop
TEMP_DIR=temp
ZIP_FILE=$TEMP_DIR/impeller_sdk.zip

mkdir -p $DEST_DIR
mkdir -p $TEMP_DIR

# Download an archive with prebuilt artifacts.
wget \
  --no-verbose \
  --output-document $ZIP_FILE \
  https://storage.googleapis.com/flutter_infra_release/flutter/$FLUTTER_SHA/$PLATFORM_ARCH/impeller_sdk.zip

# Extract the C header file, pre-build library.
unzip $ZIP_FILE -d $DEST_DIR
