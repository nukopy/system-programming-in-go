dcb:
	docker-compose build

dcu:
	docker-compose up -d

dcin:
	docker-compose exec system-programming-in-go /bin/bash

dcd:
	docker-compose down

# help
.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
