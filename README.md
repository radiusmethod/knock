# knock

To run:

`go main.go`

Once started you'll see the sequence.  You can use something like netcat to trigger the sequence.

`nc -u 127.0.0.1 <portnumber>`

Once you have knocked successfully you should be able to get to http://localhost:9999.

### Future ideas:
* Use something like TOTP to generate the ports.
* Use a timeout mechanism to reset the list of knocks received.
* Perhaps use an IP address to index by, since the list is global.
* Allow the webserver to be accessed only the IP that knocked it.