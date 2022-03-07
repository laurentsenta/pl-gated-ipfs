all:
	go build . && ./pl-gated-ipfs -folder ./my-test-folder
