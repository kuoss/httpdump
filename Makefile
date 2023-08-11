VERSION := v0.1.0
IMAGE := ghcr.io/kuoss/httpinfo:$(VERSION)

run:
	go run ./...

docker:
	docker build -t $(IMAGE) --build-arg VERSION=$(VERSION) . && docker push $(IMAGE)
