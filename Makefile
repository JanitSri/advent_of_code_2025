run-go:
	go test ./solutions/util.go ./solutions/$(arg)_test.go -v -timeout 0