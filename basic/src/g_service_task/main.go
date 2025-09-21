// create module
// > cd basic\src\g_service_task
// > go mod init commandservicetask

// run the program
// > go run main.go
// or
// > go run basic\src\g_service_task\main.go

// create an executable
// > go build -o my_service_manager.exe
// or
// create shortcut
// "C:\Program Files\Go\bin\go.exe" run D:\_cd\prj\github\Go\basic\src\g_service_task\main.go

/*

list shcedules:

schtasks | more



sample cmd:

schtasks /query /TN "\Microsoft\Windows\WindowsUpdate\Scheduled Start"
schtasks /delete /TN "\Microsoft\Windows\WindowsUpdate\Scheduled Start" /F

schtasks /query /TN "\Microsoft\VisualStudio\Updates\BackgroundDownload"
schtasks /delete /TN "\Microsoft\VisualStudio\Updates\BackgroundDownload" /F

schtasks /query /TN "\Microsoft\Windows\UpdateOrchestrator\Schedule Scan"
schtasks /disable /TN "\Microsoft\Windows\UpdateOrchestrator\Schedule Scan"
  access is denied.
*/

package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Type    string
	Name    string
	Enabled bool
}

// A global slice to hold the names and status of all services to be managed.
// This is a single source of truth for the entire program.
var servicesToRun = []*Service{
	{Type: "Schedule", Name: `\Microsoft\Windows\WindowsUpdate\Scheduled Start`, Enabled: true},
	{Type: "Schedule", Name: `\Microsoft\VisualStudio\Updates\BackgroundDownload`, Enabled: true},
	{Type: "Schedule", Name: `\MicrosoftEdgeUpdateTaskMachineCore{8E5FADBC-9823-4EC8-B98F-22260C5B6E21}`, Enabled: true},
	{Type: "Schedule", Name: `\MicrosoftEdgeUpdateTaskMachineUA{A87C3E27-A5DB-4384-9E36-7F41D1262268}`, Enabled: true},
	{Type: "Schedule", Name: `\Opera Air scheduled Autoupdate 1744512974`, Enabled: true},
	{Type: "Schedule", Name: `\GoogleSystem\GoogleUpdater\GoogleUpdaterTaskSystem141.0.7376.0{869CDDC3-2C52-44D1-8032-E47A34D2E142}`, Enabled: true},
	{Type: "Schedule", Name: `\Microsoft\Windows\Windows Error Reporting\QueueReporting`, Enabled: true},
	{Type: "Schedule", Name: `\Microsoft\Windows\Application Experience\Microsoft Compatibility Appraiser`, Enabled: true},
	{Type: "Schedule", Name: `\Microsoft\Windows\Data Integrity Scan\Data Integrity Check And Scan`, Enabled: true},
	{Type: "Schedule", Name: `\Microsoft\Windows\InstallService\ScanForUpdates`, Enabled: true},
	{Type: "Schedule", Name: `\Microsoft\Windows\InstallService\WakeUpAndScanForUpdates`, Enabled: true},
	{Type: "Schedule", Name: `\Microsoft\Windows\PushToInstall\Registration`, Enabled: true},
	// {Type: "Schedule", Name: `\Microsoft\Windows\UpdateOrchestrator\Schedule Scan`, Enabled: true},
	// {Type: "Schedule", Name: `\Microsoft\Windows\UpdateOrchestrator\Schedule Work`, Enabled: true},
	// {Type: "Schedule", Name: `\Microsoft\Windows\UpdateOrchestrator\USO_UxBroker`, Enabled: true},
	// {Type: "Schedule", Name: ``, Enabled: true},
	// {Type: "Schedule", Name: ``, Enabled: true},
	// {Type: "Schedule", Name: ``, Enabled: true},

	{Type: "Service", Name: "edgeupdate", Enabled: true},
	{Type: "Service", Name: "wuauserv", Enabled: true},
	{Type: "Service", Name: "sysmain", Enabled: true},
	{Type: "Service", Name: "WSearch", Enabled: true},
	{Type: "Service", Name: "UsoSvc", Enabled: true},
	{Type: "Service", Name: "postgresql-x64-17", Enabled: true}, // AA1 comment

	{Type: "Service", Name: "vmcompute", Enabled: true},  // AA1 comment
	{Type: "Service", Name: "WslService", Enabled: true}, // AA1 comment
	// {Type: "Service", Name: "DPS", Enabled: true},        // AA1 can crash? taskmgr
	{Type: "Service", Name: "AppXSVC", Enabled: true},
	{Type: "Service", Name: "InstallService", Enabled: true},
	{Type: "Service", Name: "PcaSvc", Enabled: true},
	{Type: "Service", Name: "camsvc", Enabled: true},
	// {Type: "Service", Name: "StateRepository", Enabled: true}, // AA1 can crash? taskmgr
	// {Type: "Service", Name: "", Enabled: true},

	{Type: "Task", Name: "setup.exe", Enabled: true},
	{Type: "Task", Name: "updater.exe", Enabled: true},
	{Type: "Task", Name: "background.exe", Enabled: true},
	{Type: "Task", Name: "TiWorker.exe", Enabled: true},
	{Type: "Task", Name: "microsoftedgeupdate.exe", Enabled: true},
	{Type: "Task", Name: "officeclicktorun.exe", Enabled: true},
	{Type: "Task", Name: "opera_autoupdate.exe", Enabled: true},
	{Type: "Task", Name: "mobsync.exe", Enabled: true},
	{Type: "Task", Name: "postgres.exe", Enabled: true},
	{Type: "Task", Name: "pg_ctl.exe", Enabled: true},
	{Type: "Task", Name: "mobsync.exe", Enabled: true},

	{Type: "Task", Name: "perfwatson2.exe", Enabled: true},
	{Type: "Task", Name: "BackgroundDownload.exe", Enabled: true},
	{Type: "Task", Name: "CompatTelRunner.exe", Enabled: true},
	{Type: "Task", Name: "PhoneExperienceHost.exe", Enabled: true},
	// {Type: "Task", Name: "backgroundTaskHost.exe", Enabled: true},
	{Type: "Task", Name: "TSCUpdClt.exe", Enabled: true},
	{Type: "Task", Name: "HxTsr.exe", Enabled: true},

	// {Type: "Task", Name: "RuntimeBroker.exe", Enabled: true},
	// {Type: "Task", Name: "sihost.exe", Enabled: true},
	{Type: "Task", Name: "gamingservices.exe", Enabled: true},
	{Type: "Task", Name: "gamingservicesnet.exe", Enabled: true},
	{Type: "Task", Name: "MoNotificationUx.exe", Enabled: true},
	{Type: "Task", Name: "SearchProtocolHost.exe", Enabled: true},
	{Type: "Task", Name: "vctip.exe", Enabled: true},
	{Type: "Task", Name: "usoclient.exe", Enabled: true},       // AA1 exploring since i could not delete scheduler
	{Type: "Task", Name: "MusNotification.exe", Enabled: true}, // AA1 exploring since i could not delete scheduler
	// {Type: "Task", Name: ".exe", Enabled: true},

}

// A global variable to control the pause duration between attempts.
var sleepInterval = 2300 * time.Millisecond

// Global mutex to safely update the service list.
var mutex sync.Mutex

// main()
// filePath := "D:\\_cd\\prj\\github\\PyWinOptiPC\\config\\commands.txt"

// content, err := readTextFile(filePath)
// if err != nil {
// 	log.Printf("Failed to read file '%s': %v", filePath, err)
// 	return
// } else {
// 	fmt.Printf("Successfully read file '%s'. Content:\n%s\n", filePath, content)
// }

// fmt.Println("--------------------------------------------------")
// func readTextFile(filePath string) (string, error) {
// 	fileInfo, err := os.Stat(filePath)

// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return "", fmt.Errorf("file does not exist at path: %s", filePath)
// 		}
// 		// Return other potential errors (e.g., permissions issues).
// 		return "", fmt.Errorf("error accessing file: %w", err)
// 	}

// 	if fileInfo.IsDir() {
// 		return "", fmt.Errorf("path is a directory, not a file: %s", filePath)
// 	}

// 	// read an entire file into a byte slice.
// 	fileContentBytes, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return "", fmt.Errorf("error reading file content: %w", err)
// 	}

// 	// Convert the byte slice to a string and return it with a nil error.
// 	return string(fileContentBytes), nil
// }

func stopService(wg *sync.WaitGroup, service *Service) {
	defer wg.Done()

	if strings.TrimSpace(service.Name) == "" {
		return
	}

	cmd := exec.Command("net", "stop", service.Name)
	output, err := cmd.CombinedOutput()

	if output != nil {
		if strings.Contains(string(output), "service is not started.") {
			// fmt.Printf("%s ", service.Name)
			return
		} else if strings.Contains(string(output), "was stopped successfully") {
			fmt.Printf("%s(=s=) ", service.Name)
			return
		} else if strings.Contains(string(output), "Please try again later") {
			fmt.Printf("%s(=>=) ", service.Name)
			return
		} else if strings.Contains(string(output), "service is stopping") {
			fmt.Printf("%s(=~=) ", service.Name)
			return
		} else if strings.Contains(string(output), "service name is invalid") {
			fmt.Printf("%s(!) ", service.Name)

			mutex.Lock()
			service.Enabled = false
			mutex.Unlock()
			return
		} else {
			fmt.Printf("\n--- net stop %s:\n%s\n", service.Name, output)
		}
	}
	if err != nil {
		log.Printf("\n!!! net stop %s: %v\n", service.Name, err)
	}
}

func killTask(wg *sync.WaitGroup, process *Service) {
	defer wg.Done()

	if strings.TrimSpace(process.Name) == "" {
		return
	}

	cmd := exec.Command("taskkill", "-f", "-im", process.Name)
	output, err := cmd.CombinedOutput()

	if output != nil {
		if strings.Contains(string(output), "not found") {
			// fmt.Printf("%s ", process.Name)
			return
		} else if strings.Contains(string(output), "terminated") {
			fmt.Printf("%s(=s=) ", process.Name)
			return
		} else {
			fmt.Printf("\n--- taskkill %s:\n%s\n", process.Name, output)
		}
	}
	if err != nil {
		if !strings.Contains(err.Error(), "exit status 128") {
			fmt.Printf("\n!!! taskkill %s: %v\n", process.Name, err)
		}
	}
}

func delSchedule(wg *sync.WaitGroup, schedule *Service) {
	defer wg.Done()

	if strings.TrimSpace(schedule.Name) == "" {
		return
	}

	// var cmd *exec.Cmd
	cmd := exec.Command("schtasks", "/delete", "/TN", schedule.Name, "/F")
	output, err := cmd.CombinedOutput()

	if output != nil {
		if strings.Contains(string(output), "was successfully deleted") {
			// fmt.Printf("%s ", shedule.Name)
			return
		} else if strings.Contains(string(output), "cannot find the file specified") {
			// fmt.Printf("%s ", shedule.Name)
			return
		} else {
			fmt.Printf("\n--- schtasks %s:\n%s\n", schedule.Name, output)
		}
	}
	if err != nil {
		if !strings.Contains(err.Error(), "exit status 128") {
			fmt.Printf("\n!!! schtasks %s: %v\n", schedule.Name, err)
		}
	}
}

func main() {
	fmt.Println("Waiting for processes to terminate...")
	fmt.Println()

	// This loop will run indefinitely until you stop the program with Ctrl+C.
	for {
		var wg sync.WaitGroup
		activeCount := 0

		for _, service := range servicesToRun {
			// Lock the mutex to safely check the service status.
			mutex.Lock()
			isEnabled := service.Enabled
			mutex.Unlock()

			if isEnabled {
				wg.Add(1)
				activeCount++
				switch service.Type {
				case "Service":
					go stopService(&wg, service)
				case "Task":
					go killTask(&wg, service)
				case "Schedule":
					go delSchedule(&wg, service)
				}
			}
		}

		// If all services have been disabled, exit the program gracefully.
		if activeCount == 0 {
			fmt.Println("\n\nAll services have failed. Exiting.")
			return
		}

		wg.Wait()
		fmt.Println("...")
		time.Sleep(sleepInterval)
	}
}
