package core

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	addOpens1 = "--add-opens=java.base/jdk.internal.org.objectweb.asm=ALL-UNNAMED"
	addOpens2 = "--add-opens=java.base/jdk.internal.org.objectweb.asm.tree=ALL-UNNAMED"
	agentLine = "-javaagent:" + jaNetfilter + "=jetbrains"
)

func setUserDataPath() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "JetBrains")
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "JetBrains")
	default: // Linux
		return filepath.Join(os.Getenv("HOME"), ".config", "JetBrains")
	}
}

func getDefaultVmoptionsFilename(appNormName string, osName string) string {
	app := strings.ToLower(appNormName)
	switch osName {
	case "windows":
		return app + "64.exe.vmoptions"
	default:
		return app + ".vmoptions"
	}
}

type CrackRequest struct {
	App     string  `json:"app"`
	Status  string  `json:"status"`
	License License `json:"license"`
}

func GetCrackStatus(app string) string {

	var vmoptionsFiles []string
	jetBrainsPath := setUserDataPath()
	jetBrainsPathDirs, err := os.ReadDir(jetBrainsPath)
	if err != nil {
		log.Printf("❌ Failed to read directory: %v", err)
	}
	envKey := strings.ToUpper(appNormName(app)) + "_VM_OPTIONS"

	cracked := true
	found := false

	//Traversal of the configuration directory and collect *.vmoptions files in all version directories
	for _, dir := range jetBrainsPathDirs {
		if !dir.IsDir() {
			continue
		}

		if !strings.HasPrefix(dir.Name(), appPathPrefixName(app)) {
			continue
		}
		found = true
		appPath := filepath.Join(jetBrainsPath, dir.Name())
		foundFiles, _ := filepath.Glob(filepath.Join(appPath, "*.vmoptions"))

		//If no vmoptions files are found, they are marked as uncracked
		if len(foundFiles) == 0 {
			cracked = false
			log.Printf("🔒 %s: UnCracked (no vmoptions found) 📂 %s\n", app, appPath)
			continue
		}
		vmoptionsFiles = append(vmoptionsFiles, foundFiles...)

	}

	if found {
		envPath := os.Getenv(envKey)
		if envPath != "" {
			if os.Getenv("DEBUG") == "1" {
				log.Printf("🧩 Found env: %s=%s", envKey, envPath)
			}
			if _, err := os.Stat(envPath); err == nil {
				vmoptionsFiles = append(vmoptionsFiles, envPath)
			}
		}
	}
	// check vmoptionfiles
	for _, file := range vmoptionsFiles {
		data, err := os.ReadFile(file)
		if err != nil {
			cracked = false
			continue
		}
		content := string(data)
		lines := strings.Split(content, "\n")
		hasAddOpens1 := false
		hasAddOpens2 := false
		hasAgentLine := false
		for _, line := range lines {
			trim := strings.TrimSpace(line)
			if trim == addOpens1 {
				hasAddOpens1 = true
			}
			if trim == addOpens2 {
				hasAddOpens2 = true
			}
			if trim == agentLine {
				hasAgentLine = true
			}
		}
		if !(hasAddOpens1 && hasAddOpens2 && hasAgentLine) {
			cracked = false
			log.Printf("🔒 %s Current Status: UnCracked 📄 %s\n", app, file)
		} else {
			log.Printf("🎉 %s Current Status: Cracked 📄 %s\n", app, file)
		}
	}
	if !found {
		log.Printf("ℹ️  %s: Uninstall \n", app)
		return "Uninstall"
	}
	if cracked {
		return "Cracked"
	}
	return "UnCracked"
}

func getVmoptionsFiles(app string) []string {

	var vmoptionsFiles []string
	jetBrainsPath := setUserDataPath()
	jetBrainsPathDirs, err := os.ReadDir(jetBrainsPath)
	if err != nil {
		log.Printf("❌ Failed to read directory: %v", err)
	}
	envKey := strings.ToUpper(appNormName(app)) + "_VM_OPTIONS"
	envPath := os.Getenv(envKey)
	if envPath != "" {
		if os.Getenv("DEBUG") == "1" {
			log.Printf("🧩 Found env: %s=%s", envKey, envPath)
		}
		if _, err := os.Stat(envPath); err == nil {
			vmoptionsFiles = append(vmoptionsFiles, envPath)
		}
	}
	for _, dir := range jetBrainsPathDirs {
		if !dir.IsDir() {
			continue
		}
		if !strings.HasPrefix(dir.Name(), appPathPrefixName(app)) {
			continue
		}
		appPath := filepath.Join(jetBrainsPath, dir.Name())
		foundFiles, _ := filepath.Glob(filepath.Join(appPath, "*.vmoptions"))
		if len(foundFiles) == 0 {
			defaultFilePath := filepath.Join(appPath, getDefaultVmoptionsFilename(appNormName(app), runtime.GOOS))
			log.Printf("📂 No .vmoptions file found. 📄 Create vmoptions file path: %s", defaultFilePath)
			err := os.WriteFile(defaultFilePath, []byte(""), 0644)
			if err != nil {
				log.Printf("❌ Failed to create default vmoptions file: %v", err)
			} else {
				vmoptionsFiles = append(vmoptionsFiles, defaultFilePath)
				if os.Getenv("DEBUG") == "1" {
					log.Printf("📂 Found1 %d .vmoptions files: %v", len(foundFiles), foundFiles)
					log.Printf("📂 Found1v %d .vmoptions files: %v", len(vmoptionsFiles), vmoptionsFiles)
				}

			}
		} else {
			vmoptionsFiles = append(vmoptionsFiles, foundFiles...)
			if os.Getenv("DEBUG") == "1" {
				log.Printf("📂 Found2 %d .vmoptions files: %v", len(foundFiles), foundFiles)
				log.Printf("📂 Found2v %d .vmoptions files: %v", len(vmoptionsFiles), vmoptionsFiles)
			}
		}

	}
	if os.Getenv("DEBUG") == "1" {
		log.Printf("📂 SUM Found %d .vmoptions files: %v", len(vmoptionsFiles), vmoptionsFiles)
	}
	return vmoptionsFiles
}

func getKeyFiles(app string) []string {

	var keyFiles []string
	jetBrainsPath := setUserDataPath()
	jetBrainsPathDirs, err := os.ReadDir(jetBrainsPath)
	if err != nil {
		log.Printf("❌ Failed to read directory: %v", err)
	}

	for _, dir := range jetBrainsPathDirs {
		if !dir.IsDir() {
			continue
		}
		if !strings.HasPrefix(dir.Name(), appPathPrefixName(app)) {
			continue
		}
		appPath := filepath.Join(jetBrainsPath, dir.Name())
		// keyFile
		keyFile := filepath.Join(appPath, strings.ToLower(appNormName(app))+".key")
		keyFiles = append(keyFiles, keyFile)
	}
	if os.Getenv("DEBUG") == "1" {
		log.Printf("📂 SUM Found %d key files: %v", len(keyFiles), keyFiles)
	}
	return keyFiles
}
func appPathPrefixName(app string) string {
	appPathPrefixName := app
	if appPathPrefixName == "IntelliJ IDEA" {
		appPathPrefixName = "IntelliJIdea"
	}
	return appPathPrefixName
}
func appNormName(app string) string {
	appNormName := app
	if appNormName == "IntelliJ IDEA" {
		appNormName = "IDEA"
	}
	return appNormName
}
func CrackHandler(c *gin.Context) {
	var req CrackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("❌ Parsing request failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	app := req.App
	action := ""
	status := req.Status

	log.Printf("🚀 Start processing the request: app=%s, status=%s", app, status)

	log.Printf("🚀 GetCrackStatus Start...")
	currentStatus := GetCrackStatus(appPathPrefixName(app))
	log.Printf("🚀 Start %s %s...", app, status)

	vmoptionsFiles := getVmoptionsFiles(app)
	keyFiles := getKeyFiles(app)

	found := false
	setkey := false
	backup := false
	if status == "Cracked" {
		if currentStatus == "UnCracked" {
			for _, file := range vmoptionsFiles {
				found = true
				backupPath := file + ".jetbra-free.bak"
				data, err := os.ReadFile(file)
				if err == nil {
					log.Printf("📦 Backup vmoptions: %s", backupPath)
					_ = os.WriteFile(backupPath, data, 0644)
				}
				err = editVmoptionsFile(file)
				if err != nil {
					fmt.Printf("Failed to patch %s: %v\n", file, err)
				}
			}
			// Backup Key
			for _, file := range keyFiles {
				keyBackupPath := file + ".jetbra-free.bak"
				data, err := os.ReadFile(file)
				if err == nil {
					_ = os.WriteFile(keyBackupPath, data, 0644)
					log.Printf("🔑 Backup key: %s", keyBackupPath)
				}
			}
			backup = true
		}
		// setKeyFile
		licenseStr, _ := GenerateLicense(&req.License)
		for _, file := range keyFiles {
			if err := setKeyFile(file, licenseStr); err != nil {
				log.Printf("❌ Failed to set key file: %v", err)
			}
		}
		setkey = true

		if setkey && !backup {
			action = "CrackedWithoutBackup"
		} else {
			action = "Cracked"
		}

	} else if status == "UnCracked" {
		for _, file := range vmoptionsFiles {
			found = true
			if err := revertVmoptionsFile(file); err != nil {
				log.Printf("❌ Restore failed: %v", err)
			}
		}
		for _, file := range keyFiles {
			if err := revertKeyFile(file); err != nil {
				log.Printf("❌ Failed to restore key file: %v", err)
			}
		}
		action = "UnCracked"
	}

	if action != "" && found {
		c.JSON(200, gin.H{"msg": action})
		return
	}
	if !found {
		c.JSON(200, gin.H{"msg": "Uninstall"})
		return
	}
	c.JSON(200, gin.H{"msg": "ERROR"})
}

func setKeyFile(keyPath string, licenseStr string) error {
	content := "\uFFFF<certificate-key>\n" + licenseStr

	encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	utf16Content, _, err := transform.String(encoder, content)
	if err != nil {
		return err
	}
	dir := filepath.Dir(keyPath)
	tmpFile, err := os.CreateTemp(dir, "keyfile-*.tmp")
	if err != nil {
		return err
	}
	tmpPath := tmpFile.Name()

	// Write content
	if _, err := tmpFile.Write([]byte(utf16Content)); err != nil {
		tmpFile.Close()
		os.Remove(tmpPath)
		return err
	}
	tmpFile.Close()

	log.Printf("🔑 Replacing key file atomically: %s", keyPath)

	// 原Sub-sex replacement
	if err := os.Rename(tmpPath, keyPath); err != nil {
		os.Remove(tmpPath)
		return err
	}
	return nil
}

func revertKeyFile(keyPath string) error {

	backupPath := keyPath + ".jetbra-free.bak"
	if _, err := os.Stat(backupPath); err == nil {
		err := os.Rename(backupPath, keyPath)
		if err != nil {
			return err
		}
		log.Printf("🔑 Revert key: %s", keyPath)
		return nil
	}
	if _, err := os.Stat(keyPath); err == nil {
		log.Printf("🔑 Delete key: %s", keyPath)
		return os.Remove(keyPath)
	}
	return nil
}

func editVmoptionsFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var (
		lines        []string
		hasAddOpens1 bool
		hasAddOpens2 bool
		hasJavaAgent bool
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, addOpens1) {
			hasAddOpens1 = true
		}
		if strings.Contains(line, addOpens2) {
			hasAddOpens2 = true
		}
		if strings.HasPrefix(line, "-javaagent:") {
			hasJavaAgent = true
			lines = append(lines, agentLine)
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Close the file in advance, otherwise rename will fail under Windows
	file.Close()

	if !hasAddOpens1 {
		lines = append(lines, addOpens1)
	}
	if !hasAddOpens2 {
		lines = append(lines, addOpens2)
	}
	if !hasJavaAgent {
		lines = append(lines, agentLine)
	}

	dir := filepath.Dir(path)
	tempFile, err := os.CreateTemp(dir, "*.vmoptions.tmp")
	if err != nil {
		return err
	}
	tempPath := tempFile.Name()

	//Write all lines (original content + append content)
	if _, err := tempFile.WriteString(strings.Join(lines, "\n") + "\n"); err != nil {
		tempFile.Close()
		os.Remove(tempPath)
		return err
	}

	//Close temporary files
	if err := tempFile.Close(); err != nil {
		os.Remove(tempPath)
		return err
	}

	//Atomic replacement of the original file
	if err := os.Rename(tempPath, path); err != nil {
		os.Remove(tempPath)
		return err
	}

	log.Printf("📄 Updated vmoptions atomically: %s", path)
	return nil

}

func revertVmoptionsFile(path string) error {
	backupPath := path + ".jetbra-free.bak"

	_, err := os.Stat(backupPath)
	if err == nil {
		data, err := os.ReadFile(backupPath)
		if err != nil {
			println("Error reading backup file:", err)
			return err
		}
		if err := os.WriteFile(path, data, 0644); err != nil {
			println("Error writing to vmoptions file:", err)
			return err
		}
		log.Println("📄 Revert vmoptions:", path)
		return os.Remove(backupPath)
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trim := strings.TrimSpace(line)
		if trim == addOpens1 {
			continue
		}
		if trim == addOpens2 {
			continue
		}
		if strings.HasPrefix(trim, "-javaagent:") {
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	log.Printf("📄 Remove the hack in the vmoptions file: %s", path)
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}
