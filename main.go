package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aiteung/atmessage/iteung"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	// Getting filesystem statistics
	disk := DiskUsage("/")
	all := fmt.Sprintf("All: %.2f GB\n", float64(disk.All)/float64(GB))
	used := fmt.Sprintf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
	free := fmt.Sprintf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))

	msg := "*C Root Stats " + hostname + "*\n"
	msg = msg + "_Disk Space Status_\n"
	msg = msg + all + used + free

	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	memtot := fmt.Sprintf("All: %.2f GB\n", float64(memory.Total)/float64(GB))
	memused := fmt.Sprintf("Used: %.2f GB\n", float64(memory.Used)/float64(GB))
	memcached := fmt.Sprintf("Cached: %.2f GB\n", float64(memory.Cached)/float64(GB))
	memfree := fmt.Sprintf("Free: %.2f GB\n", float64(memory.Free)/float64(GB))
	msg += "_Memory Status_\n" + memtot + memcached + memused + memfree

	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	time.Sleep(time.Duration(3) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total := float64(after.Total - before.Total)
	cpuuser := fmt.Sprintf("User: %f %%\n", float64(after.User-before.User)/total*100)
	cpusys := fmt.Sprintf("System: %f %%\n", float64(after.System-before.System)/total*100)
	cpuidle := fmt.Sprintf("Idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
	//cpunice := fmt.Sprintf("cpu Nice: %f %%\n", float64(after.Nice-before.Nice)/total*100)
	//cputot := fmt.Sprintf("cpu Total: %f %%\n", float64(after.Total-before.Total)/total*100)

	msg += "_CPU Status_\n" + cpusys + cpuidle + cpuuser

	fmt.Println(msg)
	r, e := iteung.PostNotif(msg, Idgroupdebug, UrlnotifWA)
	fmt.Println(r, e)
}
