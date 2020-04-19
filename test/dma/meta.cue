modelsets: {
  test: {
    entry: "models"

    stores: {
      test: "test"
      dev:  "dev"
    }
  }
  other: {
    entry: "others"

    stores: {
      test: "test"
      dev:  "dev"
    }
  }
}

stores: {
  test: {
    id: "test"
    type: "psql"
    version: "12.2"
  }
  dev: {
    id: "local-dev"
    type: "psql"
    version: "12.2"
  }
}
