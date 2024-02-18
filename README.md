

# OneCV Take Home Assessment

Implementing Golang API where which teachers can use to perform administrative functions for their students. 

[_written in Golang, Gin, SQLC_]


## Installation Guide:

```bash
# for updating schemas on db
brew install goose

# start psql server
brew services start postgresql

# create database
createdb <DATABASE_MAIN>
createdb <DATABASE_TEST> 

# create .env and include variables:
# example:
# DATABASE_URL="postgresql://shawntyw:shawntyw@localhost/gt-onecv?sslmode=disable" (main db) 
# DATABASE_TEST_URL="postgresql://shawntyw:shawntyw@localhost/gt-onecv-test?sslmode=disable"(test db)
touch .env
```

### Generate DB schema & mock data
```bash
# required to update the Makefile's postgresql connection details 
make schema-up
```

### Running server
```bash
make run
```

### Running testcases
```bash
make test
```

### Drop DB schema
```bash
# required to update the Makefile's postgresql connection details 
make schema-down
```

