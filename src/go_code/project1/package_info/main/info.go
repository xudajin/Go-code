package main


// 包的使用细节记录信息


/*
1 包名和所在的文件夹名保存一致（建议,一般为小写字母
2 当引入一个文件要使用其他包函数或变量时，需要先引入对应的包
3 package 指令 在文件第一行，然后是import 指令
4 在import包时，路径从GOPATH 的src的下一级开始，
  src文件夹(默认存放源代码)，编辑器会自动从src文件夹下引入
5 为了让其他包的文件可以访问本包的函数，则函数名的首字母必须
大写，表示该函数为公开函数
6 go支撑给包名取 别名 防止包名过长

7 在同一个 package 下 不能有相同的函数名 与是否在不同的
文件下无关 ， 当 package name 相同时，不同文件中也不能有
相同的函数及全局变量名 ，当package name 相同时，即代表在同一
范围中


*/