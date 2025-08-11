#!/bin/bash
set -e

# ---- config ----
APP_NAME="pkneuron"
VERSION="1.2.0"                  # set your app version
MAINTAINER="Pezhman Kasraee <github@pezhmankasraee.com>"
DESCRIPTION="PKNeuron - An interactive LLM assistant for Linux terminal"

ARCH="$(dpkg --print-architecture)"  # usually amd64
WORKDIR="$(pwd)/debbuild"
DEB="${APP_NAME}_${VERSION}_${ARCH}.deb"

# ---- build directory structure ----
rm -rf "$WORKDIR"
mkdir -p "$WORKDIR/DEBIAN"
mkdir -p "$WORKDIR/usr/local/bin"

# ---- copy binary ----
cp "$APP_NAME" "$WORKDIR/usr/local/bin/$APP_NAME"

# ---- control file ----
cat > "$WORKDIR/DEBIAN/control" <<EOF
Package: $APP_NAME
Version: $VERSION
Section: utils
Priority: optional
Architecture: $ARCH
Depends:
Maintainer: $MAINTAINER
Description: $DESCRIPTION
EOF

# ---- permissions ----
chmod 755 "$WORKDIR/usr/local/bin/$APP_NAME"
chmod 755 "$WORKDIR/DEBIAN"

# ---- build the deb ----
dpkg-deb --build "$WORKDIR" "$DEB"

echo "Created $DEB"

rm -rf "$WORKDIR"
echo "Cleaned up build directory."