export GOROOT=/usr/lib/go
export GOPATH=/opt/go/multi-image-comparer/lib-multi-comparer:/opt/go/multi-image-comparer/mic-example-images:/opt/go/multi-image-comparer/pearson-correlation-coefficient
export PATH=$PATH:$GOPATH:$GOROOT/bin
export GOBIN=/opt/go/multi-image-comparer/lib-multi-comparer/bin:/opt/go/multi-image-comparer/mic-example-images/bin:/opt/go/multi-image-comparer/pearson-correlation-coefficient/bin
go env
