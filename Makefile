
#
# Docker & Docker Compose
#=
up:
	@docker-compose up -d


build:
	@docker-compose build --no-cache --force-rm


destroy:
	@docker-compose down --rmi all --volumes --remove-orphans


#
# gRPC Application Build
#
gen:
	@bash ./scripts/generate.sh

clean-pb:
	@rm -rf ./internal/rpc/pb/**/*.pb.go