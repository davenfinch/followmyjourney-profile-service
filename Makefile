BINARY=profile-service
IMAGE=gcr.io/PROJECT-ID/followmyjourney-profile-service:latest

.PHONY: build run docker-build docker-push deploy-cloudrun

build:
	go build -v -o ./build/$(BINARY) ./cmd/server

run:
	go run ./cmd/server

docker-build:
	docker build -t $(IMAGE) .

docker-push: docker-build
	docker push $(IMAGE)

deploy-cloudrun:
	gcloud run deploy followmyjourney-profile-service --image $(IMAGE) --region=us-central1 --platform=managed --allow-unauthenticated
