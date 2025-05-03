IMAGE_NAME = livreria

test:
	go test ./... -v

build-image:
	docker build -t $(IMAGE_NAME):$(version) .

run-docker:
	docker run -d --name $(IMAGE_NAME) -p 8080:8080 $(IMAGE_NAME):$(version)

run:
	docker compose down && docker compose up --build