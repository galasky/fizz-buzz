# Welcome to fizz-buzz

## Dependencies

You need to install the dependicies bellow in order to running the fizz-buzz project.

### Go

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

You need to [install Go language on your system](https://golang.org/dl/).

Then follow the [installation instructions](https://golang.org/doc/install#install).

## Getting Started


### Get the Revel framework

To get the Revel framework, run:

    go get -u github.com/revel/cmd/revel

Don't forget to set the environment variable PATH to your $GOPATH/bin folder in order to be able to execute the revel command line (restart your terminal after that).

### Clone the web server:

Go to your go workspace and create this path : src/github.com/galasky

Then go to the galasky folder

To clone the fizz-buzz run this command line:

    git clone https://github.com/galasky/fizz-buzz.git

### Start the web server:

To start the fizz-buzz run this command line:

    revel run github.com/galasky/fizz-buzz

Go to [http://localhost:9000/](http://localhost:9000/) and you'll see the web site.

### Description of Contents

The default directory structure of a generated Revel application:

    fizz-buzz               App root
      app               App sources
        controllers     App controllers
          init.go       Interceptor registration
        models          App domain models
        routes          Reverse routes (generated code)
        views           Templates
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
      messages          Message files
      public            Public assets
        css             CSS files
        js              Javascript files
        images          Image files

app

    The app directory contains the source code and templates for your application.

function

    The code the endpoint of fizz-buzz is in this folder : fizz-buzz/app/controllers/app.go

conf

    The conf directory contains the application’s configuration files. There are two main configuration files:

    * app.conf, the main configuration file for the application, which contains standard configuration parameters
    * routes, the routes definition file.


messages

    The messages directory contains all localized message files.

public

    Resources stored in the public directory are static assets that are served directly by the Web server. Typically it is split into three standard sub-directories for images, CSS stylesheets and JavaScript files.

    The names of these directories may be anything; the developer need only update the routes.

test

    Tests are kept in the tests directory. Revel provides a testing framework that makes it easy to write and run functional tests against your application.

### Follow the Go guidelines:

* The README file created within your application.
* The [Getting Started with Revel](http://revel.github.io/tutorial/index.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/samples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).
