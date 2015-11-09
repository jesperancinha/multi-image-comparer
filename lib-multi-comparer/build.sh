export GOROOT=/usr/lib/go
export GOPATH=/opt/go/multi-image-comparer/lib-multi-comparer
export PATH=$PATH:$GOPATH:$GOROOT/bin
export GOBIN=/opt/go/multi-image-comparer/lib-multi-comparer/bin
go env
# go install src/com/steelzack/multicomparer/constants.go src/com/steelzack/multicomparer/image-utils.go
# go install constants.go image-utils.go
# go build -x constants.go image-utils.go
go get "github.com/stretchr/testify/assert"
go get "github.com/stretchr/objx"
go build -x src/com/steelzack/multicomparer/image-utils.go 

