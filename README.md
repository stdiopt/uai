# UAI

A simple cli tool to pipe into [ollama](https://ollama.com/download)


## Install

install and compile from source (go is required)

```bash
go install github.com/stdiopt/uai@latest
```

## Usage:

```bash
cat file.txt | UAI_MODEL=gemma3 uai -i describe the contents
```

```bash
export UAI_MODEL=llama3.1:8b
cat file.txt | uai -i
```

```bash
UAI_MODEL=deepseek-r1:1.5b uai ask something as arguments with no input
```
