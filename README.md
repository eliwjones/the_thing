This is The Thing.

####Clone the repository:
```
git clone git@github.com:eliwjones/the_thing.git
```

####Run The Server:
```
go run go/the_server.go
```

####Process 'a_key' with The Script:
```
python python/the_script.py a_key
```

####Run Log View server:
```
while true; do { echo -e 'HTTP/1.1 200 OK\r\n'; cat html/the_html.html; } | nc -l 10065; done
```

####View The Log:
http://127.0.0.1:10065

Is The Thing working?
