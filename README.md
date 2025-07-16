# PKNeuron 🧠

![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)
![golang version](https://img.shields.io/github/go-mod/go-version/pezhmankasraee/pkneuron)
[![GitHub release](https://img.shields.io/github/v/release/pezhmankasraee/pkneuron)](https://github.com/pezhmankasraee/pksetdev/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/pezhmankasraee/pkneuron.svg)](https://pkg.go.dev/github.com/pezhmankasraee/pkneuron)

A powerful and flexible command-line application designed to streamline your interactions with Large Language Models (LLMs). Written in Go, PKNeuron provides a general-purpose interface for sending prompts and receiving responses, making it an essential tool for developers, researchers, and anyone working with AI models.

## ✨ Features

*   **General-Purpose Prompting**: Interact with various LLMs using a unified and straightforward command-line interface.
*   **Go-Powered Performance**: Enjoy fast and efficient execution thanks to its compiled Go backend.
*   **User-Friendly Interface**: Designed for ease of use, allowing you to quickly send prompts and get responses without complex setups.
*   **Extensible**: Built with flexibility in mind, aiming to support multiple LLM providers and future features.

## 🚀 Installation

### Prerequisites

*   Go 1.18 or higher

### From Source

1.  **Clone the repository:**
    bash
    git clone https://github.com/PezhmanKasraee/PKNeuron.git # Adjust with your actual repo URL
    cd PKNeuron
    

2.  **Build and install:**
    To build an executable in your current directory:
    bash
    go build -o pkneuron ./cmd/pkneuron # Assuming your main package is in cmd/pkneuron

    You can also use the `makefile` as follows:
    ```bash
    $ make build
    ```
    the binary will be created at `bin/`
    
    Or, to install  directly into your  (or ):
    bash
    go install ./cmd/pkneuron # Again, assuming main package path
    

### Pre-compiled Binaries (Coming Soon)

Pre-compiled binaries for various operating systems will be made available on the [releases page](https://github.com/PezhmanKasraee/PKNeuron/releases) (link will become active once releases exist).

## 💡 Usage

PKNeuron is designed to be intuitive. Here's how you can get started:

### Basic Prompting

To send a prompt to the default configured LLM:

bash
pkneuron "Tell me a short story about a brave knight and a wise dragon."


### Specifying Models (Coming soon)

PKNeuron will likely support specifying different models or providers via flags or configuration:

bash
pkneuron --model openai-gpt4 "Write a Python function to calculate the factorial of a number."
pkneuron --provider anthropic "Summarize the key findings of the latest IPCC report."


### Getting Help

To see all available commands and options:

bash
pkneuron --help


## ⚙️ Configuration

PKNeuron typically uses environment variables for sensitive information like API keys. Ensure you have the necessary keys set up for the LLM providers you intend to use.

For example, for OpenAI:

bash
export GEMINI_API_KEY="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"


Refer to the official documentation for specific LLM providers for details on obtaining API keys.

## 🤝 Contributing

Contributions are welcome! If you'd like to contribute, please fork the repository and open a pull request. For major changes, please open an issue first to discuss what you would like to change.

## 📄 License

PKNeuron is distributed under the [BSD 3-Clause License](LICENSE). See the  file for more details.

## 📧 About the Author

PKNeuron is developed and maintained by **Pezhman Kasraee**.

*   **Email**: github@pezhmankasraee.com
*   **GitHub**: [Pezhman Kasraee](https://github.com/PezhmanKasraee) (Adjust this link to your actual GitHub profile)