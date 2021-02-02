# Backgound
ログイン機能を実装したい

# Objective
User作成時にパスワードを暗号化して保存する。

- Emailが事前に登録されていないか確認する機能
- パスワードを暗号化する機能

を提供する。

# Module

**`FindByEmail` (`pkg/domain/repository/user.go`)**

- リクエストされたEメールアドレスが事前に登録されていないか確認する。
- 登録されていない場合はnil
- 登録されていた場合は、エラーレスポンスを返す。
- データアクセスを必要とし、ログイン時にも流用できるため、repositoryとして実装。

**`password_encryption` (`src/user/usecase.go`)**

- リクエストされたパスワードをbcryptを用いてハッシュ化する。
- ビジネスロジックであり、データアクセスを必要としないため、Usecaseに実装。
- Usecase外部から呼び出されることはないため、privateな関数として実装し、Addメソッドから呼び出す。

# Dataflow

- リクエストされたEメールアドレスが事前に登録されていないか確認する
- リクエストされたパスワードを`password_encryption`に渡し、暗号化。
- `Create`のメソッドを利用し、Userを作成。

see also [usecase](https://github.com/dev-sota/going-to-go-example/tree/main/src/user)
