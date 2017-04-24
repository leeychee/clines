Clines
============================================================

### Install

`go get -u github.com/leeychee/clines`

### Usage

1. 单次生成

`./clines -s 12 -r 500 -o out 1,2 3,4`

2. 批量生成

`./clines -s 12 -r 500  -f imgs.txt`

3. 说明

```
-f string
  	定义连线的文件路径，当此参数为空时，将从命令行参数读取
-o string
  	输出图片路径 (default "output")
-r float
  	圆半径，间接指定了图片大小 (default 500)
-s int
  	将圆按指定数字等分 (default 12)
```
