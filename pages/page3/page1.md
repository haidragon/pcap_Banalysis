#  winpcap编译与使用
* 驱动编译(npf.sys)

* 安装wdk7600 地址：
* https://www.microsoft.com/en-us/download/confirmation.aspx?id=11800
* 进入C:\Users\haidragon\Desktop\winpcap-WINPCAP_4_1_0_2980\winpcap-WINPCAP_4_1_0_2980\packetNtx\driver(对应自己下载的目录)
* 这里编译的64位 执行 CompileDriver.bat，如图1所示。
* ![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page3/images/1.png)
* packet.dll编译
* 安装vs2005(高版本要修改地方非常多)，打开目录地址：
```
C:\Users\haidragon\Desktop\winpcap-WINPCAP_4_1_0_2980\winpcap-WINPCAP_4_1_0_2980\packetNtx\Dll\Project
```
* ![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page3/images/packet.png)
* wpcap.dll编译
* 进入C:\Users\haidragon\Desktop\winpcap-WINPCAP_4_1_0_2980\winpcap-WINPCAP_4_1_0_2980\wpcap\PRJ
* 这里不要用git上的源码，少很多文件用这里的：
* 下载地址：https://www.winpcap.org/devel.htm
* 修改下packet.lib查找路径
* ![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page3/images/dll.png)
然后继续编译
* ![avatar](https://github.com/haidragon/pcap_Banalysis/blob/master/pages/page3/images/wpcap.png)
* 可以参考： https://blog.csdn.net/fengsuiyunqing/article/details/98076555

