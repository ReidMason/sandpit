# Create migration

Install the golang-migrate CLI tool

```bash
brew install golang-migrate
```

Create the migration

- `-ext` specifies the file extension
- `-dir` is the output directory
- `-seq` uses sequential numbers instead of timestamps

```bash
migrate create -ext sql -dir db/migrations -seq init
```
