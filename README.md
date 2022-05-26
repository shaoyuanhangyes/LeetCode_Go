# LeetCode -Golang

LeetCode题解 Golang版本

之前刷题都是用cpp做题的 学习了go语言后 为了更加快速的熟悉go语言的语法 将之前做过的题目用golang重构一下

很多语法和cpp完全不同 经常写混 所以以下总写错的语法需要重复记忆

## 重复记忆内容

函数声明方式  函数名(变量名 变量类型, 变量名 变量类型) 函数返回值类型

if后直接写判断条件 不用加括号  if len == 0 {}

没有while 直接用for就可以替代while功能  for后直接写循环条件 不用加括号
for low <= high    ////   for i := 1; i <= len; ++i {}

快速声明变量并赋值  low, high  := 0, len(数组名)-1

fmt.Println(数组名) 即可输出全部的数组元素

```

```

