A Golang worker that fetches the current $ cotization every hour and stores it within a PostgreSQL database

TODO:

- ~~Add PostgreSQL and insert of the data~~
- ~~Add a simple server with an endpoint that returns all the data available~~
- ~~Add query to fetch data for a chart~~
- ~~Add cors~~
- Add tests for everything
- Add viper to configure environment variables

I'm adding the structure for heroku deployments following these steps:

http://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html

Still not sure how to work with migrations in golang, lol, just loaded the db dump with:

```
cat db.dump | heroku pg:psql
```
