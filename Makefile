build:
	docker buildx build --push --platform linux/arm64 -t phuctq32/pro-magnet . \
	&& docker buildx build --push --platform linux/amd64 -t phuctq32/pro-magnet:amd64 .

dev:
	@docker compose --env-file ./prod.env up -d \
	mongo-primary mongo-secondary1 mongo-secondary2 redis \
	&& docker compose --env-file ./prod.env run --rm mongo-init

prod:
	docker compose --env-file ./prod.env up -d