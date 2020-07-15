# Usage:
# make init-test

init-test:
	@docker-compose up --build

init-dev:
	@docker-compose up --build -d