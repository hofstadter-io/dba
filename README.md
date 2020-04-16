# dba - a datastore assistant.

`dba` is a tool to help you manage your data
schemas, migrations, servers, and more.

Features:

- Schemas written in Cue. All the power of Cue and `dba` will turn into SQL,libraries, and tools.
- Migration calculation and management. `dba` tracks how your schema changes and handles the rest.
- Manage local, test, cicd, and production servers.
- Visual tool for interacting.
- Supports [list of databases]


# Notes and references:

SQL Drivers / Libraries:

- https://pkg.go.dev/mod/github.com/lib/pq
- GORM
- https://github.com/jackc/pgx
- https://github.com/xo/dburl
- https://awesome-go.com/#database
- https://awesome-go.com/#database-drivers
- https://awesome-go.com/#orm

SQL Builders:

- https://github.com/go-jet/jet
- https://github.com/jirfag/go-queryset
- https://github.com/xo/xo

SQL Parsers:

- https://github.com/lfittl/pg_query_go
- https://gowalker.org/github.com/cockroachdb/cockroach/pkg/sql/parser
- https://marianogappa.github.io/software/2019/06/05/lets-build-a-sql-parser-in-go/
- https://github.com/pingcap/parser
- https://github.com/xwb1989/sqlparser

SQL -> Go:

- https://github.com/xo/xo
- https://github.com/kyleconroy/sqlc

Migration management:

- https://www.prisma.io/docs/reference/tools-and-interfaces/prisma-migrate
- https://github.com/golang-migrate/migrate
- https://github.com/sequelize/umzug
- https://awesome-go.com/#database
- https://awesome-go.com/#orm

Other:

- https://github.com/vitessio/vitess
- https://github.com/google/go-cloud
- https://github.com/google/wire
- https://github.com/Azure/autorest
