package main

import "os"

/**
	1. 拼接地址
	2. 下载到指定目录
http://www.doczj.com/pic/view?ih=949&rn=1&doc_id=579f331e0722192e4536f6c7&o=jpg_6_0_______&pn=1&iw=638&ix=0&sign=584d1b44f78a76839d00a4ed580a8d81&type=1&iy=0&aimw=638&app_ver=2.9.8.2&ua=bd_800_800_IncredibleS_2.9.8.2_2.3.7&bid=1&app_ua=IncredibleS&uid=&cuid=&fr=3&Bdi_bear=WIFI&from=3_10000&bduss=&pid=1&screen=800_800&sys_ver=2.3.7
http://www.doczj.com/pic/view?ih=1041&rn=1&doc_id=579f331e0722192e4536f6c7&o=jpg_6_0_______&pn=2&iw=638&ix=0&sign=b2dd5681240d728fef703156b6556ede&type=1&iy=0&aimw=638&app_ver=2.9.8.2&ua=bd_800_800_IncredibleS_2.9.8.2_2.3.7&bid=1&app_ua=IncredibleS&uid=&cuid=&fr=3&Bdi_bear=WIFI&from=3_10000&bduss=&pid=1&screen=800_800&sys_ver=2.3.7
http://www.doczj.com/pic/view?ih=1041&rn=1&doc_id=579f331e0722192e4536f6c7&o=jpg_6_0_______&pn=3&iw=638&ix=0&sign=3bb5f7a73b9ef82fdc09e2c845f084a2&type=1&iy=0&aimw=638&app_ver=2.9.8.2&ua=bd_800_800_IncredibleS_2.9.8.2_2.3.7&bid=1&app_ua=IncredibleS&uid=&cuid=&fr=3&Bdi_bear=WIFI&from=3_10000&bduss=&pid=1&screen=800_800&sys_ver=2.3.7
*/

func createDir() {
	err := os.MkdirAll("doc", 777)
	if err != nil {
		panic(err)
	}
}

var baseUrl = "http://www.doczj.com/doc/539181928-%s.html"

func download(url string) {

}