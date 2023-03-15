# Go Layout Creator for Automated Go Application Layouts

Go Layout Creator is an efficient open-source tool that enables developers to create a Go application layout easily. The manual process of creating folders and files can be tedious and prone to errors. By automating this process, Go Layout Creator helps developers save time and effort, allowing them to concentrate on their coding tasks.

## Getting Started
To begin using Go Layout Creator, the user can install it by executing the following command in their terminal:
```sh
go get github.com/cezarovici/goLayouter
```
## Usage
To use the tool, developers should first create an input file that outlines the desired folder and file structure. Once the input file is ready, they can run a command in the terminal to generate the layout.

```sh
go-layout-creator input_file_name 
```

### Change Directory
The "Change Directory" feature in Go Layout Creator allows developers to specify the directory where they want to create their Go application layout. This directory should be a valid path on the developer's computer, and can be specified in the input file by adding an exclamation mark before the directory path.

For instance, if a developer wants to create a Go application layout in a directory named "myapp" on their Desktop, they can specify the path as follows:
```sh
! ~/Desktop/myapp
```
A relative path can also be used to specify the directory for creating the Go application layout. To create the layout in the current directory, a single period (".") can be used as the directory path in the input file.
```sh
! .
```
<br>

### Line indentation 
Line indentation is a powerful feature of Go Layout Creator. It allows users to create nested directories easily by indenting the desired directory or file with one or more spaces. The tool will automatically create the necessary subdirectories.

This feature is particularly useful when creating complex directory structures for projects such as Kubernetes manifests. For example, the following input file:
 ```sh
app1
 controllers
  deployment.yaml
  service.yaml
 models
  config.yaml
  database.yaml
 README.md
 ```
will generate the following directory structure:
```sh
└── app1
    ├── controllers
    │   ├── deployment.yaml
    │   └── service.yaml
    ├── models
    │   ├── config.yaml
    │   └── database.yaml
    └── README.md   # this needs to be implemented
```
The controllers and models directories are nested within the app1 directory. The files within these directories are also nested within their respective directories. This can be a real time-saver when working on large projects with many different files and directories.

Files do not have to be indented, instead, they can either be on the same line or indented with one or more spaces. For example:

```sh
project1
  scripts
    deploy.sh rollback.sh
  config
    dev.yamlprod.yaml
  README.md # this needs to be implemented
```
Output directory structure:
```sh
└── project1
    ├── config
    │   ├── dev.yaml
    │   └── prod.yaml
    ├── scripts
    │   ├── deploy.sh
    │   └── rollback.sh
    └── README.md # needs to be implemented
```
A directory structure for a Kubernetes application could include subdirectories for deployments, services, and volumes. For example, the deployments directory might contain YAML files defining Kubernetes deployments, while the services directory might contain YAML files defining Kubernetes services. The volumes directory could include subdirectories for different types of data, such as configuration files or persistent data storage.
```sh
myapp/
 ├─ manifests
 │   ├─ deployment.yaml
 │   ├─ service.yaml
 │   └─ ingress.yaml
 ├─ cmd
 │   └─ myapp
 │       └─ main.go
 ├─ internal
 │   └─ pkg
 │       ├─ db
 │       │   └─ db.go
 │       ├─ http
 │       │   └─ server.go
 │       └─ auth
 │           └─ auth.go
 └─ vendor
     ├─ github.com
     │   ├─ kubernetes
     │   └─ myapp
     └─ golang.org
         └─ x
             └─ errors
```
To create this directory structure with Go Layout Creator, the developer should use the following input file:
```sh
myapp
 manifests
  deployment.yaml
  service.yaml
  ingress.yaml
 cmd
  myapp
   main.go
 internal
  pkg
   db
    db.go
   http
    server.go
   auth
    auth.go
 vendor
  github.com
   kubernetes
   myapp
  golang.org
   x
    errors
```

### Package Name

One of the key features of Go Layout Creator is its ability to assign package names based on the directory structure. By default, the package name is set to the last directory name in the folder path. However, if a file named main.go is present, the package name is automatically set to "main". To set a custom package name for other files, users can add a line to the input file specifying the desired package name.


```sh
# package <package name>
```
#### Example 1: Default package name
```sh
myapp
  file1.go
  subfolder1
    file2.go
```
By default, Go Layout Creator will create the package with the name myapp and subfolder1. So the generated files will look like this:
```sh
 ├── file1.go # pacakge myapp
 └── subfolder1
    └──file2.go # package subfolder1
```
#### Example 2: Custom package name
```sh
myapp
# package mypkg
  file1.go
  subfolder1
  # package submypkg
    file2.go
```
The following output would look like:
```sh
myapp
 ├─ file1.go # pacakge mypkg
 └─ subfolder1
    └──file2.go # package submypkg
```
#### Example 3: Main package name
```sh
myapp
  main.go
```
Since the file name is main.go, Go Layout Creator will use the package name package main. So the generated files will look like this:
```sh
main
  main.go # package main
  go.mod # should be implemented
```
<br>

For a web application that uses the Gin framework and Docker. The input file might look something like this:

```sh
mywebapp
  Dockerfile
  main.go
  controllers
    users_controller.go
  models
    user.go
  routes
    routes.go
  services
    users_service.go
  tests
    controllers
      users_controller_test.go
    models
      user_test.go
```

In this case, Go Layout Creator will create the following directory structure with the corresponding packages for each file:

```sh
mywebapp/
├── controllers/
│   └── users_controller.go    # package controllers
├── Dockerfile
├── go.mod
├── main.go   # package main
├── models/
│   └── user.go    # package models
├── routes/
│   └── routes.go    # package routes
├── services/
│   └── users_service.go    # package services
└── tests/
    ├── controllers/
    │   └── users_controller_test.go    # package controllers_test
    └── models/
        └── user_test.go    # package models_test
```

#### Creating test files
For test files, go-layouter has the following syntax:
```sh
# t for simple tests
# tt for tdd
```

For this app
```sh
my_project
├── main
│   ├── main.go
│   └── main_test.go
└── lib
    ├── lib.go
    └── lib_test.go
```

The input file should be

```sh
! my_project
main
    # t
    main.go
lib
    # tt 
    lib.go
```

## Conclusion
Go Layout Creator simplifies the process of creating Go application layouts by automating the creation of folders and files, saving developers time and reducing the likelihood of errors. Its functionality includes setting package names, changing directories, creating files, and indenting lines, making it a valuable tool for developers seeking to improve their workflow.