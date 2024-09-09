package test

//func TestLoadTest(t *testing.T) {
//	// 要测试的目标 URL
//	target := "http://your-target-url"
//	proto := "HTTP"
//
//	// 每秒请求数
//	callsPerSecond := 100
//	// 持续时间
//	duration := 10 * time.Second
//
//	report, err := ghz.Run(
//		context.Background(),
//		target,
//		ghz.Proto(proto),
//		ghz.QPS(callsPerSecond),
//		ghz.Duration(duration),
//	)
//	if err != nil {
//		t.Errorf("Error running load test: %v", err)
//		return
//	}
//
//	// 输出详细的测试报告
//	fmt.Printf("Average Latency: %v\n", report.Latency.Average)
//	fmt.Printf("Minimum Latency: %v\n", report.Latency.Min)
//	fmt.Printf("Maximum Latency: %v\n", report.Latency.Max)
//	fmt.Printf("Total Requests: %d\n", report.Count)
//	fmt.Printf("Error Count: %d\n", report.ErrorCount)
//	fmt.Printf("Error Percentage: %.2f%%\n", report.ErrorPercentage)
//}
