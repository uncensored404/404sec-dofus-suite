# 404Sec Dofus Suite

Compagnon indépendant en Go : gestion de builds, estimation générique des dégâts, planification XP, dashboard web et API REST locale.

> Pas de bot, d'injection, de lecture mémoire ou d'interaction automatisée avec le client. Les données de démonstration sont fictives et les formules doivent être vérifiées pour la version de Dofus utilisée.

## Prérequis
Go 1.22+

## Lancer
```bash
go test ./...
go run ./cmd/dofus404 --help
go run ./cmd/dofus404 builds list
go run ./cmd/dofus404 damage --base 100 --power 250 --percent 20 --flat 12 --resist 15
go run ./cmd/dofus404 xp --current 120000 --target 250000
go run ./cmd/dofus404 api --addr 127.0.0.1:8787
```
Puis ouvre http://127.0.0.1:8787.

## API
- `GET /health`
- `GET /api/v1/builds`
- `POST /api/v1/damage` — JSON: `{"base":100,"power":250,"percent_damage":20,"flat_bonus":12,"resistance_percent":15}`
- `POST /api/v1/xp` — JSON: `{"current":120000,"target":250000}`

Market CSV d'exemple : `data/market.csv`; les données sont importées manuellement, aucune requête réseau n'est effectuée.
