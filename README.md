# Revel-Gorm-Todo

Todo Application API on Go(revel) and Gorm.

## Getting Started

A high-productivity web framework for the [Go language](http://www.golang.org/).

### Start the web server:

    revel run github.com/cncgl/revel-gorm-todo

   Run with <tt>--help</tt> for options.

### API

index
```
$ curl http://localhost:9000/api/todos
```

show
```
$ curl http://localhost:9000/api/todos/:id
```

create
```
$ curl http://localhost:9000/api/todos -H "Content-type: application/json" -X POST -d '{"status":true, "title":"Shopping"}'
```

update
```
$ curl http://localhost:9000/api/todos/:id -H "Content-type: application/json" -X PUT -d '{"status":true, "title":"Meeting"}'
```

delete
```
$ curl http://localhost:9000/api/todos/:id -X DELETE
```

### Follow the guidelines to start developing your application:

* The README file created within your application.
* The [Getting Started with Revel](http://revel.github.io/tutorial/index.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/samples/index.html).
* The [API documentation](http://revel.github.io/docs/godoc/index.html).

## Contributing
We encourage you to contribute to Revel! Please check out the [Contributing to Revel
guide](https://github.com/revel/revel/blob/master/CONTRIBUTING.md) for guidelines about how
to proceed. [Join us](https://groups.google.com/forum/#!forum/revel-framework)!

## License
MIT