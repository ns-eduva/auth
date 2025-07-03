-include .env

# Redefinir MAKEFILE_LIST pour qu'il ne contienne que le Makefile
MAKEFILE_LIST := Makefile

FILE_COMPOSE := compose.$(APP_ENV).yaml
ENV_FILE := --env-file .env
GO_COMMAND_CONTAINER_TEST := docker exec -i -e APP_ENV=test eduvaauth_dev go test

.PHONY: help cm dev devb devbnod ddev tidy addget
.DEFAULT_GOAL = help

## —— HELP ——————————————————————————————————
help: ## Afficher l'aide
	@grep -E '(^[a-zA-Z0-9\./_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— CONTAINER ——————————————————————————————————
up: ## demarre le service auth
	docker compose $(ENV_FILE) -f docker/$(FILE_COMPOSE) up -d

upb: ## demarre le service auth
	docker compose $(ENV_FILE) -f docker/$(FILE_COMPOSE) up -d --build

down: ## arrete le service auth
	docker compose $(ENV_FILE) -f docker/$(FILE_COMPOSE) down

l: # log d'un container -> make l c=container_name
	docker logs -f $(c)

lapp: ## log de l'application
	$(MAKE) l c="eduvaauth_$(APP_ENV)"

ldb: ## log de la base de données
	$(MAKE) l c="eduvaauth_$(APP_ENV)_db"

s: # ouvre un shell dans un des containers
	docker exec -it $(c) bash

sapp: ## ouvre un shell dans la container de l'application
	$(MAKE) s c="eduvaauth_$(APP_ENV)"

sdb: ## ouvre un shell dans le container de la base de données
	$(MAKE) s c="eduvaauth_$(APP_ENV)_db"

## —— GO ——————————————————————————————————
g: ## execute commande go dans le conteneur auth - usage: make g c=commande_go
	docker exec -it eduvaauth_$(APP_ENV) go $(c)

gg: ## ajoute une dependance - usage: make gg dep=path_de_la_dependance
	$(MAKE) g c="get $(dep)"

tidy: ## execute go mod tidy pour nettoyer les dependances
	$(MAKE) g c="mod tidy"

## —— TEST ——————————————————————————————————
t: ## execute tous les tests
	$(GO_COMMAND_CONTAINER_TEST) ./...

tf: ## execute les tests avec un package - usage: make tf file=chemin/vers/le/package
	$(GO_COMMAND_CONTAINER_TEST) $(file)

tv: ## execute tous les tests avec verbose
	$(GO_COMMAND_CONTAINER_TEST) -v ./...

tvf: ## execute les tests avec verbose d'un package - usage: make tvf file=chemin/vers/le/package
	$(GO_COMMAND_CONTAINER_TEST) -v $(file)

## —— MIGRATION DB ——————————————————————————————————
cm: ## créé un fichier de migration - usage: make cm file=nom_du_fichier
	docker exec -it eduvaauth_$(APP_ENV) migrationcreate $(file)

## —— HELPER DB ——————————————————————————————————
see-db: ## affiche les collections de la base de données
	docker exec -it eduvaauth_$(APP_ENV)_db mongosh --eval "show dbs"

see-coll: ## affiche les collections d'une base de données - usage: make see-coll db=nom_de_la_base
	docker exec -it eduvaauth_$(APP_ENV)_db mongosh $(db) --eval "show collections"