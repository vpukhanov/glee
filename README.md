# glee

This command-line utility lets you manipulate .git repository's `info/exclude` file, which works like `.gitignore` but is completely local and not included in the repository itself.

[Learn more about .git/info/exclude](https://git-scm.com/docs/gitignore)

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

    # Display help
    glee help
