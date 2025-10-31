package market

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// TestGetKlinesAster 测试Aster Futures API K线获取
func TestGetKlinesAster(t *testing.T) {
	// 直接测试HTTP请求，避免使用可能有问题的函数
	url := "https://fapi.asterdex.com/fapi/v3/klines?symbol=BTCUSDT&interval=3m&limit=10"
	
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

// TestGetOpenInterestAster 测试Aster Open Interest API
func TestGetOpenInterestAster(t *testing.T) {
	// 设置为aster交易所
	oldExchange := GetDefaultExchange()
	SetDefaultExchange("aster")
	defer SetDefaultExchange(oldExchange)

	oiData, err := getOpenInterestData("BTCUSDT")
	if err != nil {
		t.Fatalf("获取Aster OI失败: %v", err)
	}

	if oiData.Latest <= 0 {
		t.Errorf("OI Latest为0或负数: %.4f", oiData.Latest)
	}

	t.Logf("✅ Open Interest: Latest=%.4f, Average=%.4f", oiData.Latest, oiData.Average)
}

// TestGetFundingRateAster 测试Aster Funding Rate API
func TestGetFundingRateAster(t *testing.T) {
	// 设置为aster交易所
	oldExchange := GetDefaultExchange()
	SetDefaultExchange("aster")
	defer SetDefaultExchange(oldExchange)

	rate, err := getFundingRate("BTCUSDT")
	if err != nil {
		t.Fatalf("获取Aster Funding Rate失败: %v", err)
	}

	t.Logf("✅ Funding Rate: %.8f", rate)
}

// TestMarketGetCompleteAster 测试完整的market.Get流程（包含所有API）
func TestMarketGetCompleteAster(t *testing.T) {
	// 设置为aster交易所
	oldExchange := GetDefaultExchange()
	SetDefaultExchange("aster")
	defer SetDefaultExchange(oldExchange)

	data, err := Get("BTCUSDT")
	if err != nil {
		t.Fatalf("获取市场数据失败: %v", err)
	}

	// 验证所有数据
	if data.CurrentPrice == 0 {
		t.Errorf("CurrentPrice为0")
	}
	if data.OpenInterest == nil {
		t.Errorf("OpenInterest为nil")
	} else if data.OpenInterest.Latest <= 0 {
		t.Errorf("OpenInterest.Latest为0")
	}

	t.Logf("✅ 完整市场数据获取成功")
	t.Logf("当前价格: %.2f", data.CurrentPrice)
	t.Logf("Open Interest: %.4f", data.OpenInterest.Latest)
	t.Logf("Funding Rate: %.8f", data.FundingRate)
	t.Logf("EMA20: %.2f", data.CurrentEMA20)
	t.Logf("MACD: %.4f", data.CurrentMACD)
	t.Logf("RSI7: %.2f", data.CurrentRSI7)
}

