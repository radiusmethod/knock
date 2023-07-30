# knock

To run:

`go run main.go`

Once started you'll see the sequence.  You can use something like netcat to trigger the sequence.

`nc -u 127.0.0.1 <portnumber>`

Here is how to do 3 knocks in sequence (`arr` is the knock port sequence):

```shell
arr=(38752 39259 31959); for i in "${arr[@]}"; do echo "ping" | nc -u -w0 127.0.0.1 $i; done
```

Once you have knocked successfully you should be able to get to http://localhost:9999.

### Future ideas:
* Use something like TOTP to generate the ports.
* Perhaps use an IP address to index by, since the list is global.
* Allow the webserver to be accessed only the IP that knocked it.
* Disconnect the webserver after a period of time