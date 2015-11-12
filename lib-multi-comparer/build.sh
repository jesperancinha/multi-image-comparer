export GOROOT=/usr/lib/go
export GOPATH=/opt/go/multi-image-comparer/lib
export PATH=$PATH:$GOPATH:$GOROOT/bin
export GOBIN=/opt/go/multi-image-comparer/lib-multi-comparer/bin
go env 
go get "github.com/stretchr/testify/assert"
go get "github.com/stretchr/objx"
go get "github.com/steelzack/lib-multi-comparer/multicomparer"

export GOPATH=/opt/go/multi-image-comparer/lib-multi-comparer
go build -x src/github.com/steelzack/lib-multi-comparer/multicomparer/image-utils.go 

