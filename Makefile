MAIN_PACKAGE_PATH:=./main.go
BINARY_NAME:=app
.PHONY:	build
build:
								go	build	-o=${BINARY_NAME}	${MAIN_PACKAGE_PATH}