#!/bin/bash
Namespace="lishimeng"

# shellcheck disable=SC2046
Version=$(git describe --tags $(git rev-list --tags --max-count=1))
# shellcheck disable=SC2154
GitCommit=$(git log --pretty=format:"%h" -1)
BuildTime=$(date +%FT%T%z)

checkout_tag(){
  git checkout "${Version}"
}

common(){
  echo ""
}

build_image(){
  local Name=$1
  local AppPath=$2
  print_app_info "${Name}" "${AppPath}"

  docker build -t "${Namespace}/${Name}:${Version}" \
  --build-arg NAME="${Name}" \
  --build-arg VERSION="${Version}" \
  --build-arg BUILD_TIME="${BuildTime}" \
  --build-arg COMMIT="${GitCommit}" \
  --build-arg APP_PATH="${AppPath}" -f "./${AppPath}/Dockerfile" .
}

print_app_info(){
  local Name=$1
  local AppPath=$2
  echo "****************************************"
  echo "App:${Name}[${Namespace}]"
  echo "Version:${Version}"
  echo "Commit:${GitCommit}"
  echo "Build:${BuildTime}"
  echo "Main_Path:${AppPath}"
  echo "****************************************"
  echo ""
}

push_image(){
  local Name=$1
  echo "****************************************"
  echo "Push:${Namespace}:${Name}:${Version}"
  echo "****************************************"
  echo ""
  docker tag  "${Namespace}/${Name}:${Version}" "${Namespace}/${Name}"
  docker push "${Namespace}/${Name}:${Version}"
  docker push "${Namespace}/${Name}"
}

build_all(){
  common
  checkout_tag
  build_image 'passport' 'cmd/passport'
  build_image 'passport-profile' 'cmd/profile'
}

push_all(){
  common
  push_image 'passport'
  push_image 'passport-profile'
}

case  $1 in
    push)
		push_all
        ;;
    *)
		build_all
        ;;
esac

