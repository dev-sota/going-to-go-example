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

**`password_decoder` (`src/user/usecase.go`)**

- リクエストされたパスワードをbcryptを用いて復号化する。
- ビジネスロジックであり、データアクセスを必要としないため、Usecaseに実装。
- Usecase外部から呼び出されることはないため、privateな関数として実装し、Addメソッドから呼び出す。

# Dataflow

- リクエストされたEメールアドレスから登録されているユーザーを探す。
- 登録されているパスワードを`password_decoder`に渡し、復号化。
- リクエストされたパスワードと一致しているか確認。
- `github.com/go-chi/jwtauth`を使用し、JWTを発行して、レスポンスとして返す。

see also [usecase](https://github.com/dev-sota/going-to-go-example/tree/main/src/user)
