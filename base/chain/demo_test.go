package chain

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	i := int(1)
	fmt.Printf("%v", &i)

	/*pool := grpool.New(10)
	fmt.Println(time.Now())
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		sonI := i
		// 从线程池中开启一个线程跑任务
		pool.Add(context.Background(), func(ctx context.Context) {
			// todo initSDK(line)
			// 线程执行完毕后 等待组 -1
			log.Println(sonI)
			defer wg.Done()
		})
	}
	// 等待所有线程执行完毕
	wg.Wait()
	fmt.Println(time.Now())*/
}
