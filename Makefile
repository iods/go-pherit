
up:
	@docker-compose up -d


build:
	@docker-compose build --no-cache --force-rm


destroy:
	@docker-compose down --rmi all --volumes --remove-orphans

buffer:
	@bash ./scripts/generate.sh