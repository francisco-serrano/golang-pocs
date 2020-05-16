package proc

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Usage struct {
	CpuUsage float32 `json:"cpu_usage"`
	TotalMem int     `json:"total_mem"`
	FreeMem  int     `json:"free_mem"`
}

func GetUsage() (*Usage, error) {
	cpuUsage, err := getCpuUsage()
	if err != nil {
		return nil, err
	}

	totalMem, freeMem, err := getMemUsage()
	if err != nil {
		return nil, err
	}

	return &Usage{
		CpuUsage: cpuUsage,
		TotalMem: totalMem,
		FreeMem:  freeMem,
	}, nil
}

func getCpuUsage() (float32, error) {
	var idle, total uint64

	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val
				if i == 4 {
					idle = val
				}
			}

			return float32(total-idle) / float32(total), nil
		}
	}

	return float32(total-idle) / float32(total), nil
}

func getMemUsage() (int, int, error) {
	contents, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, 0, err
	}

	lines := strings.Split(string(contents), "\n")

	totalMem := lines[0]
	freeMem := lines[1]

	regex, err := regexp.Compile("([0-9]+)")
	if err != nil {
		log.Fatal(err)
	}

	total, err := strconv.Atoi(regex.FindStringSubmatch(totalMem)[0])
	if err != nil {
		return 0, 0, err
	}

	free, err := strconv.Atoi(regex.FindStringSubmatch(freeMem)[0])
	if err != nil {
		return 0, 0, err
	}

	return total, free, nil
}

func Run() {
	usage, err := GetUsage()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(usage)
}
