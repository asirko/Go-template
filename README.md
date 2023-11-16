# Go Template

## Description

Ce projet est un template très basique d'application backend en GO. 
Il est basé sur l'architecture hexagonale.

L'objectif est de compléter ce template pour avoir une base de code réutilisable pour les projets futurs.

L'architecture et les choix pris sont disponibles [ici](./docs/static/architecture.md).

## Reste à faire

- passer en UUID pour les users
- mettre à jour le swagger et validé qu'il se génère seul (voir go-redoc?)
- mettre à jour le readme 
- mettre un logger en place
- revoir l'ORM pour utiliser un système de seeding
- avoir de meilleurs messages d'erreur pour les 400
- gérer les codes d'erreurs
- gérer une authentification avec double token (l'access dans la réponse le refresh dans les cookies)
- Utiliser gvm pour gérer les versions de go et simplifier les prérequis 
- vérifier la génération automatique de la doc swagger et la documentation de la base de données
- Mettre en place de la CI (voir golangci)
- regrouper password utils et paseto dans un pkg utils
- déplacer les utils dans un package utils
- Pas besoin d'avoir 2 niveaux pour handler/http, garder uniquement http

## Démarrer

### Prérequis

Avoir [Go](https://go.dev/dl/) 1.21 ou + et [Task](https://taskfile.dev/installation/) installé sur la machine :

    ```bash
    go version && task --version
    ```

### Installation

```bash
task install
cp .env.example .env
```

### Lancer le projet

```bash
task service:up
task migrate:up
task dev
```
