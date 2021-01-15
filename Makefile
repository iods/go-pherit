
up:
	@docker-compose up -d


build:
	@docker-compose build --no-cache --force-rm


destroy:
	@docker-compose down --rmi all --volumes --remove-orphans


roman-build:
	@docker-compose -f build/compose/docker-compose.roman.yml build --no-cache --force-rm
	@docker-compose -f build/compose/docker-compose.roman.yml up -d

roman-destroy:
	@docker-compose -f build/compose/docker-compose.roman.yml down --rmi all --volumes --remove-orphans