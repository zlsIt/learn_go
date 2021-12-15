package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	//case1()
	//case2()
	case3()
}

func case1() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	n := <-ch1
	fmt.Println(n)
}

// 需求
// 定义两个chan  往第一个chan里面写值， 在读取第一个chan里面的值，往第二个chan里面写 在把第二个chan的数据读出来
func case2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		//for n := range ch1 {
		//	ch2 <- n + 1
		//}
		for {
			n, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- n * n
		}
		close(ch2)
	}()
	for n := range ch2 {
		fmt.Println(n)
	}
}

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

//需求：
//计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
//随机生成数字进行计算
func case3() {
	var wg sync.WaitGroup
	jobCh := make(chan *Job, 128)
	resCh := make(chan *Result, 128)
	for i := 0; i < 64; i++ {
		go func(jobCh chan *Job, resCh chan *Result) {
			for job := range jobCh {
				var sum int
				randNum := job.RandNum
				for randNum != 0 {
					tmp := randNum % 10
					sum += tmp
					randNum /= 10
				}
				res := &Result{
					job: job,
					sum: sum,
				}
				resCh <- res
			}
		}(jobCh, resCh)
	}
	go func(resCh chan *Result) {
		for res := range resCh {
			fmt.Printf("job id:%v randnum:%v result:%d\n", res.job.Id,
				res.job.RandNum, res.sum)
			wg.Done()
		}
	}(resCh)
	var id int
	for id < 100 {
		wg.Add(1)
		id++
		job := &Job{
			Id:      id,
			RandNum: rand.Int(),
		}
		jobCh <- job
	}
	wg.Wait()
}
