package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"os"
)

type ProcessInfo struct {
	Name        string
	PID         int
	MemoryKB    int 
	MemoryBytes int 
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <process_name_or_pid>")
		return
	}

	for {
		processes, err := getMemoryInfo(os.Args[1])
		if err != nil {
			fmt.Println("Error:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if len(processes) == 0 {
			fmt.Println("Process not found")
		} else {
			printGraph(processes)
		}

		time.Sleep(5 * time.Second)
	}
}

func getMemoryInfo(query string) ([]ProcessInfo, error) {
	var processes []ProcessInfo

	// If query is an integer, treat it as PID
	if pid, err := strconv.Atoi(query); err == nil {
		if process, err := getProcessInfoByPID(pid); err == nil {
			processes = append(processes, process)
		} else {
			return nil, err
		}
	} else {
		// If query is a string, treat it as process name
		out, err := exec.Command("tasklist", "/fo", "csv", "/nh").Output()
		if err != nil {
			return nil, err
		}

		scanner := bufio.NewScanner(strings.NewReader(string(out)))
		for scanner.Scan() {
			fields := strings.Split(scanner.Text(), ",")
			name := strings.Trim(fields[0], "\"")
			pid, _ := strconv.Atoi(strings.Trim(fields[1], "\""))
			memoryKB, _ := strconv.Atoi(strings.Trim(fields[4], "\"")) // Memory in KB
			memoryBytes := memoryKB * 1024

			if strings.Contains(name, query) {
				processes = append(processes, ProcessInfo{Name: name, PID: pid, MemoryKB: memoryKB, MemoryBytes: memoryBytes})
			}
		}
	}

	return processes, nil
}

func getProcessInfoByPID(pid int) (ProcessInfo, error) {
	out, err := exec.Command("tasklist", "/fi", "PID eq "+strconv.Itoa(pid), "/fo", "csv", "/nh").Output()
	if err != nil {
		return ProcessInfo{}, err
	}

	fields := strings.Split(string(out), ",")
	name := strings.Trim(fields[0], "\"")
	memoryKB, _ := strconv.Atoi(strings.Trim(fields[4], "\"")) // Memory in KB
	memoryBytes := memoryKB * 1024

	return ProcessInfo{Name: name, PID: pid, MemoryKB: memoryKB, MemoryBytes: memoryBytes}, nil
}

func printGraph(processes []ProcessInfo) {
	fmt.Println("Memory Usage Graph:")
	for _, p := range processes {
		fmt.Printf("[%s - PID: %d]: %s\n", p.Name, p.PID, strings.Repeat("*", p.MemoryKB))
	}
}
