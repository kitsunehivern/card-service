env "dev" {
  url = "postgres://user:pass@localhost:5432/carddb?sslmode=disable"
  schemas = ["public"]
}

env "local" {
  url = "postgres://user:pass@localhost:5432/carddb?sslmode=disable"
  schemas = ["public"]
  migration {
    dir = "file://migration"
  }
}
