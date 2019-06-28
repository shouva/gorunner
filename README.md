# GoRunner
this tool to pull go project and automatic build and run it. identical to:
```
cd to\go\path\project
git pull 
#if new push there
  kill current process
  go build -o output
  ./output

```

you can setting path of js source and delay time.

## How to Use

  1. install
     ```go get github.com/shouva/gorunner```
   
  2. run it!, process will error. this normal.
  ```
  gorunner
  ```
  3. edit gorunner.json
  ```
  {
      "path":"~/codes/go/src/github.com/mosleim/gochanged",
      "delay": 10,
      "branch": "master",
      "output": "output"
  }
  ```
  4. run again!
  ```
  gorunner
  ```
  5. Happy code and running. :-)


This repository is powered by [CV. Otoritech](https://otoritech.com) and [Shouva Store](https://shouva.com).