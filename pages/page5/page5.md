# wireshark编译与调试
* mac(10.15.2) 

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
qtbase5-dev 
qtmultimedia5-dev 
libpulse-dev 
libhamlib-dev
libqt5svg5-dev
编译
mkdir build
cd build
cmake ..
make 
```
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/linuxbuild.png)
# win(64位)失败 一直卡在cygwin 少glib2 不知道要哪个glib2
* apt工具用这个 git上的文件乱码了
* https://raw.githubusercontent.com/transcode-open/apt-cyg/master/apt-cyg
* 这个地方可以下载缺少的dll 但是不知道是否带马
* https://www.pconlife.com/download/otherfile/179150/aa26678ca78b915765ece4563214cf2f/
* 其它参考：
* https://www.wireshark.org/lists/wireshark-dev/201603/msg00057.html
* https://blog.csdn.net/dreamnow1201/article/details/80200785
```
用cygwin编译
麻烦就是切换不同软件源安装依赖库
cygcheck make 用于检查make依赖哪些库
libffi-devel (中间有要安装这个 忘记哪个了 先放这里)
```
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/4.png)
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/5.png)
# win(32位）qt环境变量一定要设置成32的不然报模块计算机类型与目标计算机类型冲突
```
mkdir vs_build_32
cd vs_build_32
cmake .. -G "Visual Studio 14 2015"
```
设置的环境变量有：
```
QT5_BASE_DIR = C:\Qt\Qt5.13.0\5.13.0\msvc2017
WIRESHARK_LIB_DIR=c:\wireshark-win64-libs-3.2(名称要看脚本 参考：https://blog.csdn.net/dreamnow1201/article/details/80087174)
WIRESHARK_VERSION_EXTRA=v1.0.0.1(随便)
PLATFORM=win32
```
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/vs32.png)
![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page5/images/vs32debug.png)
# Chocolatey 方式编译
* 如果用Chocolatey 参考：
* https://blog.csdn.net/weixin_40411459/article/details/94742456
* https://www.wireshark.org/docs/wsdg_html_chunked/ChSetupWin32.html#ChSetupGit
```
//iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
这里用cmd(root）才安装成功
@powershell -NoProfile -ExecutionPolicy Bypass -Command "iex ((new-object net.webclient).DownloadString('https://chocolatey.org/install.ps1'))" && SET PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin
```
