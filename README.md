# Go Application README

## 🔧 **Docker を使った環境構築手順**

1. **Docker イメージを作成する**

```bash
$ docker compose build
```

2. **スキーマグレーションの実行**

```bash
$ docker-compose run --rm -e SQLDEF_ACTION=apply mysqldef
```

3. **Docker コンテナの起動**

```bash
$ docker compose up -d
```

---
