# Go Layout Creator for Automated Go Application Layouts

Go Layout Creator is an efficient open-source tool that allows developers to create a Go application layout with ease. The process of manually creating folders and files can be tedious and error-prone. By automating this process, Go Layout Creator helps developers save time and effort, and focus on what they do best - coding.

## Getting Started
To get started with Go Layout Creator, you can install it by running the following command in your terminal:
```sh
go get github.com/cezarovici/goLayouter
# not working yet
```
## Usage
To use Go Layout Creator, create an input file with the desired folder and file structure, then run the following command:

```sh
# not working, but good idea
go-layout-creator -i input_file_name -o output_folder_name
```

### Change Directory
You can use the "Change Directory" feature in Go Layout Creator to specify the directory in which you want to create your Go application layout. The directory should be a valid path on your computer, and you can specify it in the input file by prefixing the directory path with an exclamation mark.

For example, if you want to create the Go application layout in a directory called "myapp" on your Desktop, you can specify the path as follows:
```sh
! ~/Desktop/myapp
```
Alternatively, you can use a relative path to specify the directory. For example, if you want to create the Go application layout in the current directory, you can simply use a single period (".") as follows:
```sh
! .
```
Once you've specified the directory, Go Layout Creator will create the Go application layout in that directory.
<br>

### Line indentation 
Line indentation is a powerful feature of Go Layout Creator that allows you to create nested directories with ease. All you need to do is indent the desired directory or file with one or more spaces, and Go Layout Creator will automatically create the necessary subdirectories.

This is particularly helpful when creating complex directory structures for projects such as Kubernetes manifests. For example, the following input file:
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
As you can see, the controllers and models directories are nested within the app1 directory. The files within these directories are also nested within their respective directories. This can be a real time-saver when working on large projects with many different files and directories.

Note that files do not have to be indented, instead, they can either be on the same line or indented with one or more spaces. For example:

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

Now, let's consider a more complex example using Kubernetes directory structure. Suppose you want to create a Kubernetes application with the following directory structure:
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
To create this directory structure with Go Layout Creator, you would use the following input file:
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

One of the essential features of Go Layout Creator is its package name functionality. By default, Go Layout Creator assigns the package name as the last directory name in the folder path. However, if you create a file called main.go, the package name is automatically set to package main. To specify a different package name for other files, simply add the following line to the input file:
```sh
# package <package name>
```
#### Example 1: Default package name
Suppose you have the following input file:
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
Suppose you have the following input file:
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
Suppose you have the following input file:
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

Suppose you want to create the file structure for a web application that uses the Gin framework and Docker. Your input file might look something like this:

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
For test files, use the following syntax:
```sh
# t for simple tests
# tt for tdd
```
If you want to create a folder structure like this:
```sh
my_project
├── main
│   ├── main.go
│   └── main_test.go
└── lib
    ├── lib.go
    └── lib_test.go
```

You can specify it in the input file as follows:

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
Go Layout Creator simplifies the process of creating Go application layouts. By automating the process of creating folders and files, developers can save time and reduce the likelihood of errors. With its package name, change directory, file creation, and line indentation functionality, Go Layout Creator is a valuable tool for developers looking to improve their workflow.