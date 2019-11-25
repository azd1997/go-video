package main

import "log"

// 视频播放和上传视频文件都是http长连接，尤其是播放视频
// 如果不限流，那随着播放量的增多，会占满服务器

// bucket-token算法
// bucket : token1, token2, ... , token n
//
// request 获取 token
// response 还回 token

// 其实就是个资源池感觉

// 并发安全

type ConnLimiter struct {
	concurrentConn int
	bucket 	chan int
}

func NewConnLimit(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,		// “连接池”数目
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.\n")
		return false
	}

	cl.bucket <- 1		// 获取token
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket	// 将获得的这一个token从chan取出，就相当于释放了token。什么时候会释放conn，就是新的conn到来的时候
	log.Printf("New connection coming: %d.\n", c)
}
