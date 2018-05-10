# picki
This application is called *picki* and is actually just a tool to download files from an HTTP-server which offers a specific file structure.

# Usage

# Installation
Just use `pic install foo@1.2.3`.

# Configuration
The configuration is stored in a file called `pic.conf` and should be in the same folder where the `pic` command is exeuted.

The file contains the following attributes:

```
# This is a comment

# The base-url where to search
url:http://foo.com/pic

# The output folder where all libraries are saved
output_folder:./libs
```

# The server
For example: You have two java libraries `foo` and `bar`. For `foo` you want to host version 1.2.3 and 1.4.5 and for `bar` you want to host the versions 0.1, 0.2 and 0.3.

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

The `info.txt` file simply contains the folder list. Someone who does not know the server can look into this file an see what versions are online.

The `meta` file contains meta information about the library, e.g. the file extension

# Meta file
The meta file is stored next to the actual library and is just named `meta` (no file ending).

This file contains meta information about the library, e.g. the file ending. Here're all attributes:

```
# This is a comment

# The file extension of the library
ext:
```
