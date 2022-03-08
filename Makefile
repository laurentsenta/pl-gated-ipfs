all: randomize
	go build . && ./pl-gated-ipfs -folder ./my-test-folder

randomize:
	for i in `ls ./my-test-folder`; do openssl rand -hex 10 >> "./my-test-folder/$$i"; done;