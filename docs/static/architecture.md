# Architecture de l'application

## Librairies

- GIN : Framework web responsable du routage d'API et de validation des données, il est configurable et ressemble à ExpressJS
- SquirreL : ORM très basique

## Outils

- Task : Task runner pour gérer les tâches de développement
- Swaggo : Générateur de documentation d'API
- Docker et docker-compose : conteneur d'exécution en mode dev et constitution d'image pour la production
- Air : Reloader pour le développement
- Golang-migrate : Migration de la base de données

## Structure

- `cmd` : contient les fichiers d'entrée de l'application
- `internal` : contient le code de l'application

TODO
J'aimerai simplifier la structure, mais tant que la DI n'est pas refaite, je ne peux pas le faire.