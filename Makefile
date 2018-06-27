clean:
	rm -f tunnel

tunnel: clean
	CGO_ENABLED=0 go build ./cmd/tunnel

tunnel-image: tunnel
	go get github.com/openshift/imagebuilder/cmd/imagebuilder
	imagebuilder -f Dockerfile -t docker.io/jimminter/tunnel:latest .

tunnel-push: tunnel-image
	docker push docker.io/jimminter/tunnel:latest

.PHONY: clean tunnel-image tunnel-push
