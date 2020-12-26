package main

/**
 * @Author: SimonWang00
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/12/26 14:56
 */

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
)


type Watch struct {
	watch *fsnotify.Watcher
}


func (w *Watch) watchDir(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = w.watch.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	log.Println("监控服务已经启动")
	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					// 新建文件夹
					if ev.Op&fsnotify.Create == fsnotify.Create {
						log.Println("创建文件 : ", ev.Name)
						fi, err := os.Stat(ev.Name)
						// 添加文件夹了，就添加监控
						if err == nil && fi.IsDir() {
							w.watch.Add(ev.Name)
							log.Println("添加监控 : ", ev.Name)
						}
					}
					// 写文件夹
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("写入文件 : ", ev.Name)
					}
					// 删文件
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("删除文件 : ", ev.Name)
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Remove(ev.Name)
							log.Println("删除监控 : ", ev.Name)
						}
					}
					// 文件重命名
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.Println("重命名文件 : ", ev.Name)
						w.watch.Remove(ev.Name)
					}
				}
			case err := <-w.watch.Errors:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}()
}

func main() {
	watch, _ := fsnotify.NewWatcher()
	w := Watch{
		watch: watch,
	}
	w.watchDir("D:\\Go\\src\\my-gotools\\10.工具库-文件夹监控\\watch")
	select {}
}
