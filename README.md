# gopository
Simple artifact repository written in Go.

## Build
```
go get github.com/trubeck/gopository

go install github.com/trubeck/gopository
```

## Setup

1. Create folder where all artifacts should be stored. This is the `base-path`

2. Create folders for each package you want to deliver. *Note: The name of the folders
will be the names of the packages* (`package-name`)

3. The file name of the artifacts has to be the following
    ```
    <name>_<major>.<minor>.<patch>.<file ending>
    ```
    
    e.g.:
    ```
    MySoftware_1.1.0.jar
    ```
    *To prevent problems, please do not contain dots or slashes in the filename before the
     version*
     
4. Copy the artifacts into the corresponding folders to get the following structure

    ```
    <base-path>
        |- SoftwareOne
        |       |- SoftwareOne_0.0.1.exe
        |       |- SoftwareOne_0.2.3.exe
        |
        |- SoftwareTwo
                |- SoftwareTwo_1.2.1.jar
                |- SoftwareTwo_2.3.5.jar
    ```
     
5. Start the gopository

    ```
    gopository --path "<base-path>" --host "<host>" --port "<port>"
    ```
    *The Host defaults to `localhost` and the post to `8080`*
    
6. If you added new artifact, restart the gorepository.

### SSL

If you want the connection to be SSL encrypted use the arguments `--sslCert` and
`--sslKey` with the paths to the right files.

## Usage

Get a List of all package names:

```
http://<host>:<port>/packages
```

Get a List of all packages and available versions:

```
http://<host>:<port>/versions
```

Download latest artifact of a package

```
http://<host>:<port>/download/<package-name>/latest
```

Download specific version of a package

```
http://<host>:<port>/download/<package-name>/1-1-0
```

## ToDo

- [ ] Comment code
- [ ] Authentication
- [ ] Auto-reload if a new artifact is added
- [ ] Web-GUI for up- and download