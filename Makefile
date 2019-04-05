# Makefile variables
PROJECT_NAME=go-szondi

DOCKERFILE_URL="https://gist.githubusercontent.com/vavas/b7b6099d8696e61b848d8a2866104c2b/raw/31bb7cb8544827f458ab931864a21e68ab9614a5/dokerfile"

# go build
compile: dep
	mkdir -p dist
	go build -o dist/${PROJECT_NAME}

# go dep
dep:
	govendor sync
depinfo:
	govendor list && govendor status