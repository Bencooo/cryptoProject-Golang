# Kraken Data Collector – Go + SQLite

Ce projet en **Go** récupère des données publiques depuis l'API **Kraken** et les stocke dans une base de données **SQLite**.

---

## Fonctionnalités

- Récupère l'heure système depuis `/Time`
- Récupère les paires de trading depuis `/AssetPairs`
- Récupère les prix de 10 paires avec `/Ticker`
- Stocke les données dans 3 tables :
  - `system_status`
  - `asset_pairs`
  - `ticker_info`
- Utilise **goroutines**, **waitgroups** et **channels** pour paralléliser les appels API
- Option pour automatiser l'insertion toutes les 5 minutes (en cours)

---

## Schéma de la base de données

### `system_status`
| id | unixtime | rfc1123 |
|----|----------|---------|

### `asset_pairs`
| id | pair | altname | base | quote |
|----|------|---------|------|-------|

### `ticker_info`
| id | pair | ask_price | bid_price | last_trade_price | open_price | high_24h | low_24h | volume_24h | retrieved_at |
|----|------|-----------|-----------|------------------|------------|----------|---------|------------|---------------|

---

## Lancer le projet

### 1. Cloner le repo
```bash
git clone https://github.com/ton-utilisateur/kraken-api-go.git
cd kraken-api-go
```

### 2. Lancer le projet
```bash
go run .
```

## Remarques personnelles

> J’ai voulu prendre le temps de bien structurer mon projet avec une architecture propre en séparant la base de données dans un dossier `/database`, et en utilisant les modules Go.  
> Cela m’a pris un peu de temps car les modules ne fonctionnaient pas comme prévu au début, donc j’ai finalement remis tout dans le `main.go` pour assurer la fonctionnalité principale.

> Ce qu’il me manque pour compléter totalement le projet :
> - L’export des données de `ticker_info` dans un **fichier CSV**
> - La mise en place d’un **serveur web simple** pour permettre de **récupérer ce CSV** via un endpoint HTTP

> En dehors de cela, le reste n’était pas particulièrement difficile à réaliser.
