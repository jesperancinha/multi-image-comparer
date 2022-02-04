coverage:
	cd pearson-correlation-coefficient && go test -coverprofile=../coverage-pc.out
	cd tanimoto-correlation-coefficient && go test -coverprofile=../coverage-tc.out
