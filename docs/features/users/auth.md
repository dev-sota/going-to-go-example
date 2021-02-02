# Backgound
ログイン機能を実装したい

# Objective
トークンの検証

- Authorizationヘッダーからトークンを取得する
- トークンをmiddleware検証してユーザーを識別する
- ログイン制約 (非ログインユーザの場合エラーを出す)

を提供する。

# Module

**`JwtAuthenticator` (`cmd/api/server/middleware/auth.go`)**

- Authorizationヘッダーからトークンを取得する
- 取得したトークンを"github.com/go-chi/jwtauth" を用いて検証

**`Verify` (`src/user/usecase.go`)**
- トークン情報(user_id)からユーザー検証する

# Dataflow

- ヘッダーからトークンを取得してコンテキストに追加する
- `JwtAuthenticator`を用いてトークンを検証する
- `JwtAuthenticator`から`Verify`を呼び出して、トークン情報(user_id)からユーザー検証する

see also [middleware](https://github.com/dev-sota/going-to-go-example/tree/main/cmd/api/middleware/auth.go)
see also [usecase](https://github.com/dev-sota/going-to-go-example/tree/main/src/user)
