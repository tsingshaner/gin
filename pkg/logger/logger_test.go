package logger

import "testing"

// 10000            155188 ns/op               0 B/op          0 allocs/op
func BenchmarkInfo(b *testing.B) {
	InitLog(LevelInfo)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Info("test message")
	}
}

// 549096804                2.460 ns/op           0 B/op          0 allocs/op
func BenchmarkInfoSkip(b *testing.B) {
	InitLog(LevelWarn)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Info("test message")
	}
}
