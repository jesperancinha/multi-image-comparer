echo ---- Creating system variables...
export GOROOT=/usr/lib/go
export GOPATH=/opt/go/multi-image-comparer/lib
export PATH=$PATH:$GOPATH:$GOROOT/bin
go env

echo ---- Loading commong libraries
go get "github.com/stretchr/testify/assert"
echo -------- Loaded github.com/stretchr/testify/assert
go get "github.com/stretchr/objx"
echo -------- Loaded github.com/stretchr/objx
go get "github.com/steelzack/images/points"
echo -------- Loaded github.com/stretchr/objx
go get "github.com/steelzack/images/points"
echo -------- Loaded github.com/steelzack/images/points

echo ---- Creating pkg builds
# cd lib-multi-comparer
# go install -v -gcflags "-N -l" ./...
# cd ..
