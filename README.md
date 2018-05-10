#picki
This application is called *picki* and is actually just a tool to download files from an HTTP-server which offers a specific file structure.

#Usecase
picki can be used to download libraries. Yeah, that's it.

#The server
For example: You have two java libraries `foo` and `bar`. For `foo` you want to host version 1.2.3 and 1.4.5 and for `bar` you want to host the versions 0.1, 0.2 and 0.3.

The server must have this structure:

```
./
├── bar_0-1-0
│   ├── bar_0-1-0.jar
│   └── dependencies.json
├── bar_0-2-0
│   ├── bar_0-2-0.jar
│   └── dependencies.json
├── bar_0-3-0
│   ├── bar_0-3-0.jar
│   └── dependencies.json
├── foo_1-2-3
│   ├── foo_1-2-3.jar
│   └── dependencies.json
├── foo_1-4-5
│   ├── foo_1-4-5.jar
│   └── dependencies.json
└── info.txt
```

The `info.txt` file simply contains the folder list. Someone who does not know the server can look into this file an see what versions are online.
