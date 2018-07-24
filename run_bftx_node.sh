echo "Checking if golang is installed"
command -v go >/dev/null 2>&1 || { echo >&2 "Go-bftx requires Go but it appears not to be installed. Please check your system has Golang installed (https://golang.org/) before installing Go-bftx. Aborting."; exit 1; }

echo "Downloding go-bftx code"
go get github.com/blockfreight/go-bftx

cd $GOPATH//src/github.com/blockfreight/go-bftx

echo "Resolving dependencies"
dep ensure

echo "Initializing Blockfreight Node"
bftx init

echo "Starting Blockfreight Node"
bftnode