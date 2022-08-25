#!/bin/bash
NAME="passport"
MAIN_PATH="cmd/passport/main.go"
ORG="lishimeng"

# shellcheck disable=SC2046
VERSION=$(git describe --tags $(git rev-list --tags --max-count=1))
# shellcheck disable=SC2154
COMMIT=$(git log --pretty=format:"%h" -1)

build_application(){
  git checkout "${VERSION}"
  docker build -t "${ORG}/${NAME}:${VERSION}" \
  --build-arg NAME="${NAME}" \
  --build-arg VERSION="${VERSION}" \
  --build-arg COMMIT="${COMMIT}" \
  --build-arg MAIN_PATH="${MAIN_PATH}" .
}

print_app_info(){
  echo "****************************************"
  echo "App:${NAME}"
  echo "Version:${VERSION}"
  echo "Commit:${COMMIT}"
  echo "Main_Path:${MAIN_PATH}"
  echo "****************************************"
  echo ""
}

print_app_info
build_application