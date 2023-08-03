# RESTWrench

## Install
### To get both submodules use    
```$ git clone --recurse-submodules git@github.com:gstrand99/RESTWrench.git ```
### To update those submodules use
```$ git submodule update --remote ```

## Commands
### Main tool 
```$ wrench```
### Create the queries and tests
| Name | Command | Description |
| ------- | ----------- | ------------- |
| Main: | ```$ wrench create``` | |
| POST: | ```$ wrench create post``` | |
| GET: | ```$ wrench create get``` | |
| PUT: | ```$ wrench create put``` | |
| DELETE: | ```$ wrench create delete``` | |
| Test: | ```$ wrench create test``` | |
### Run the queries and tests
| Name | Command | Description |
| ------- | ----------- | ------------- |
| Main: | ```$ wrench run``` | |
| POST: | ```$ wrench run post``` | |
| GET: | ```$ wrench run get``` | |
| PUT: | ```$ wrench run put``` | |
| DELETE: | ```$ wrench run delete``` | |
| Test: | ```$ wrench run test``` | |
| All: | ```$ wrench run all``` | |
### Set your config
| Name | Command | Description |
| ------- | ----------- | ------------- |
| Main: | ```$ wrench config``` | |
| Create: | ```$ wrench config create``` | |
| Set: | ```$ wrench config set``` | |
| Main: | ```$ wrench config``` | |


### Structure
```
RESTWrench/
├── wrench/         # Core functionality
│   ├── crud/       # CRUD operations
│   ├── config/     # Configuration management
│   └── test/       # Test running
├── tui/            # TUI interface
│   └── main.go     # Entry point for the TUI application
├── cli/            # CLI interface
│   └── main.go     # Entry point for the CLI application
└── nvim/           # Neovim plugin
    └── main.lua    # Entry point for the Neovim plugin
```
