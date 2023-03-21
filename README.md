# 科大讯飞api

```shell
go get -u github.com/fzkun/xfyun
```
```shell
emotion, err := NewXFYun().GetNaturalLanguage(&config.Config{
    AppID:  "",
    ApiKey: "",
    }).Emotion("好开心")
```