# Fly

**Fly** is a Golang-based hot-reload tool designed to watch changes in your Go project files and automatically rebuild and restart your application. It aims to streamline development by providing fast feedback during the development cycle.

## Features

- Watches for file changes in your Go project directory.
- Automatically rebuilds and restarts the application on file changes.
- Handles system interrupts (`Ctrl + C`) and program termination signals for graceful shutdowns.
- Configurable target directory for watching file changes.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/fly.git
    cd fly
    ```

2. Build the tool:
    ```bash
    go build -o fly
    ```

## Usage

### Basic Usage
To run Fly and watch the current directory for file changes, simply execute:

```bash
./fly
```

### Specify a Target Folder

You can also specify a target folder to watch by using the `--target` flag:

```bash
./fly --target=./path/to/your/project
```

### Example

```bash
./fly --target=./src
```

This command will watch the `./src` folder and rebuild the project whenever a file changes.

## Handling Interruptions

Fly supports graceful shutdown. Press `Ctrl + C` during the execution, and Fly will stop the current process, perform cleanup, and exit safely.

## How It Works

1. Fly uses [fsnotify](https://github.com/fsnotify/fsnotify) to watch for file changes in the specified directory.
2. When a change is detected, it runs the `go build` command to rebuild your project.
3. After a successful build, Fly will restart the newly built binary.
4. If Fly is interrupted (`SIGINT`, `SIGTERM`), it cleans up by terminating the currently running process and then exits.

## Contributing

Contributions are welcome! If you find a bug or have an idea for a feature, feel free to submit an issue or a pull request.

1. Fork the repository
2. Create a new feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add your feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Notes:
- The `Installation` and `Usage` sections explain how to build and run the tool.
- The `Handling Interruptions` section explains how the program handles `Ctrl + C` signals.
- The `How It Works` section provides a high-level explanation of the tool's functionality.
- There's a placeholder for a `LICENSE` section, which can be replaced with the actual license file if needed.