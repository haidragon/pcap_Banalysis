# wireshark编译与调试 (MAC)
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
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/1.png)
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/2.png)
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/3.jpg)
