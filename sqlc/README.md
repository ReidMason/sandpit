### Setup

Generate the boilerplate by running this command

```bash
sqlc generate
```

### Thoughts

Initial impressions not bad but lacking some functionality.\
Still requires you to write the SQL manually with no validation against the actual database. So if your `schema.sql` file is wrong then the queries will be too.\
No migration support out of the box which is a shame.
