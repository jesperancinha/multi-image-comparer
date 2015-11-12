echo ---- Creating system variables...
export GOROOT=/usr/lib/go
export GOPATH=/opt/go/multi-image-comparer/lib
export PATH=$PATH:$GOPATH:$GOROOT/bin
export GOBIN=/opt/go/multi-image-comparer/lib-multi-comparer/bin
go env

echo ---- Loading commong libraries
go get "github.com/stretchr/testify/assert"
echo -------- Loaded github.com/stretchr/testify/assert
go get "github.com/stretchr/objx"
echo -------- Loaded github.com/stretchr/objx

echo ---- Creating pkg builds
export GOPATH=/opt/go/multi-image-comparer/lib-multi-comparer
cd lib-multi-comparer
go build -x src/github.com/steelzack/lib-multi-comparer/multicomparer/image-utils.go
go install -v -gcflags "-N -l" ./...
cd ..
