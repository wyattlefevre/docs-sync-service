# atlas.hcl
env "docker" {
  url = "postgresql://myuser:mypassword@postgres:5432/mydb?sslmode=disable"
  migration {
    dir = "file://schema/migrations"
    format = atlas
  }
}

