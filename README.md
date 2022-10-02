# cookiecutter-go-web-app

A project template for a web application written in Go based on the [Fiber](https://github.com/gofiber/fiber) web framework. This incorporates:

* Error handling with `github.com/rs/zerlog` and `github.com/pkg/errors`
* Database migrations
* Tools to load YAML-based configurations

To use this template, you must have [Cookiecutter](https://github.com/cookiecutter/cookiecutter) installed.

```
python3 -m pip install --user --upgrade cookiecutter

cookiecutter https://github.com/codemicro/cookiecutter-go-web-app.git
```