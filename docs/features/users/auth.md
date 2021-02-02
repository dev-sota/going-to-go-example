# Backgound
ログイン機能を実装したい

# Objective
トークンの検証

- Authorizationヘッダーからトークンを取得する
- トークンをmiddleware検証してユーザーを識別する
- ログイン制約 (非ログインユーザの場合エラーを出す)

を提供する。

# Module

**`cmd/api/server/middleware/auth.go`**

- ヘッダにあるトークンを取得してコンテキストに追加する
- 取得したトークンをデコード
- デコードしたトークン情報からユーザー検証する

# Dataflow

- ヘッダにあるトークンを取得してコンテキストに追加する
- トークンからユーザー検証して認可する

see also [middleware](https://github.com/dev-sota/going-to-go-example/tree/main/cmd/api/middleware/auth.go)
