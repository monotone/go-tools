package singleton

import (
	"fmt"
	"os"
	"syscall"
)

type Singleton struct {
	filename string
	rm       bool
	writePid bool
	f        *os.File
}

func NewSingleton(filename string) *Singleton {
	return &Singleton{
		filename: filename,
	}
}

func (t *Singleton) RemoveWhenStop(rm bool) {
	t.rm = rm
}

func (t *Singleton) WritePID(w bool) {
	t.writePid = w
}

func (t *Singleton) Start() error {
	var err error
	t.f, err = os.OpenFile(t.filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	// 给文件加锁
	err = syscall.Flock(int(t.f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", t.filename, err)
	}

	// 写入进程 ID
	if t.writePid {
		_, err = t.f.WriteString(fmt.Sprintf("%d", syscall.Getpid()))
	}

	return err
}

func (t *Singleton) Stop() {
	if t.f != nil {
		syscall.Flock(int(t.f.Fd()), syscall.LOCK_UN)
		t.f.Close()
	}

	if t.rm {
		os.Remove(t.filename)
	}
}
