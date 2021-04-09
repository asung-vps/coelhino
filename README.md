# Starfire

## TODO:
- Implement init command:
  - Create TOML project file in optionally specified directory
  - Want to set an initConfig function to detect project, but need to figure out how to handle uninitialized project case.
    - Could have a struct that parses TOML file and a flag indicating whether the context is initialized or not
- Need to determine differentiation between fmt.Println and log.Println

## Example of a transaction
```toml
description = " This is a test transaction"

[configuration]
default-protocol = "https"

[[requests]]
name = "request1"
description = "This is a test request"
protocol = "http"
url = "google.com"
path = "/"
type = "GET"
headers = [
	["Accept-Language", "en-US,en;q=0.5"],
	["Accept-Encoding", "gzip, deflate"]
]

```