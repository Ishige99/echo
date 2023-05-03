# echo

## 開発環境

- MacOS Monterey v12.2.1 M1
- Goland
- go version go1.20 darwin/arm64

## Local Build

- Docker version 20.10.22
- Docker Compose version v2.15.1
- MySQL v8.0.33

```bash
$ docker-compose up -d

// Operate DB
$ docker-compose exec mysql bash
bash-4.4# mysql -u root -p
> Enter password: root
```

## packages

### Gorilla

https://github.com/gorilla/mux  
  
特定のメソッドでの処理を記述したいため使用しました。  
Gorillaを使用すれば、`main.go`の`HandleFunc`をより簡潔に書けます。

## DB Scheme

### user

ユーザー情報を保存するテーブル

| Column | Type | Options |
| -- | -- | -- |
| id | BIGINT | NOT NULL, PRIMARY KEY, AUTO_INCREMENT |
| name | VARCHAR | NOT NULL |
| avatar | VARCHAR | NOT NULL |
| is_delete | BOOLEAN | NOT NULL |
| created_at | TIMESTAMP | NOT NULL, CURRENT_TIMESTAMP |
| updated_at | TIMESTAMP | NOT NULL, CURRENT_TIMESTAMP |

### post

ユーザーの投稿を保存するテーブル

| Column | Type | Options |
| -- | -- | -- |
| id | BIGINT | NOT NULL, AUTO_INCREMENT, PRIMARY KEY |
| user_id | BIGINT | NOT NULL |
| content | TEXT | NOT NULL |
| is_delete | BOOLEAN | NOT NULL |
| created_at | TIMESTAMP | NOT NULL, CURRENT_TIMESTAMP |
| updated_at | TIMESTAMP | NOT NULL, CURRENT_TIMESTAMP |