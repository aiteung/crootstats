package main

import (
	"fmt"
	"os"

	"github.com/aiteung/atmessage/iteung"
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

	msg := "*CrootStat*\n"
	msg = msg + "_Disk Space Status_\n" + hostname + "\n"
	msg = msg + all + used + free
	fmt.Println(msg)
	r, e := iteung.PostNotif(msg, Idgroupdebug, UrlnotifWA)
	fmt.Println(r, e)
}
