package common

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// 使用真实有效的美国ip
// https://lite.ip2location.com/united-states-of-america-ip-address-ranges
// https://cdn-lite.ip2location.com/datasets/US.json?_=1683336720620
//
//	async function getIpRange() {
//	  const results = await fetch(`https://cdn-lite.ip2location.com/datasets/US.json?_=${Date.now()}`)
//	    .then((res) => res.json())
//	    .then((res) => {
//	      const limitCount = 10000;
//	      return res.data.filter((x) => parseInt(x[2].replace(/,/g,'')) >= limitCount).map((x) => `{"${x[0]}", "${x[1]}"}, //${x[2]}`);
//	    });
//	    console.log(`results : `,results);
//	    return results.join('\n');
//	}
//
// copy(await getIpRange());
var IP_RANGE = [][]string{
	{"3.2.50.0", "3.5.31.255"},         //192,000
	{"3.5.74.0", "3.5.133.255"},        //15,360
	{"3.12.0.0", "3.23.255.255"},       //786,432
	{"3.30.0.0", "3.33.34.255"},        //205,568
	{"3.33.36.0", "3.33.255.255"},      //56,320
	{"3.40.0.0", "3.63.255.255"},       //1,572,864
	{"3.80.0.0", "3.95.255.255"},       //1,048,576
	{"3.100.0.0", "3.103.255.255"},     //262,144
	{"3.116.0.0", "3.119.255.255"},     //262,144
	{"3.128.0.0", "3.247.255.255"},     //7,864,320
	{"4.0.0.0", "4.1.179.255"},         //111,616
	{"4.1.181.0", "4.14.241.255"},      //867,584
	{"4.15.21.0", "4.16.47.255"},       //72,448
	{"4.16.55.0", "4.18.65.255"},       //133,888
	{"4.18.68.0", "4.28.135.255"},      //672,768
	{"4.28.139.0", "4.31.207.255"},     //214,272
	{"4.31.209.0", "4.33.203.255"},     //129,792
	{"4.33.234.0", "4.37.0.255"},       //202,496
	{"4.37.2.0", "4.42.31.255"},        //335,360
	{"4.42.35.0", "4.55.87.255"},       //865,536
	{"4.55.95.0", "4.59.175.255"},      //282,880
	{"4.59.179.0", "4.71.152.255"},     //779,776
	{"4.71.154.0", "4.143.255.255"},    //4,744,704
	{"4.148.0.0", "4.157.255.255"},     //655,360
	{"4.184.0.0", "4.184.55.255"},      //14,336
	{"4.198.160.0", "4.198.255.255"},   //24,576
	{"4.203.96.0", "4.203.255.255"},    //40,960
	{"4.227.0.0", "4.227.255.255"},     //65,536
	{"4.232.200.0", "4.232.255.255"},   //14,336
	{"4.236.0.0", "4.236.255.255"},     //65,536
	{"4.242.0.0", "4.242.255.255"},     //65,536
	{"4.246.0.0", "4.246.255.255"},     //65,536
	{"4.248.128.0", "4.249.255.255"},   //98,304
	{"4.255.0.0", "4.255.255.255"},     //65,536
	{"5.78.0.0", "5.78.255.255"},       //65,536
	{"5.153.23.0", "5.153.63.255"},     //10,496
	{"6.0.0.0", "8.3.111.255"},         //33,779,712
	{"8.3.128.0", "8.5.250.255"},       //162,560
	{"8.5.252.0", "8.7.243.255"},       //129,024
	{"8.7.245.0", "8.10.5.255"},        //135,424
	{"8.10.8.0", "8.14.121.255"},       //291,328
	{"8.14.123.0", "8.14.196.255"},     //18,944
	{"8.15.0.0", "8.17.204.255"},       //183,552
	{"8.17.207.0", "8.18.49.255"},      //25,344
	{"8.18.51.0", "8.18.112.255"},      //15,872
	{"8.18.197.0", "8.19.7.255"},       //17,152
	{"8.19.9.0", "8.20.99.255"},        //88,832
	{"8.20.128.0", "8.20.252.255"},     //32,000
	{"8.21.42.0", "8.21.109.255"},      //17,408
	{"8.21.112.0", "8.22.175.255"},     //81,920
	{"8.22.177.0", "8.23.138.255"},     //55,808
	{"8.23.140.0", "8.23.239.255"},     //25,600
	{"8.24.16.0", "8.24.241.255"},      //57,856
	{"8.24.245.0", "8.25.95.255"},      //27,392
	{"8.25.99.0", "8.25.248.255"},      //38,400
	{"8.25.250.0", "8.26.93.255"},      //25,600
	{"8.26.95.0", "8.26.175.255"},      //20,736
	{"8.26.181.0", "8.27.63.255"},      //35,584
	{"8.27.80.0", "8.28.3.255"},        //46,080
	{"8.28.21.0", "8.28.81.255"},       //15,616
	{"8.28.83.0", "8.28.126.255"},      //11,264
	{"8.28.128.0", "8.28.200.255"},     //18,688
	{"8.28.214.0", "8.29.104.255"},     //37,632
	{"8.29.106.0", "8.29.223.255"},     //30,208
	{"8.29.225.0", "8.30.207.255"},     //61,184
	{"8.30.235.0", "8.33.95.255"},      //161,024
	{"8.33.138.0", "8.34.68.255"},      //47,872
	{"8.34.72.0", "8.34.145.255"},      //18,944
	{"8.34.147.0", "8.34.199.255"},     //13,568
	{"8.34.224.0", "8.35.56.255"},      //22,784
	{"8.35.60.0", "8.35.148.255"},      //22,784
	{"8.35.150.0", "8.35.210.255"},     //15,616
	{"8.35.212.0", "8.36.76.255"},      //30,976
	{"8.36.78.0", "8.36.129.255"},      //13,312
	{"8.36.131.0", "8.36.215.255"},     //21,760
	{"8.36.221.0", "8.37.40.255"},      //19,456
	{"8.37.44.0", "8.38.146.255"},      //91,904
	{"8.38.173.0", "8.39.5.255"},       //22,784
	{"8.39.19.0", "8.39.124.255"},      //27,136
	{"8.39.145.0", "8.39.200.255"},     //14,336
	{"8.39.216.0", "8.40.25.255"},      //16,896
	{"8.40.32.0", "8.40.106.255"},      //19,200
	{"8.40.141.0", "8.41.4.255"},       //30,720
	{"8.41.40.0", "8.42.7.255"},        //57,344
	{"8.42.9.0", "8.42.50.255"},        //10,752
	{"8.42.56.0", "8.42.160.255"},      //26,880
	{"8.42.173.0", "8.42.244.255"},     //18,432
	{"8.42.246.0", "8.43.120.255"},     //33,536
	{"8.43.124.0", "8.43.223.255"},     //25,600
	{"8.44.7.0", "8.44.57.255"},        //13,056
	{"8.44.93.0", "8.45.95.255"},       //66,304
	{"8.45.97.0", "8.46.112.255"},      //69,632
	{"8.46.119.0", "8.47.68.255"},      //52,736
	{"8.47.70.0", "8.49.215.255"},      //168,448
	{"8.49.217.0", "8.50.11.255"},      //13,056
	{"8.50.21.0", "8.51.0.255"},        //60,416
	{"8.51.64.0", "8.127.255.255"},     //5,029,888
	{"8.192.0.0", "8.192.108.255"},     //27,904
	{"8.192.110.0", "8.207.255.255"},   //1,020,416
	{"8.221.0.0", "8.221.127.255"},     //32,768
	{"8.224.0.0", "8.238.42.255"},      //928,512
	{"8.238.53.0", "8.238.142.255"},    //23,040
	{"8.238.205.0", "8.241.255.255"},   //209,664
	{"8.244.80.0", "8.244.130.255"},    //13,056
	{"8.244.132.0", "8.244.255.255"},   //31,744
	{"8.245.64.0", "8.245.127.255"},    //16,384
	{"8.245.160.0", "8.245.255.255"},   //24,576
	{"8.246.139.0", "8.246.191.255"},   //13,568
	{"8.246.201.0", "9.9.8.255"},       //1,196,032
	{"9.9.10.0", "9.255.255.255"},      //16,184,832
	{"11.0.0.0", "11.210.23.255"},      //13,768,704
	{"11.210.25.0", "12.5.185.255"},    //3,383,552
	{"12.5.188.0", "12.19.87.255"},     //891,904
	{"12.19.96.0", "12.24.140.255"},    //339,200
	{"12.24.142.0", "12.35.147.255"},   //722,432
	{"12.35.149.0", "12.41.127.255"},   //387,840
	{"12.41.136.0", "12.46.103.255"},   //319,488
	{"12.46.106.0", "12.129.111.255"},  //5,441,024
	{"12.129.113.0", "12.139.119.255"}, //657,152
	{"12.139.121.0", "12.144.81.255"},  //317,696
	{"12.144.88.0", "12.159.147.255"},  //998,400
	{"12.159.152.0", "12.174.223.255"}, //1,001,472
	{"12.175.0.0", "12.184.30.255"},    //597,760
	{"12.184.32.0", "12.192.62.255"},   //532,224
	{"12.192.64.0", "12.196.47.255"},   //258,048
	{"12.196.63.0", "12.204.9.255"},    //510,720
	{"12.204.16.0", "12.206.179.255"},  //173,056
	{"12.206.184.0", "12.208.168.255"}, //127,232
	{"12.208.172.0", "13.32.0.255"},    //5,199,104
	{"13.32.176.0", "13.32.215.255"},   //10,240
	{"13.34.93.0", "13.34.255.255"},    //41,728
	{"13.35.73.0", "13.35.127.255"},    //14,080
	{"13.44.0.0", "13.47.255.255"},     //262,144
	{"13.52.0.0", "13.52.255.255"},     //65,536
	{"13.56.0.0", "13.66.255.255"},     //720,896
	{"13.67.128.0", "13.68.255.255"},   //98,304
	{"13.71.192.0", "13.72.191.255"},   //65,536
	{"13.73.32.0", "13.73.95.255"},     //16,384
	{"13.77.64.0", "13.77.255.255"},    //49,152
	{"13.78.128.0", "13.78.255.255"},   //32,768
	{"13.82.0.0", "13.86.255.255"},     //327,680
	{"13.87.127.0", "13.88.199.255"},   //84,224
	{"13.89.0.0", "13.92.255.255"},     //262,144
	{"13.93.128.0", "13.93.255.255"},   //32,768
	{"13.96.0.0", "13.103.255.255"},    //524,288
	{"13.104.1.0", "13.104.41.255"},    //10,496
	{"13.105.204.0", "13.105.255.255"}, //13,312
	{"13.107.55.0", "13.107.136.255"},  //20,992
	{"13.107.141.0", "13.107.197.255"}, //14,592
	{"13.107.255.0", "13.111.255.255"}, //262,400
	{"13.116.0.0", "13.120.63.255"},    //278,528
	{"13.120.128.0", "13.121.64.255"},  //49,408
	{"13.121.128.0", "13.122.63.255"},  //49,152
	{"13.122.128.0", "13.123.255.255"}, //98,304
	{"13.128.0.0", "13.199.255.255"},   //4,718,592
	{"13.216.0.0", "13.224.15.255"},    //528,384
	{"13.226.9.0", "13.226.56.255"},    //12,288
	{"13.226.176.0", "13.226.243.255"}, //17,408
	{"13.240.0.0", "13.243.255.255"},   //262,144
	{"13.248.128.0", "13.248.255.255"}, //32,768
	{"13.249.34.0", "13.249.143.255"},  //28,160
	{"13.249.176.0", "13.249.240.255"}, //16,640
	{"13.252.0.0", "13.255.255.255"},   //262,144
	{"15.0.0.0", "15.106.75.255"},      //6,966,272
	{"15.106.77.0", "15.109.211.255"},  //231,168
	{"15.109.213.0", "15.113.77.255"},  //227,584
	{"15.113.79.0", "15.114.96.255"},   //70,144
	{"15.114.98.0", "15.118.101.255"},  //263,168
	{"15.118.103.0", "15.119.207.255"}, //92,416
	{"15.119.209.0", "15.122.23.255"},  //149,248
	{"15.122.25.0", "15.124.124.255"},  //156,672
	{"15.124.128.0", "15.127.94.255"},  //188,160
	{"15.127.98.0", "15.145.19.255"},   //1,159,680
	{"15.145.24.0", "15.151.255.255"},  //452,608
	{"15.153.0.0", "15.155.255.255"},   //196,608
	{"15.158.192.0", "15.159.255.255"}, //81,920
	{"15.162.0.0", "15.163.255.255"},   //131,072
	{"15.166.0.0", "15.167.255.255"},   //131,072
	{"15.169.0.0", "15.177.23.255"},    //530,432
	{"15.177.101.0", "15.183.255.255"}, //432,896
	{"15.186.0.0", "15.187.255.255"},   //131,072
	{"15.189.0.0", "15.189.255.255"},   //65,536
	{"15.190.64.0", "15.192.255.255"},  //180,224
	{"15.193.11.0", "15.195.184.255"},  //175,616
	{"15.195.186.0", "15.205.255.255"}, //673,280
	{"15.208.0.0", "15.219.200.255"},   //772,352
	{"15.219.202.0", "15.220.43.255"},  //25,088
	{"15.221.54.0", "15.221.127.255"},  //18,944
	{"15.221.153.0", "15.221.255.255"}, //26,368
	{"15.224.0.0", "15.227.255.255"},   //262,144
	{"15.231.0.0", "15.234.255.255"},   //262,144
	{"15.238.0.0", "15.247.255.255"},   //655,360
	{"15.248.72.0", "16.0.89.255"},     //528,896
}

// 獲取真實有效的隨機IP
func GetRandomIP() string {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	// 生成隨機索引
	randomIndex := rng.Intn(len(IP_RANGE))

	// 獲取隨機 IP 地址範圍
	startIP := IP_RANGE[randomIndex][0]
	endIP := IP_RANGE[randomIndex][1]

	// 將起始 IP 地址轉換為整數形式
	startIPInt := ipToUint32(net.ParseIP(startIP))
	// 將結束 IP 地址轉換為整數形式
	endIPInt := ipToUint32(net.ParseIP(endIP))

	// 生成隨機 IP 地址
	randomIPInt := rng.Uint32()%(endIPInt-startIPInt+1) + startIPInt
	randomIP := uint32ToIP(randomIPInt)

	return randomIP
}

// 將 IP 地址轉換為 uint32
func ipToUint32(ip net.IP) uint32 {
	ip = ip.To4()
	var result uint32
	result += uint32(ip[0]) << 24
	result += uint32(ip[1]) << 16
	result += uint32(ip[2]) << 8
	result += uint32(ip[3])
	return result
}

// 將 uint32 轉換為 IP 地址
func uint32ToIP(intIP uint32) string {
	ip := fmt.Sprintf("%d.%d.%d.%d", byte(intIP>>24), byte(intIP>>16), byte(intIP>>8), byte(intIP))
	return ip
}
