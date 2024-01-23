# Fiber - Templ - HTMX - TailwindCSS Template
## Usage
1. Clone the repository to your local machine.
2. Ensure that you have the prerequisites installed.
3. Run `make init` to install project dependencies.
4. Use `make run` to generate templates and run the application during development in watch mode.
5. For a production build, run `make build`.

There is a Makefile designed to streamline the development, templating, and building processes for your Go project, utilizing the specified package manager. By default, the package manager is set to [pnpm](https://pnpm.io/), but you can easily change it to your preferred package manager.

## Prerequisites
Before using the Makefile, ensure that you have the following prerequisites installed on your system:
- [Go](https://golang.org/) (required for building and running the application)
- Your preferred package manager (e.g., [pnpm](https://pnpm.io/))
- [Templ](https://templ.guide/)

## Configuration
You can change the author and package name in `go.mod` and `package.json` files.
You can change the default package manager for javascript environment by modifying the `package_manager` variable at the beginning of the Makefile.
```make
package_manager = YOUR_PACKAGE_MANAGER
```
Replace `YOUR_PACKAGE_MANAGER` with the name of your preferred package manager.
You can change pretty much everything.
## Targets

### `all`
The default target is set to `build`. Running `make` or `make all` will build the Go application.

### `templ`
Generates code using [templ](https://github.com/dannyvankooten/templ) based on the templates located in the `./components` directory.

### `tailwind`
Builds Tailwind CSS using the specified package manager. It depends on the `init` target to install the required dependencies.

### `run`
Runs the Go application in templ's watch mode, detects changes and restarts accordingly.

### `build`
Builds the Go application and outputs the executable as `app`.

### `init`
Installs the project dependencies using the specified package manager.
