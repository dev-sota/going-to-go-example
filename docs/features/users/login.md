# Backgound
ログイン機能を実装したい

# Objective
ログイン時にパスワードを復号化して、ユーザー認証をする。

- Emailが事前に登録されているか確認する機能
- パスワードを復号化する機能
- JWTを発行する機能

を提供する。

# Module

**`FindByEmail` (`pkg/domain/repository/user.go`)**
- `docs/features/users/create.md`で作成。

**`password.Authorize` (`pkg/password/pasword.go`)**
- リクエストされたパスワードをbcryptを用いて復号化する。
- `password.Encrypt`同様。

**`value.NewToken(user_id)` (`pkg/value/token.go`)**
- `github.com/go-chi/jwtauth`を使用し、ユーザーIDからトークンを作成する。

# Dataflow

- リクエストされたEメールアドレスから登録されているユーザーを探す。
- 登録されているパスワードを`password.Authorize`に渡し、復号化。
- リクエストされたパスワードと一致しているか確認。
- JWTを発行して、レスポンスとして返す。

see also [usecase](https://github.com/dev-sota/going-to-go-example/tree/main/src/user)
