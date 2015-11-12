echo ---- Creating system variables...
export GOROOT=/usr/lib/go
export PATH=$PATH:$GOPATH:$GOROOT/bin
go env

export GOPATH=/opt/go/multi-image-comparer/pearson-correlation-coefficient
echo ---- Loading commong libraries
go get "github.com/stretchr/testify/assert"
echo -------- Loaded github.com/stretchr/testify/assert
go get "github.com/stretchr/objx"
echo -------- Loaded github.com/stretchr/objx
go get "github.com/steelzack/images/points"
echo -------- Loaded github.com/stretchr/objx
go get "github.com/steelzack/images/points"
echo -------- Loaded github.com/steelzack/images/points

export GOPATH=/opt/go/multi-image-comparer/tanimoto-correlation-coefficient
echo ---- Loading commong libraries
go get "github.com/stretchr/testify/assert"
echo -------- Loaded github.com/stretchr/testify/assert
go get "github.com/stretchr/objx"
echo -------- Loaded github.com/stretchr/objx
go get "github.com/steelzack/images/points"
echo -------- Loaded github.com/stretchr/objx
go get "github.com/steelzack/images/points"
echo -------- Loaded github.com/steelzack/images/points


