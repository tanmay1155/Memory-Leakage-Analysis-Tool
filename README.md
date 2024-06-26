# Memory Leakage Analysis Tool (MLAT) - Go Language

MLAT is a command-line tool written in Go that allows you to analyze memory leakage in processes running on your system. It identifies processes with memory leaks and provides insights into their memory usage patterns.

# Features

Memory Leakage Detection: MLAT identifies processes exhibiting memory leakage by analyzing their memory usage over time.

Process Monitoring: The tool continuously monitors the memory usage of specified processes and alerts users when memory leakage is detected.

Detailed Analysis: MLAT provides detailed reports on memory usage trends, allowing users to diagnose and troubleshoot memory leakage issues effectively.

# Getting Started

To use MLAT, follow these steps:

Build the Tool: Use the Go compiler to build the MLAT executable from the provided source code files.

Run the Tool: Execute the compiled MLAT executable from the command line, specifying the process name or PID to monitor.

Monitor Memory Usage: MLAT will start monitoring the specified process and analyzing its memory usage. It will alert you if memory leakage is detected.

Diagnose Memory Leakage: Use the detailed reports generated by MLAT to diagnose and troubleshoot memory leakage issues in the monitored process.

# Example Usage

#Build MLAT executable
go run main.go <process_name/Pid>

#Run MLAT to monitor a process by name
go run main.go memoryleak.exe

#Run MLAT to monitor a process by PID
go run main.go 12345

# Contributing

Contributions to MLAT are welcome. If you have suggestions for new features, improvements, or bug fixes, feel free to open an issue or submit a pull request.
