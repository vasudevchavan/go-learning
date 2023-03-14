package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	// " Commands taken from https://www.nivas.hr/blog/2017/09/19/measuring-disk-io-performance-macos/" +
	// "Random read/write performance" +
	// "fio --randrepeat=1 --ioengine=posixaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4k --iodepth=64 --size=4G --readwrite=randrw --rwmixread=75" +
	// "Random read performance" +
	// "fio --randrepeat=1 --ioengine=posixaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4k --iodepth=64 --size=4G --readwrite=randread" +
	// "Random write performance" +
	// "fio --randrepeat=1 --ioengine=posixaio --direct=1 --gtod_reduce=1 --name=test --filename=test --bs=4k --iodepth=64 --size=4G --readwrite=randwrite" +
	// "Commands obtained from https://cloud.google.com/compute/docs/disks/benchmarking-pd-performance" +
	// "Test write throughput by performing sequential writes with multiple parallel streams" +
	// "fio --name=write_throughput --directory=$TEST_DIR --numjobs=16 	--size=10G --time_based --runtime=60s --ramp_time=2s --ioengine=libaio --direct=1 --verify=0 --bs=1M --iodepth=64 --rw=write --group_reporting=1" +
	// "Test write IOPS by performing random writes, using an I/O block size of 4 KB and an I/O depth of at least 256" +
	// "fio --name=write_iops --directory=$TEST_DIR --size=10G --time_based --runtime=60s --ramp_time=2s --ioengine=libaio --direct=1 --verify=0 --bs=4K --iodepth=256 --rw=randwrite --group_reporting=1" +
	// "Test read throughput by performing sequential reads with multiple parallel streams (16+), using an I/O block size of 1 MB and an I/O depth of at least 64" +
	// "fio --name=read_throughput --directory=$TEST_DIR --numjobs=16 --size=10G --time_based --runtime=60s --ramp_time=2s --ioengine=libaio --direct=1 --verify=0 --bs=1M --iodepth=64 --rw=read --group_reporting=1" +
	// "Test read IOPS by performing random reads, using an I/O block size of 4 KB and an I/O depth of at least 256" +
	// "fio --name=read_iops --directory=$TEST_DIR --size=10G --time_based --runtime=60s --ramp_time=2s --ioengine=libaio --direct=1 --verify=0 --bs=4K --iodepth=256 --rw=randread --group_reporting=1" +
	// "find the list of variable to be used and specific directory can be specified in Arg[6] in the command" +
	// "Random read/write performance" +
	// "1 posixaio 1 1 test test 4k 64 1G randrw 75" +
	// "Random read performance" +
	// "1 posixaio 1 1 test test 4k 64 1G randread 75" +
	// "Random write performance" +
	// "1 posixaio 1 1 test test 4k 64 1G randwrite 75" +
	// "Based on operation type the performance benchmark results will be calculated"

	prg := "fio"
	arg_len := len(os.Args[1:])

	if arg_len == 0 {

		// This portion of the code will get fio performance benchmark and will be run inside default location defined
		// Location can be defined inside Dockerfile during multistage docker build

		fmt.Println("Performing over all benchmark")
		fmt.Println("Performing Random read/write performance")
		cmd := exec.Command(prg, "--randrepeat=1", "--ioengine=posixaio", "--direct=1", "--gtod_reduce=1", "--name=test", "--filename=test", "--bs=4k", "--iodepth=64", "--size=4G", "--readwrite=randrw", "--rwmixread=75")
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(stdout))

		fmt.Println("Performing Random read performance")
		cmd2 := exec.Command(prg, "--randrepeat=1", "--ioengine=posixaio", "--direct=1", "--gtod_reduce=1", "--name=test", "--filename=test", "--bs=4k", "--iodepth=64", "--size=4G", "--readwrite=randread")
		stdout2, err2 := cmd2.Output()
		if err2 != nil {
			fmt.Println(err2.Error())
			return
		}
		fmt.Print(string(stdout2))

		fmt.Println("Performing Random write performance")
		cmd3 := exec.Command(prg, "--randrepeat=1 ", "--ioengine=posixaio ", "--direct=1 ", "--gtod_reduce=1 ", "--name=test ", "--filename=test ", "--bs=4k ", "--iodepth=64 ", "--size=4G ", "--readwrite=randwrite")
		stdout3, err3 := cmd3.Output()
		if err != nil {
			fmt.Println(err3.Error())
			return
		}
		fmt.Print(string(stdout3))

	} else if arg_len == 1 {

		// This portion of the code will get fio performance benchmark from location passed as a argument'
		// *** Note Location has a arument must contain exact path of the file.

		arg := os.Args[1]
		location := "--filename=" + arg
		fmt.Println("Performing over all benchmark")
		fmt.Println("Performing Random read/write performance")
		cmd := exec.Command(prg, "--randrepeat=1", "--ioengine=posixaio", "--direct=1", "--gtod_reduce=1", "--name=test", location, "--bs=4k", "--iodepth=64", "--size=4G", "--readwrite=randrw", "--rwmixread=75")
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(stdout))

		fmt.Println("Performing Random read performance")
		cmd2 := exec.Command(prg, "--randrepeat=1", "--ioengine=posixaio", "--direct=1", "--gtod_reduce=1", "--name=test", location, "--bs=4k", "--iodepth=64", "--size=4G", "--readwrite=randread")
		stdout2, err2 := cmd2.Output()
		if err2 != nil {
			fmt.Println(err2.Error())
			return
		}
		fmt.Print(string(stdout2))

		fmt.Println("Performing Random write performance")
		cmd3 := exec.Command(prg, "--randrepeat=1 ", "--ioengine=posixaio ", "--direct=1 ", "--gtod_reduce=1 ", "--name=test ", location, "--bs=4k ", "--iodepth=64 ", "--size=4G ", "--readwrite=randwrite")
		stdout3, err3 := cmd3.Output()
		if err != nil {
			fmt.Println(err3.Error())
			return
		}
		fmt.Print(string(stdout3))

	} else {

		// This portion of the code will get fio performance benchmark from location passed as a argument
		// *** Note Location has a arument must contain exact path of the file.

		arg := os.Args[1]
		arg2 := os.Args[2]
		arg3 := os.Args[3]
		arg4 := os.Args[4]
		arg5 := os.Args[5]
		arg6 := os.Args[6]
		arg7 := os.Args[7]
		arg8 := os.Args[8]
		arg9 := os.Args[9]
		arg10 := os.Args[10]
		arg11 := os.Args[11]

		if arg10 == "randrw" {
			fmt.Println("Performing Random read/write performance")
			cmd := exec.Command(prg, "--randrepeat="+arg, "--ioengine="+arg2,
				"--direct="+arg3, "--gtod_reduce="+arg4, "--name="+arg5,
				"--filename="+arg6, "--bs="+arg7, "--iodepth="+arg8, "--size="+arg9, "--readwrite="+arg10, "--rwmixread="+arg11)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Print(string(stdout))

		} else if arg10 == "randread" {
			fmt.Println("Performing Random read performance")
			cmd := exec.Command(prg, "--randrepeat="+arg, "--ioengine="+arg2,
				"--direct="+arg3, "--gtod_reduce="+arg4, "--name="+arg5,
				"--filename="+arg6, "--bs="+arg7, "--iodepth="+arg8, "--size="+arg9, "--readwrite="+arg10)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Print(string(stdout))

		} else {
			fmt.Println("Random write performance")
			cmd := exec.Command(prg, "--randrepeat="+arg, "--ioengine="+arg2,
				"--direct="+arg3, "--gtod_reduce="+arg4, "--name="+arg5,
				"--filename="+arg6, "--bs="+arg7, "--iodepth="+arg8, "--size="+arg9, "--readwrite="+arg10)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Print(string(stdout))
		}
	}

}
