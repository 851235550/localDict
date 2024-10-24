# localDict

## Preparation
Install golang dependencies.
```shell
go get ./...
```

## Build binary
```go
go build
```

## How to use
```shell
./dict -e "Now"
Now
adv. 现在，此刻；马上，立刻；（用于引起注意或用于转换话题）好了；到现在为止，迄今；这样一来；（用于婉转地强调请求或命令）好了；你要知道（用于信息承上启下）；然而；用于叙述过去的事当时，那时；现在到了这个地步；嗯（在决定下面说什么时的停顿用语）；（用于讽刺性反问的结尾）哦；<文>时而
adj. 现在的；<非正式>流行的，时髦的
n. 现在，此刻
conj. 由于，既然

./dict -e "Now you have created"
Now you have created
现在你已经创建了
```

```shell
./dict --help
Usage of ./dict
  -e string
        Input an English word or sentence
```
