# wireshark编译与调试
*  下载wireshark源码 

 输入命令：
```
安装 相关依赖
https://www.wireshark.org/lists/wireshark-dev/201603/msg00057.html
mkdir build
cd build 
cmake ..
make 
cmake .. -G "Xcode" // 如果报错就先 选择xcode sudo xcode-select --switch /Applications/Xcode.app/
```

https://raw.githubusercontent.com/transcode-open/apt-cyg/master/apt-cyg
cygwin 一些命令
cygcheck make
安装工具
