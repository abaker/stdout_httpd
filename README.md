`tail -n 1` of stdout as an HTTP server

### WTF?

I wanted to turn [rtlamr](https://github.com/bemasher/rtlamr) into a RESTful sensor for Home Assistant

### Contrived Example

```
while true; do date; sleep 1; done | stdout_httpd
```

Then `curl localhost` to print the date

### Usage

```
$ stdout_httpd -h
Usage of stdout_httpd:
  -debug
        enable debug logging
  -port int
        http port (default 80)
```
