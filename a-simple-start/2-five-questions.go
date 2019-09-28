/*

	Five basics question ?

		- How do we run the code in our project ?
			- Use go commande in terminal
				- go run filename.go
				- go run main.go (1-hellow-world.go)

		Go CLI

			- go build 		-> Compiles a bunch of go source code files
			- go run 			-> Compiles and executes one or two files
			- go fmt 			-> Formats all the code in each file on the current directory
			- go install 	-> Compiles and "installs" a package.
			- go get 			-> Downloads the raw source dode of someone else's package
			- go test 		-> Runs any tests associated with the current project



		- What does 'package main' mean ?
			- Package == Project == Workspace

			2 Differents type of packages
				- Executable 	-> Generates a file that we can run
				- Reusable		-> Code used as 'helpers'. Good place to put reusable logic. (Like libraries)

		-	How do we know when we are making one or the other?
			-	Package main -> go build -> main.exe (If we ran this file the function named 'main' would be automatically ran)

			Executable package -> package main -> Defines a package that can be compiled and then *executed*.
				Must have a func called 'main'


				Reusable package 	-> package calculator -> Defines a package that can be used as dependency (helper code)
													-> package uploader		-> Defines a package that can be used as a dependency (helper code)


		- What does 'import "fmt" mean ?
		- What is that 'func' thing ?
		- How is the main.go file organized ?
*/
