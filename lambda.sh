#!/bin/sh

set -eu

os="linux"
arch="amd64"

_build_pkg() {
	base_dir="$1"
	pkg="$2"

	(
		cd "${base_dir}/lambda/$pkg"
		GOOS="${os}" GOARCH="${arch}" go build -o "${base_dir}/build/${os}/${arch}/${pkg}"
	)
}

_build() {
	base_dir="$1"

	rm -rf "${base_dir}/build"
	_build_pkg "${base_dir}" "convert"
}

_deploy() {
	base_dir="$1"

	_build "${base_dir}"

	deploy_zip_path="${base_dir}/deploy.zip"
	(
		cd "${base_dir}/build/${os}/${arch}"
		zip "${deploy_zip_path}" ./*
	)
	aws lambda update-function-code --function-name JapaneseCalendarSkill --zip-file "fileb://${deploy_zip_path}"
	rm "${deploy_zip_path}"
}

_main() {
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
