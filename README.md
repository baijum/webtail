# Webtail

Webtail stream file content to web.

Usage:

```bash
webtail -addr=:8081 [file.log]
```

The ``-addr`` option can be used to specify the host and port number.
If ``-addr`` option is not given, by default `8081` will be used.

The optional argument is the path to log file that you want to stream.

If the argument is not given, the standard input (stdin) will be used.
This helps to use `tail` together with `webtail`:

```bash
tail -f file.log | webtail
```

After running this program, you can open the URL: [http://localhost:8081](http://localhost:8081) in a browser.
