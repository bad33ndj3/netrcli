# netrcli

This .netrc CLI extracts specific parts (name, login, password, account) from a `.netrc` file for a given machine.

## Installation

You can install the program using `go install`:

```sh
go install github.com/bad33ndj3/netrcli@latest
```

## Usage

The program requires three arguments: the path to the `.netrc` file, the machine name, and the part to print (name, login, password, account).

```sh
netrcli --netrc <path_to_netrc_file> --machine <machine_name> --part <part_to_print>
```

### Example

```sh
netrcli --netrc ~/.netrc --machine gitlab.com --part login
```

### Arguments

- `--netrc` : Path to the `.netrc` file (required).
- `--machine` : Machine name to extract information for (required).
- `--part` : Part to print (required). Valid values are `name`, `login`, `password`, `account`.

### Error Handling

- If any of the required arguments are missing, the program will exit with an appropriate error message.
- If the specified part is not one of `name`, `login`, `password`, `account`, the program will exit with an error.
- If the specified machine is not found in the `.netrc` file, the program will exit with an error.
