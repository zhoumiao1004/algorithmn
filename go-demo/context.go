package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // 同步等待

func step1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "name", "miozhou")
	return child
}

func step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 18)
	return child
}

func step3(ctx context.Context) {
	fmt.Println(ctx.Value("name"))
	fmt.Println(ctx.Value("age"))
}

func f1() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*100)
	defer cancel()
	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(err)
	}
}

func f2() {
	parent, cancel := context.WithTimeout(context.TODO(), time.Millisecond*1000)
	defer cancel()
	t0 := time.Now()

	time.Sleep(time.Millisecond * 500)

	child, cancel2 := context.WithTimeout(parent, time.Millisecond*1000)
	defer cancel2()
	t1 := time.Now()

	select {
	case <-child.Done():
		err := child.Err()
		t3 := time.Now()
		fmt.Println(t3.Sub(t0).Milliseconds(), t3.Sub(t1).Milliseconds())
		fmt.Println(err)
	}
}

func f3() {
	ctx, cancel := context.WithCancel(context.TODO())
	t0 := time.Now()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cancel()
	}()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		t3 := time.Now()
		fmt.Println(t3.Sub(t0).Milliseconds())
		fmt.Println(err)
	}
}

func main() {
	grandpa := context.TODO()
	father := step1(grandpa)
	grandson := step2(father)
	step3(grandson)
	// f1()
	// f2() // 1000 500 parent先到期，导致child Done管道被关闭，解除阻塞
	f3()

	t0 := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		ip, err := GetIp(ctx)
		fmt.Println(ip, err)
		// wg.Done()
	}()
	// 起一个协程2s后主动把上面的子协程停掉
	go func() {
		time.Sleep(2 * time.Second)
		// 取消协程
		cancel()
	}()
	wg.Wait()

	fmt.Println(time.Since(t0))
}

func GetIp(ctx context.Context) (ip string, err error) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("收到Done消息", ctx.Err())
			err = ctx.Err()
			wg.Done()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "192.168.200.1"
	wg.Done()
	return
}
