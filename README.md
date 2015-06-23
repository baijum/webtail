# Webtail

Webtail stream file content to web.

Usage:

```bash
go get github.com/baijum/webtail
webtail -addr=:8080 [file.log]
```

The ``-addr`` option can be used to specify the host and port number.
If ``-addr`` option is not given, by default `8080` will be used.

The second argument is the path to log file that you want to stream.

If the second argument is not given, the standard input (stdin) will be used.
This helps to use `tail` together with `webtail`:

```bash
tail -f file.log | webtail
```

After running this program, you can open the URL: http://localhost:8080 in a browser.
