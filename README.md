# Outdated README!
This version of the README is for the old bash-script. I'm currently rewriting picl in go. The file may change until I'm finished with the reimplementation.

# picl
Picl (pick library) is just a dumb tool to download libraries of a specific version.

The idea is that a simple HTTP-server with a certain file structure is able to server versioned libraries.

I wrote this to not have the overhead of Maven, Gradle, ... when I just want to download libraries and update them.

# Usage

Full help-message at the end of this file.

## Installing a library
Just use `picl install foo@1.2.3`.

## Removing a library
Just use `picl remove foo@1.2.3`.

## Configuration (local)
When executing, picl searches for a file called "picl.conf" in the execution directory.
Use the option `-c other.conf`/`--config=other.conf` to use a non-default config file.

The file contains the following attributes:

```
# This is a comment

# The base-url where to search
url:http://foo.com/picl

# The output folder where all libraries are saved
output_folder:./libs
```

The default folder (if not in the config or CLI specified) is `./libs` (so the `libs` folder on the execution directory).

# The server

## What do you need
An HTTP-server with optional write access (via FTP or such) if you want to add/remove files ;)

## Set up
For example: You have two java libraries `foo` and `bar`. For `foo` you want to host version `1.2.3` and `1.4.5` and for `bar` you want to host the versions `0.1`, `0.2` and `0.3`.

The server must have this structure:

```
./
├── bar@0.1.0
│   ├── bar@0.1.0.jar
│   └── meta
├── bar@0.2.0
│   ├── bar@0.2.0.jar
│   └── meta
├── bar@0.3.0
│   ├── bar@0.3.0.jar
│   └── meta
├── foo@1.2.3
│   ├── foo@1.2.3.jar
│   └── meta
├── foo@1.4.5
│   ├── foo@1.4.5.jar
│   └── meta
└── info.txt
```

The `info.txt` file is optional, not used by picl and simply contains the folder list. Someone who does not know the server can look into this file an see what versions are online.

The `meta` file *must* be inside a libraries folder and contains meta information about the library, e.g. the file extension.

# Meta file
The meta file has to be stored next to the actual library and is just named `meta` (no file ending).

This file contains meta information about the library, e.g. the file ending. Here're all attributes:

```
# This is a comment

# The file extension of the library
ext:jar
```
The `ext` value should be defined, otherwise picl tries to load a file without ending (probably not what you want).

# --help
```
Usage: picl [OPTIONS...] {COMMAND} [OPTIONS...] {LIBRARY}

Loads and manages simple libraries and files.

Normal options:
  -h, --help      Shows this help message

Commands:
  install         Installs the given library
  remove          Uninstalls/removes the given library

Options for the "install" command (all of these are optional):
  -h, --help      Shows this help message
  -c, --config    Specifies the configuration file that should be used. This
                  is "./picl.conf" by default.
  -o, --output    Specifies the output folder where all libraries should be
                  stored. This is "./libs" by default.
  -u, --url       The base url where picl downloads files from.

Options for the "remove" command (all of these are optional):
  -h, --help      Shows this help message

Library name at the end:
  This name if the library name inclusing the version you wan't do deal with.
  Ths name has the following format:

      my-library@3.5.1

  There must be a name and there must be a version. The version is basically
  the string that is behind the "@" and is not parsed. It just has to exist
  on the server and the format "x.y.z" (e.g. 3.5.1) is recommended.
```
