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
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/1.png)
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/2.png)
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/3.jpg)
# linux(kali)
```
安装依赖库
libgcrypt20-dev 
libgcrypt20
libglib2.0-dev
flex 
bison
//qt5-default
//qttools5-dev
wget http://download.qt.io/archive/qt/5.12/5.12.0/qt-opensource-linux-x64-5.12.0.run
chmod +x qt-opensource-linux-x64-5.12.0.run
./qt-opensource-linux-x64-5.12.0.run
编译
mkdir build
cd build
cmake ..
make 

```
# win 
* https://raw.githubusercontent.com/transcode-open/apt-cyg/master/apt-cyg
* https://www.pconlife.com/download/otherfile/179150/aa26678ca78b915765ece4563214cf2f/
```
用cygwin编译
麻烦就是切换不同软件源安装依赖库
cygcheck make 用于检查make依赖哪些库
```
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/4.png)
https://blog.csdn.net/dreamnow1201/article/details/80200785
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/5.png)
