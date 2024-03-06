# Offers Market

Demo project for a marketplace where you can find the best offers in town

## Build templates

To generate the template files run:

```sh
make build-templates
```

## Build styles

You'll need to build the styles each time you make a change in the styles you use in templates.
This is because Tailwind CSS doesn't include all utility classes, just what you use.

Build the styles running:

```sh
make build-styles
```

## DB Migrations

Use [Goose](https://github.com/pressly/goose) for DB Migrations.
