# gotth-go-tailwind-templ-htmx-echo

## Introduction

The tech stack :

- Go
- HTMX
- Templ
- Tailwindcss
- Echo

## Core Technologies

- [Go](https://go.dev/) - Version 1.22.0 or greater required
- [Templ](https://templ.guide/)
- [Air](https://github.com/cosmtrek/air)
- [Htmx](https://htmx.org/)
- [Tailwindcss](https://tailwindcss.com/)
- [Echo](https://echo.labstack.com/) - Version 1.22.0 or greater required

## Installation

### Clone the Repository

```bash
git clone https://github.com/marishlucas/gotth-go-tailwind-templ-htmx-echo <destination-directory>
```

```bash
cd <destination-directory>
```

### Install Dependencies

```bash
go get -u (optional, upgrade all dependecies)
go mod tidy
```

### Change .env.example file to .env and include a PORT variable

```bash
mv .env.example .env;
```

```bash
echo "PORT=8080" > .env
```

## Build Steps and Serving

This project requires a build step. The following are commands needed to build your html and css output.

### Templ HTML Generation

With templ installed and the binary somewhere on your PATH, run the following to generate your HTML components and templates and start the proxy server.

> **_NOTE:_** Replace 8080 with your port number (or source it from the .env through your preffered method).

```bash
templ generate --watch --proxy="http://localhost:8080" --cmd="air"
```

### CSS File Generation

With the [Tailwind Binary](https://tailwindcss.com/blog/standalone-cli) installed and moved somewhere on your PATH, run the following to generate your CSS output for your tailwind classes (remove --watch to simply build and not hot reload)

```bash
tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
```

### Serving with Air

With the [Air Binary](https://github.com/cosmtrek/air) installed and moved somewhere on your PATH, the first command (templ generate) will generate the HTML components and templates and then automatically run air to start the server.

To configure air, you can modify .air.toml in the root of the project. (it will be auto-generated after the first time you run air in your repo)

## Project Overview

Note, htmx and your tailwind output are included in the head of this template:

```html
<script src="https://unpkg.com/htmx.org@2.0.1"></script>
<link rel="stylesheet" href="/static/css/output.css"></link>
```

### Inspired from

[templ-quickstart](https://github.com/phillip-england/templ-quickstart)
