# Revel Example

### Start the web server:

   revel run

### Go to http://localhost:9000/ and you'll see:

    "It works"

### Create user

```zsh
% curl -v -XPOST -F "name=example 1" localhost:9000/users
% curl -v -XPOST -F "name=example 2" localhost:9000/users
% curl -v -XPOST -F "name=example 3" localhost:9000/users
```

### Get users

```zsh
% curl -s http://localhost:9000/users | jq .
[
  {
    "UserId": 1,
    "Name": "example 1"
  },
  {
    "UserId": 2,
    "Name": "example 2"
  },
  {
    "UserId": 3,
    "Name": "example 3"
  }
]
```

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        models/
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites
