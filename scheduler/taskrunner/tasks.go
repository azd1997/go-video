package taskrunner

import (
	"errors"
	"github.com/azd1997/go-video/scheduler/dbops"
	"log"
	"os"
	"sync"
)

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeleteRecord(3)	// 假定每次读3条
	if err != nil {
		log.Printf("Video clear dispatcher error: %v\n", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("all tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {

	errMap := &sync.Map{}		// 用于下面goroutine存储错误信息，使这里返回

ForLoop:
	for {
		select {
		case vid := <- dc:
			// 这里进行删除视频操作，由于使用go func，有可能删一组视频还没有完全删除，就又把他们中的某几个给添加到待删除组，也就是说会作许多无用功，但这里无所谓了，反正最后都会删掉
			go func(id string) {
				// 删文件
				if err := deleteVideo(id); err != nil && err != os.ErrNotExist {
					errMap.Store(id, err)
					return
				}

				// 前面没出错或者 要删除的文件不存在的话，都应该删除这条待删除记录并通知客户端
				// 删除好做，通知呢？
				// 要想当用户想删的视频不存在能通知用户，三种策略：
				// 1. 接收用户删视频请求时就检查视频是否存在，不存在直接回复。（说实话没搞懂这个删除作调度干嘛，删除操作本身就很快，还非加个数据库操作）
				// 2. 执行删除任务时想办法把不存在信息返回回去。但这里用的sync.Map.Range没法保留这些数据，除非我再用全局变量或channel把相应的vid传出去，再response
				// 3. 不通知客户端，如果是不存在就当做是删除成功了，把这条任务删除。暂时采取这种做法

				// 删待删记录
				if err := dbops.DelVideoDeleteRecord(id); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid.(string))
		default:
			break ForLoop
		}
	}

	var err error
	errMap.Range(func(key, value interface{}) bool {		// Range遍历会在遍历到第一个错误时就退出
		if err = value.(error); err != nil {
			return false
		}
		return true
	})


	return nil
}

func deleteVideo(vid string) error {
	if err := os.Remove(VIDEO_DIR + vid); err != nil && !os.IsNotExist(err) {
		log.Printf("deleteVideo: %v\n", err)
		return err
	} else if os.IsNotExist(err) {
		log.Printf("deleteVideo: %v: %s\n", os.ErrNotExist, VIDEO_DIR + vid)
		return os.ErrNotExist
	}
	return nil
}