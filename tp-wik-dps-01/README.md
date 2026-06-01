# Ping Server

Petit serveur HTTP développé en Go avec uniquement la bibliothèque standard.

## Fonctionnalités

* Écoute sur un port configurable via la variable d'environnement `PING_LISTEN_PORT`
* Retourne les en-têtes (headers) de la requête au format JSON
* Accepte uniquement les requêtes `GET` sur l'endpoint `/ping`
* Retourne une erreur `404 Not Found` pour toute autre route ou méthode HTTP
* Aucune dépendance externe

## Structure du projet

```text
tp-wik-dps-01/
├── main.go
├── handlers/
│   └── ping.go
├── config/
│   └── config.go
├── go.mod
└── README.md
```

## Prérequis

* Go 1.20 ou supérieur

## Installation

Cloner le projet puis initialiser les dépendances :

```bash
go mod tidy
```

## Configuration

Le serveur utilise la variable d'environnement suivante :

PING_LISTEN_PORT | Port d'écoute du serveur | 8080      

### Exemple

Linux / macOS :

```bash
export PING_LISTEN_PORT=9000
```

Windows PowerShell :

```powershell
$env:PING_LISTEN_PORT="9000"
```

## Lancement

```bash
go run .
```

Ou :

```bash
go build -o tp-wik-dps-01
./tp-wik-dps-01
```

## Exemple de test

```bash
curl http://localhost:8080/ping
```

```bash
curl -X POST http://localhost:8080/ping
```

Le premier appel retourne les headers de la requête au format JSON, tandis que le second retourne une erreur 404.

## Technologies utilisées

* Go
* net/http
* encoding/json
* os