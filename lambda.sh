#!/bin/sh

set -eu

os="linux"
arch="amd64"

_build_pkg() {
	base_dir="$1"
	pkg="$2"

	(
		cd "${base_dir}/lambda/$pkg"
		GOOS="${os}" GOARCH="${arch}" go build
	)
}

_build() {
	base_dir="$1"

	_build_pkg "${base_dir}" "convert"
}

_deploy_pkg() {
	base_dir="$1"
	pkg="$2"

	_build_pkg "${base_dir}" "${pkg}"

	(
		cd "${base_dir}/lambda/$pkg"
		build-lambda-zip "${pkg}"
		aws lambda update-function-code --function-name JapaneseCalendarSkill --zip-file "fileb://./${pkg}.zip"
	)
}

_deploy() {
	base_dir="$1"

	_deploy_pkg "${base_dir}" "convert"
}

_main() {
	go get github.com/aws/aws-lambda-go/cmd/build-lambda-zip

	base_dir=$(cd "$(dirname $0)" || exit 1 ; pwd)

	sub_cmd="$1"
	shift
	case $sub_cmd in
		"build" ) _build  "${base_dir}" "$@" ;;
		"deploy") _deploy "${base_dir}" "$@" ;;
	  *)"default"
	esac
}

_main "$@"
