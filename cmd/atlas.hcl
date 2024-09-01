data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "./migrate",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://qingshaner:123456@127.0.0.1:5432/gin"
  migration {
    dir = "file://./migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
