modelsets: {
  test: {
    models: {
      entry: "models"
    }

    stores: {
      test: "test" @reln("many-to-many=foo")
      dev:  "dev"
    }
  }
} @attrtest("modelset=test")

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
} @attrtest("stores='test'") @foogoo("cow='moo',cat='meow'")
