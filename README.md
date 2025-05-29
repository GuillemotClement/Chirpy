# chirpy

Projet Bootdev d'un serveur en Go

## Lancer le server

```shell
go build -o out && ./out
```

## Lib

### Goose

(Goose)[https://github.com/pressly/goose]
Outils de migration de changement de la DB.

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Pour verifier la version installer

```shell
goose -version
```

Pour faire une migration avec Goose

```shel
goose postgres <connection_string> up

# url db
# protocol://username:password@host:port/database
```

La migration est lancer depuis le dossier du fichier sql

Dans un dossier `sql/schema`, on peut creer un fichier `001_users.sql` pour modifier la base

```sql
-- +goose Up
CREATE TABLE ...

-- +goose Down
DROP TABLE users;
```

### SQLC

Outil qui permet de generer du Go depuis les requetes SQl.

Ce n'est pas un ORM mais un outil qui permet de travailler avec des requetes SQl plus facilement et avec un typage.

```shell
brew install sqlc
```

L'outil est a lancer en etant a la racine du projet

La configuration dans un fichier `sqlc.yaml` placer a la racine du projet

```yaml
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```

Creer les requete dans `sql/queries`

Pour generer le code :

```shell
sqlc generate
```

Pour corriger le probleme de package

```shell
go get github.com/google/uuid
```

Importation du driver PG

```shell
go get github.com/lib/pq
```

Importer ensuite dans le main

```go
import _ "github.com/lib/pq"
```

### .env

Pour utiliser le .env

```shell
go get github.com/joho/godotenv
```

Puis au debut du `main.go` :

```go
godotenv.Load()
```

Pour recupere des valeur

```go
dbURL := os.GetEnv("DB_URL")
```
