# Webzen - A Go UI Framework for WebAssembly

![GitHub last commit](https://img.shields.io/github/last-commit/dimkauzh/webzen)
![GitHub license](https://img.shields.io/github/license/dimkauzh/webzen)
![Lines of code](https://tokei.rs/b1/github/dimkauzh/webzen?category=lines)

Webzen is a Go UI/Render Framework that targets WebAssembly, enabling you to build web applications with Go. It helps you build ui's in the web and render things easy and fast. This project leverages `syscall/js` to interact with the JavaScript runtime in the browser.

> **Please note that Webzen is currently under high maintenance and is not production-ready.** The project is actively being developed and improved, which is why there is only a `dev` branch available.

## Building
### How to Run the Example

To run the example provided in this repository, follow these steps:
1. Install needed dependencies:

First, make sure you have Go installed, secondly, run make install to install everything needed:

```shell
make setup
```

2. Clone this repository to your local machine:

```shell
git clone https://github.com/dimkauzh/webzen.git
```
3. Navigate to the project folder:

```shell
cd webzen
```
4. Build and run the example using make and wasmserve. This will start a local development server at localhost:8080:
```shell
make example
```
Open your web browser and go to http://localhost:8080 to see the example in action.

### Building a Ready Application
To build a production-ready WebAssembly application with Webzen, follow these steps:

Compile your Go code to WebAssembly using the GOOS=js and GOARCH=wasm flags. Replace build_path/file_name.wasm and file_path/file_name.go with your desired output file path and source file, respectively:

```shell
GOOS=js GOARCH=wasm go build -o build_path/file_name.wasm file_path/file_name.go
```
You can then include the generated .wasm file in your web application and load it using JavaScript.

Feel free to explore the example provided in this repository to better understand how Webzen works and how you can create your own Go-powered web applications.

## License
This project is licensed under the GPLv3 License - see the LICENSE file for details.
