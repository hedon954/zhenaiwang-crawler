package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/**
 fetcher: 网页抓取器（公用）
 */

//确定原网站是什么编码格式的
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

//登录信息
//const cookie string = "sid=4806b106-101a-453c-a8ea-2487998b8f2a; ec=WPCTkzcB-1621520259689-e1bdb29fd3fd4-1383851470; notificationPreAuthorizeSwitch=41034; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621520315; __channelId=905821%2C0; token=1233097342.1621559603658.3891c07896ff3ac13793d5ae73640ccb; refreshToken=1233097342.1621646003658.5ae95462ae5b2e5cf26a51bc8c058e83; recommendId=%7B%22main-flow%22%3A%22recall-v1%22%2C%22off-feat%22%3A%22v1%22%2C%22feat-config%22%3A%22v1%22%2C%22model-version%22%3A%22v11%22%7D; _pc_myzhenai_showdialog_=1; _pc_myzhenai_memberid_=%22%2C1233097342%22; _efmdata=4z0tUM7xO4rdj4LEzzDwpcpyNqH2tbRnfmzY1Xd8GSGqRdYn5bLcEdkkCKbdNtux%2FBE9MVtwFh4SgSwUPwLF8AugJksqGjhOh%2B5n7vo4mds%3D; _exid=7leMW4cU0JEoh682%2BOCtZbkf1haUgcwWOtsRsWcVhFFqZDWvtfaizT2pgzMX86ym%2BtJDiefDY%2B43RUEPakpSRg%3D%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621559661"
//const cookie string = "sid=4806b106-101a-453c-a8ea-2487998b8f2a; ec=WPCTkzcB-1621520259689-e1bdb29fd3fd4-1383851470; FSSBBIl1UgzbN7NO=5wqNxvDOtSD46rZDbDa6DRQ27eFRWMAtTG1nV2b.t5dQpM4CRLY1X72X8_3PpKjQGcAZE0MwjSRWRVRTOkioVGa; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621520315; token=1233097342.1621559603658.3891c07896ff3ac13793d5ae73640ccb; refreshToken=1233097342.1621646003658.5ae95462ae5b2e5cf26a51bc8c058e83; recommendId=%7B%22main-flow%22%3A%22recall-v1%22%2C%22off-feat%22%3A%22v1%22%2C%22feat-config%22%3A%22v1%22%2C%22model-version%22%3A%22v11%22%7D; _pc_myzhenai_showdialog_=1; _pc_myzhenai_memberid_=%22%2C1233097342%22; _exid=NuVyFIFeA5DY%2BC7Yu4es%2F7jtqggx5Q3ttyZmdmtHwalnfCv4ZH7UFkE6HBKbxGnSWmYqeET2FT6g74hFkmCijA%3D%3D; _efmdata=4z0tUM7xO4rdj4LEzzDwpcpyNqH2tbRnfmzY1Xd8GSGqRdYn5bLcEdkkCKbdNtuxgJjDUmI7k0Oo4EN%2B37Bd94gBwlDvUIX9abjlH8uDrXo%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621588274; FSSBBIl1UgzbN7NP=53twyJKcv8slqqqmyOO1EsqAKbLqWQPXxRhTwiFkiaIqqA7.wIywh9TLJ7a2DZU8MwXvEykn6pfdf_VFskNJwwnPi9NOtFqNwsa7CV2augCuBDRY9W9TOFGP.50C03FEHOS8IPzwQr3z37Qmzh9NYWQP7clW1vnmJ.AnsNkRoUcRRFGuojwCi3AJnRxzXQMMEa0Y8yhNZzXyhf518TsKKbEU.k3ideCG3NpcRXg6fRwT.cy2607Cg30gVThzRHR.64VCQymcVhYV0sMPIpKK_wy7HkZWcAt3x7814s__pxFCjDNl78aIhIzCDco_XROF5Q"
//const cookie string = "sid=4806b106-101a-453c-a8ea-2487998b8f2a; ec=WPCTkzcB-1621520259689-e1bdb29fd3fd4-1383851470; FSSBBIl1UgzbN7NO=5wqNxvDOtSD46rZDbDa6DRQ27eFRWMAtTG1nV2b.t5dQpM4CRLY1X72X8_3PpKjQGcAZE0MwjSRWRVRTOkioVGa; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621520315; token=1233097342.1621559603658.3891c07896ff3ac13793d5ae73640ccb; refreshToken=1233097342.1621646003658.5ae95462ae5b2e5cf26a51bc8c058e83; recommendId=%7B%22main-flow%22%3A%22recall-v1%22%2C%22off-feat%22%3A%22v1%22%2C%22feat-config%22%3A%22v1%22%2C%22model-version%22%3A%22v11%22%7D; _pc_myzhenai_showdialog_=1; _pc_myzhenai_memberid_=%22%2C1233097342%22; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621598781; _efmdata=4z0tUM7xO4rdj4LEzzDwpcpyNqH2tbRnfmzY1Xd8GSGqRdYn5bLcEdkkCKbdNtuxw4atfLARtaLyDU0Z0%2BsNva8FhluabXqfwXS1vmcrRmM%3D; _exid=7ESeCl3ndKdil4ChF5m9soicn8CztxWx9yVrQGUgjTGPMDejzBoHbSbnOV%2FF1Wcg0X7zbxqD69iDx0t7zBLL3w%3D%3D; FSSBBIl1UgzbN7NP=53t8PAKcvLOaqqqmyOzo3NqK3UvzN_UB1ioeucr_H7RgbouUevcODEzsG6kFu4Hdzj3e4CXwi_HwJt3D79OW_OO3lxh38ZNnuUajKjJKizHpsWJVykY2j0gcPGsnsValC69FTOvK6W5JFiguFFeycP2QkSk1gGBgPfU1dHAIUF8Kqr0Ewo1NJaCczIYUX1WFGLVtyGI1fmGwHq8.Era0HBpRCnRhF6zNrO2ebcdr.ClymxuFGHImyeMZKyB90qJcJNe4qgYWQyDK8rh4U9schpk"
const userAgent string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36"


//限速：10ms抓取一次
var rateLimiter  = time.Tick(100 * time.Millisecond)

//抓取网页内容
//return 1：返回内容
//return 2: 错误
func Fetch(url string) ([]byte, error) {

	<- rateLimiter

	fmt.Println("Now is fetching", url)
	//读取网站
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil{
		fmt.Println("Fetch error when request", err)
		return nil, err
	}
	//模拟登录
	header := request.Header
	//header.Add("cookie",cookie)
	header.Add("user-agent", userAgent)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Fetch error when login", err)
		return nil, err
	}
	//最终要关闭response
	defer resp.Body.Close()
	//转编码格式，统一转为utf-8，避免乱码
	bodyBufferReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyBufferReader)
	utf8Reader := transform.NewReader(bodyBufferReader, e.NewDecoder())
	//请求失败了就不继续往下走了
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	//读取主页内容
	return ioutil.ReadAll(utf8Reader)
}


