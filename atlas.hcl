# atlas.hcl
env "docker" {
  url = "postgresql://myuser:mypassword@localhost:1234/mydb?sslmode=disable"
  migration {
    dir = "file://migrations"
    format = atlas
  }
}

