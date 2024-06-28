# glee

This command-line utility lets you manipulate .git repository's `info/exclude` file, which works like `.gitignore` but is completely local and not included in the repository itself.

[Learn more about .git/info/exclude](https://git-scm.com/docs/gitignore)

## Installation

### Using Go

If you have Go installed on your system, you can install `glee` directly using the `go install` command:

```sh
go install github.com/vpukhanov/glee@latest
```

This will download the source code, compile it, and install the `glee` binary in your `$GOPATH/bin` directory. Make sure your `$GOPATH/bin` is added to your system's `PATH` to run `glee` from anywhere.

### From Releases

1. Visit the [Releases page](https://github.com/vpukhanov/glee/releases) of the glee repository.
2. Download the archive for your operating system and architecture.
3. Extract the archive:
   - On macOS/Linux: `tar -xzf glee_<version>_<os>_<arch>.tar.gz`
   - On Windows: Use a tool like 7-Zip to extract the `.zip` file
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

## Usage

    # Add files to exclude list
    glee add filename1 filename2

    # Add multiple files to exclude list using glob pattern
    glee add filename*.txt

    # Add glob pattern itself to exclude list
    glee add filename\*.txt

    # Display entries in the exclude list
    glee list

    # Clear the exclude list
    glee clear

    # Open exclude file in the text editor
    glee edit

    # Display help
    glee help

```

```

```

```
