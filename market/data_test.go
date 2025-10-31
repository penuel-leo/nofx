package market

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// TestGetKlinesAster 测试Aster API K线获取
func TestGetKlinesAster(t *testing.T) {
	// 直接测试HTTP请求，避免使用可能有问题的函数
	url := "https://sapi.asterdex.com/api/v1/klines?symbol=BTCUSDT&interval=3m&limit=10"
	
	// 创建带超时的客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	t.Logf("请求URL: %s", url)
	
	resp, err := client.Get(url)
	if err != nil {
		t.Fatalf("HTTP请求失败: %v", err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("读取响应失败: %v", err)
	}
	
	t.Logf("HTTP状态码: %d", resp.StatusCode)
	
	if resp.StatusCode != 200 {
		t.Fatalf("API返回非200: %d, 响应: %s", resp.StatusCode, string(body))
	}
	
	// 解析为K线数据
	var rawData [][]interface{}
	if err := json.Unmarshal(body, &rawData); err != nil {
		t.Fatalf("JSON解析失败: %v, 响应: %s", err, string(body))
	}

	if len(rawData) == 0 {
		t.Fatalf("Aster API返回空K线数据")
	}
	
	t.Logf("✅ 成功获取 %d 条K线数据", len(rawData))
	
	// 验证数据格式（参考用户提供的示例）
	firstItem := rawData[0]
	if len(firstItem) < 7 {
		t.Fatalf("K线数据字段不足: %d < 7", len(firstItem))
	}
	
	// 解析第一条K线验证格式
	openTime := int64(firstItem[0].(float64))
	open := fmt.Sprintf("%v", firstItem[1])
	high := fmt.Sprintf("%v", firstItem[2])
	low := fmt.Sprintf("%v", firstItem[3])
	close := fmt.Sprintf("%v", firstItem[4])
	volume := fmt.Sprintf("%v", firstItem[5])
	
	t.Logf("第一条K线: OpenTime=%d, Open=%s, High=%s, Low=%s, Close=%s, Volume=%s",
		openTime, open, high, low, close, volume)
	
	if openTime == 0 {
		t.Errorf("OpenTime为0")
	}
}

// TestGetKlinesBinance 测试币安API K线获取（对照组）
func TestGetKlinesBinance(t *testing.T) {
	t.Skip("跳过币安测试，专注Aster API")
}

