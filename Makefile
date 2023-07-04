BINARY_NAME=gato
NEW_LINE=echo

build: 
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

run: build 
	./${BINARY_NAME}-linux

test: run
	${NEW_LINE}
	echo "Test 1: reading two files"
	./${BINARY_NAME}-linux test1.txt

	${NEW_LINE}
	echo "Test 2: reading file from stdin"
	head -n1 test1.txt | ./${BINARY_NAME}-linux -

	${NEW_LINE}
	echo "Test 3: reading file and display lines"
	./${BINARY_NAME}-linux -n test1.txt

	${NEW_LINE}
	echo "Test 3: reading file from stdin and display lines skipping empty lines"
	sed G test1.txt | ./${BINARY_NAME}-linux -b | head -n10	

clean: 
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows