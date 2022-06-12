
<br />
<div align="center">
 

<h3 align="center">GoInject</h3>

  <p align="center">
    GoInject is a dynamic-link library injection program written in Go.
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

![Alt text](/docs/goinject.PNG "GoInject")

GoInject is a dynamic-link library injector that is written in Go. It makes use of Windows system calls and wrappers. The graphical user interface is based on the Fyne toolkit.

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Go](https://go.dev/)
* [Fyne](https://fyne.io/)


<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

To build the executable from source, follow these steps.


### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/tjandy98/goinject.git
   ```
3. Install dependencies
   ```sh
   go mod tidy
   ```
4. Generate executable binary
   ```sh
   go build 
   ```

   > To build for 32-bit and 64-bit, set the environment variable `GOARCH` to `amd64` for 64-bit and `386` for 32-bit.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

A target process should be selected from the list of active processes. Upon selecting a process, the target process name and the target process ID will be displayed.
A valid DLL path should be set either by manually entering the path, or using the built-in file dialog.

Finally, the DLL can be injected into the target process by clicking on the Inject button.



<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap
- [ ] Add process filter
- [ ] Save configurations
- [ ] Add more injection techniques



<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request.
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. 






<p align="right">(<a href="#top">back to top</a>)</p>


