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
