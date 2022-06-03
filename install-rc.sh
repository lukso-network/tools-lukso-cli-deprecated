#!/usr/bin/env bash


# for Apple M1s
if [ "$(uname -s)" == "Darwin" ] && [ "$(uname -m)" == "arm64" ]
then
ARCHITECTURE="amd64"
else
ARCHITECTURE=$(uname -m)
ARCHITECTURE=${ARCHITECTURE/x86_64/amd64}
ARCHITECTURE=${ARCHITECTURE/aarch64/arm64}
fi
readonly os_arch_suffix="$(uname -s | tr '[:upper:]' '[:lower:]')-$ARCHITECTURE"

PLATFORM=""
case "$OSTYPE" in
darwin*) PLATFORM="darwin" ;;
linux*) PLATFORM="linux" ;;
msys*) PLATFORM="windows" ;;
cygwin*) PLATFORM="windows" ;;
*) exit 1 ;;
esac
readonly PLATFORM

if [ "$PLATFORM" == "windows" ]; then
    ARCHITECTURE="amd64.exe"
elif [[ "$os_arch_suffix" == *"arm64"* ]]; then
    ARCHITECTURE="arm64"
fi

if [[ "$ARCHITECTURE" == "armv7l" ]]; then
    color "31" "32-bit ARM is not supported. Please install a 64-bit operating system."
    exit 1
fi

download() {
  URL="$1";
  echo $URL
  LOCATION="$2";
  if [[ $PLATFORM == "linux" ]]; then
    wget -O $LOCATION $URL;
  fi

  if [[ $PLATFORM == "darwin" ]]; then
    curl -o $LOCATION -Lk $URL;
  fi
}


# create binary location if not exists
mkdir -p /usr/local/bin/
# download lukso and give exec permission
download https://github.com/lukso-network/lukso-cli/releases/download/v0.2.0/lukso-cli-${PLATFORM}-${ARCHITECTURE} /usr/local/bin/lukso
chmod +x /usr/local/bin/lukso


echo ""
echo "#################### LUKSO CLI ####################"
echo "use \"lukso --help\" to check available options"
echo "###############################################################"

