# sensitive-words-match

支持三种匹配方式
1. 敏感词
   使用DFA匹配算法；不支持标点符号；
3. 组合词的方式
   `加你|微信` 表示同时包含 `加你` 和 `微信` 才会命中；
   `象牙#猛犸象牙` 包含`象牙`，同时不包含`猛犸象牙`时才会命中；其中#后面的一个词必须包含前面一个词；
5. 正则方式
   以reg@打头，后面跟正则字符串；
### Installation

```sh 
go get https://github.com/dongweifly/sensitive-words-match
```

### Usage

```go
	service := match.NewMatchService()
	service.Build([]string{
		"fuck",
		"加你|微",
		"象牙#猛犸象牙#非象牙#象牙不准卖",
		`reg@(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}`,
	})

	fmt.Println(service.Match("阿斯蒂芬fuck", '*'))
	fmt.Println(service.Match("Hello，微信，加你吗", '*'))
	fmt.Println(service.Match("Hello, 加你吗, 微信", '*'))
	fmt.Println(service.Match("买象牙吗", '*'))
	fmt.Println(service.Match("猛犸象牙筷子", '*'))
	fmt.Println(service.Match("电话号码:15210753706@@@@@", '*'))
}
```

## RoadMap


## Contact
