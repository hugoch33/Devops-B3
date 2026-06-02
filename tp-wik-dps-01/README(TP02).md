# WIK-DPS-TP02 - Dockerisation de l'API (WIK-DPS-TP01)

Ce TP contient deux images Docker pour l'API Go du TP01:

- une image **single-stage** (`Dockerfile`)
- une image **multi-stage** (`Dockerfile.multi`)

Les deux versions respectent les contraintes:

- ordre des layers optimisé pour accélérer les rebuilds
- exécution avec un utilisateur dédié non-root
- possibilité de scanner les vulnérabilités de l'image

## Prérequis

- Docker Desktop (ou Docker Engine)
- Projet API TP01 présent dans `../tp-wik-dps-01`
- Outil de scan:
  - Docker Scout (`docker scout`)

## Structure

```text
tp-wik-dps-01/
├── Dockerfile
├── Dockerfile.multi
└── .dockerignore

tp-wik-dps-02/
└── README.md
```

## Image 1: Single-stage

Depuis la racine du workspace:

```powershell
cd C:\Users\chama\Devops-B3\tp-wik-dps-01
docker build -t wik-dps-api:single -f Dockerfile .
```

### Exécution

```powershell
docker run --rm -p 8080:8080 -e PING_LISTEN_PORT=8080 wik-dps-api:single
```

Test:

```powershell
curl http://localhost:8080/ping
```

## Image 2: Multi-stage

```powershell
cd C:\Users\chama\Devops-B3\tp-wik-dps-01
docker build -t wik-dps-api:multi -f Dockerfile.multi .
```

### Exécution

```powershell
docker run --rm -p 8080:8080 -e PING_LISTEN_PORT=8080 wik-dps-api:multi
```

## Utilisateur d'exécution non-root

Dans les deux Dockerfiles:

- création de `appuser` / `appgroup`
- directive `USER appuser`

## Scan de vulnérabilités

```powershell
docker scout cves wik-dps-api:single
docker scout cves wik-dps-api:multi
```