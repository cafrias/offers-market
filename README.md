# Offers Market

Demo project for a marketplace where you can find the best offers in town

## Dev Environment

You should have installed:

- [ Goose ](https://github.com/pressly/goose): for running migrations
- [ Air ](https://github.com/cosmtrek/air): for dev reloading
- [ templ ](https://github.com/a-h/templ): for generating the templates
- [ node ](https://github.com/nvm-sh/nvm): to install tailwind css
- [ protoc ](https://grpc.io/docs/protoc-installation): to build protobuffer files

## Build templates

To generate the template files run:

```sh
make build-templates
```

## Build Protobuffer files

Run:

```sh
make build-proto
```

## Build styles

You'll need to build the styles each time you make a change in the styles you use in templates.
This is because Tailwind CSS doesn't include all utility classes, just what you use.

Build the styles running:

```sh
make build-styles
```
