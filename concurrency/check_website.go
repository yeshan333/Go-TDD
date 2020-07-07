package concurrency

// WebsiteChecker URL 器
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites : 使用 WebsiteChecker 检查 urls slice
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)   // 存储检查结果的 map
	resultChannel := make(chan result) // 创建一个 result 的 channel

	for _, url := range urls {
		go func(u string) {
			// 引入 channel 避免 race condition
			resultChannel <- result{u, wc(u)} // send statement
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel // receive expression
		results[result.string] = result.bool
	}

	return results
}
