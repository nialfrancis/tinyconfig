# TinyConfig
A bare minimum YAML config file loader module for go.

## How to use
* `GET`、`POST`、`PUT`、`DELETE`（Common HTTP methods）
* `go get github.com/nialfrancis/tinyconfig`
* Set up your config file. Eg:
```
---

- variable_one: 'test'
- variable_two: 'test2'
```
* Import it
```
import (
  "github.com/nialfrancis/tinyconfig"
)
```
* Call it with a parameter being whichever directory you've put it in or the code below to load from the same directory as the executable
```
pwd, _ = filepath.Abs(filepath.Dir(os.Args[0]))
config := tinyconfig.ReadConfig(pwd)
```
