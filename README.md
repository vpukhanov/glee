# glee

Glee is a command-line utility that simplifies the management of Git's local exclude patterns. It provides an intuitive interface for manipulating the `.git/info/exclude` file, which functions similarly to `.gitignore` but remains entirely local and is not shared with the repository.


[Learn more about .git/info/exclude in the Git documentation](https://git-scm.com/docs/gitignore)


## Usage

```sh
# Add specific files to the exclude list
glee add filename1 filename2

# Add multiple files using a glob pattern
glee add *.txt

# Add a glob pattern itself to the exclude list (escaping the special character)
glee add \*.txt

# Display current entries in the exclude list
glee list

# Clear all entries from the exclude list
glee clear

# Open the exclude file in your default text editor
glee edit

# Display help and available commands
glee help
```

## Installation

### Using Go

If you have Go installed on your system, you can install `glee` directly using the `go install` command:

```sh
go install github.com/vpukhanov/glee@latest
```

This will download the source code, compile it, and install the `glee` binary in your `$GOPATH/bin` directory. Make sure your `$GOPATH/bin` is added to your system's `PATH` to run `glee` from anywhere.

### From Releases

1. Visit the [Releases page](https://github.com/vpukhanov/glee/releases) of the repository.
2. Download the archive for your operating system and architecture.
3. Extract the archive:
   - On macOS: Double-click the .zip file or use unzip glee\*\*\*\*.zip
   - On Linux: tar -xzf glee\*\*\*\*.tar.gz
   - On Windows: Extract the .zip file using File Explorer or a tool like 7-Zip
4. Move the `glee` binary to a directory in your system's `PATH`.

### Building from Source

To build `glee` from source:

1. Clone the repository:
   ```sh
   git clone https://github.com/vpukhanov/glee.git
   ```
2. Navigate to the project directory:
   ```sh
   cd glee
   ```
3. Build the project:
   ```sh
   go build
   ```
4. (Optional) Move the resulting `glee` binary to a directory in your `PATH`.

### Verifying Installation

After installation, you can verify that `glee` is installed correctly by running:

```sh
glee version
```

This should display the version of `glee` you have installed.

## Questions?

If you have any questions or need further clarification, feel free to open an issue for discussion.
