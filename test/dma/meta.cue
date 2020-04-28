modelsets: {
  test: {
    entry: "models"

    stores: {
      test: "test"
      dev:  "dev"
    }
  }
  other: {
    entry: "others" @entry(hof=geb)

    stores: {
      test: "test" @hof("geb")
      dev:  "dev"
    }
  } @hof(geb=weird) @geb()
} @model(hof=goo)

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

ex: {
  foo: attrs(modelsets)
  goo: attrs(modelsets.other)
}
