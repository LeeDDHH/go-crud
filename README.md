# go-crud

- Go を使って簡単な CRUD アプリを作る

## 前提

- DB を設けず、Go 内で用意したスライスに対して CRUD する

## 環境

|     |       |
| :-: | :---: |
| go  | 1.19  |
| gin | 1.8.1 |

## エンドポイント

|    パス     | メソッド |                    役割                    |
| :---------: | :------: | :----------------------------------------: |
|   /movies   |   GET    |       サーバーで管理するデータを返す       |
| /movies/:id |   GET    |         指定した id のデータを返す         |
|   /movies   |   POST   |           新しいデータを登録する           |
| /movies/:id |   PUT    | 指定した id をもとに既存のデータを更新する |
| /movies/:id |  DELETE  | 指定した id をもとに既存のデータを削除する |

## 構成

```shell
.
├── README.md
├── api APIの集約
│   └── movie.go
├── data サーバー内で扱うデータの集約
│   └── movies.go
├── go.mod
├── go.sum
├── main.go
├── router ルーティングの集約
│   └── router.go
└── types typeの集約
    └── type.go
```
