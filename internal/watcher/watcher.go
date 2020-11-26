package watcher

import (
	"github.com/karrick/godirwalk"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Watch(ch chan bool, dirs, exts []string) {
	if len(dirs) == 0 {
		go folder(ch, "", exts)
	} else {
		for _, dir := range dirs {
			go folder(ch, dir, exts)
		}
	}
}

func folder (ch chan bool, dir string, exts []string) {
	root := filepath.Join(".", dir)
	for {
		oldRootSize := size(root, exts)

		for {
			t := size(root, exts)
			if t != oldRootSize {
				ch <- true
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func size(root string, exts []string) (result int64) {
	err := godirwalk.Walk(root, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.IsDir() || (len(exts) != 0 && !inArray(exts, filepath.Ext(osPathname))) {
				return nil
			}
			stat, err := os.Stat(osPathname)
			if err != nil {
				log.Fatal(err)
			}
			result += stat.Size()
			return nil
		},
		Unsorted: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	return
}

func inArray(exts []string, curr string) bool {
	for _, ext := range exts {
		if curr == "." + ext {
			return true
		}
	}
	return false
}