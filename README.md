# radiotutor
Amateur Radio training website for the UK Licences

[visit us](https://radiotutor.uk)

## Roadmap

~~Mock Tests~~ -> Login Functionality -> Course Structure -> UX

## Development Environment

If you want to develop RadioTutor, you need to do the following:

1) Set up a Go environment using your OS of choice (https://golang.org/doc/install#install)

2) Clone the RadioTutor repo: 
```
go get github.com/pe5er/radiotutor
```
3) Install the Redis package, https://redis.io/, either through your distro's package manager or by building from source

4) Start the Redis service

```
systemctl start redis.service
```
5) Build the application
```
go build
```
6) Run the application
```
./radiotutor
```

To run radiotutor in the background, use `./radiotutor &`
