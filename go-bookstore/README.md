# Book store API

This is a simple book store API written inb golang.  It has 2 main packages i.e `pkg` and `cmd`

The `pkg` folder defines the functionality of the application with reference to:
1. database interaction
database models and functions are defined in the `models` folder
2. business logic 
the api functions are defined in the `controllers` folder
3. The `config` folder defines the connection with the MSQL database
4. The `routes` define the api routes and endpoints 

The `cmd` folder contains the main file used to run the application

## Installation
To install this application, first ensure you have GO installed in your system. Then run the following command:

```bash
go get github.com/Faith-qa/Learning_Go/go-bookstore
```
This will download the source code and install the application

