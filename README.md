# HTMX Demo app

This project is intened to explore the HTMX library to avoid using Javascript frameworks to create a application that can render reusable components.

Objectives:

1.  Create a small app using the Revel frame work to render the parent page server side while injecting page specific config values.
2.  Using the HTMX library to render components inside the page based on config values.
3.  Consume the eden-content api for content.
4.  Consume the eden-auth component for JWT authentication/authorization.

## Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).

### Start the web server:

revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites

## Help

- The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
- The [Revel guides](http://revel.github.io/manual/index.html).
- The [Revel sample apps](http://revel.github.io/examples/index.html).
- The [API documentation](https://godoc.org/github.com/revel/revel).
