# Backgound
ログイン機能を実装したい

# Objective
トークンの検証

- トークン検証
    - 署名検証
    - 有効期限検証
    - ユーザー検証
- ログイン制約 (非ログインユーザの場合エラーを出す)

を提供する。

# Module

**`VerifyToken` (`cmd/api/server/middleware/auth.go`)**

- Authorizationヘッダーからトークンを取得する
- トークンの署名を検証（復号化）する
- claimからexpを取得し、検証する
- claimからuidを取得し、`Verify`からユーザーを検証する

**`user.Verify` (`src/user/usecase.go`)**
- uidからユーザーを検証する

# Dataflow

- Authorizationヘッダーからトークンを取得する
- `VerifyToken`にてトークンを検証する
- `VerifyToken`から`user.Verify`を呼び出して、uidからユーザーを検証する

see also [middleware](https://github.com/dev-sota/going-to-go-example/tree/main/cmd/api/middleware/auth.go)
see also [usecase](https://github.com/dev-sota/going-to-go-example/tree/main/src/user)
