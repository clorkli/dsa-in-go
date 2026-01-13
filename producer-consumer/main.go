package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID int
	Data string
}

func Producer(out chan<- Job) {
	fmt.Println("[生产者]开始工作...")
	
	for i := 1; i <= 5; i++ {
		//模拟生产耗时
		time.Sleep(time.Millisecond * 200)

		job := Job{
			ID : i,
			Data : fmt.Sprintf("Block-Data-%d", rand.Intn(1000)),
		}

		fmt.Printf(" ->生产任务 #%d: %s\n", job.ID, job.Data)
		out <- job
	}

	close(out)
	fmt.Println("[生产者]任务发送完毕")
}

func Consumer(id int, in <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range in {
		fmt.Printf("[工人 %d] 接到任务 #%d, 正在写入磁盘...\n", id, job.ID)

		//模拟处理耗时
		time.Sleep(time.Millisecond * 500)

		fmt.Printf("[工人 %d]任务 #%d 完成\n", id, job.ID)
	}

	fmt.Printf("[工人 %d]没有任务，闲置中...\n", id)
}

func main() {
	//创建带缓冲区
	jobs := make(chan Job, 3)

	var wg sync.WaitGroup

	workCount := 2
	for i := 1; i <= workCount; i++ {
		wg.Add(1)
		go Consumer(i, jobs, &wg)
	}

	go Producer(jobs)

	fmt.Println("主线程等待所有任务完成...")
	wg.Wait()
	fmt.Println("所有业务处理完毕，程序退出。")
}