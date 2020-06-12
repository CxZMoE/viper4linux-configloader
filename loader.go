package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"

	"github.com/CxZMoE/cxzconfutils"
)

var (
	// Is first run
	firstRun = "false"
	// Config path of this program.
	loaderConfigPath = "./v4lconfigloader.config"

	// Audio Processor Configs saved path.
	configSaveDir = ""

	// Audio Processor Main Confit (audio.conf) Path
	mainConfPath = ""
)

const banner = `
===========================================
|         viper4linuxconfigloader         |
|                                         |
|                 welcome                 |
|                                         |
| Description:                            |
|                                         |
| this program is a loader program of the |
| project: viper4linux on github,for the  |
| convinience of switching config files.  |
|                                         |
| This program does not contain any codes |
| from the viper4 team or any other teams,|
| and will be pubulish for free with MIT  |
| licence.                                |
|                                         |
| My github: https://github.com/CxZMoE    |
| Bug reports are welcomed.               |
|                                         |
===========================================
`

// Get user input stdin and apply it to variables.
func settingInput(hint string, saveVar *string) {
	for {
		fmt.Println(hint)
		fmt.Scanf("%s", saveVar)
		if cxzconfutils.ConfIsNotExist(configSaveDir) {
			fmt.Println("the path you specified does not exist,try again.")
		} else {
			break
		}
	}
}

func init() {
	// get user home dir.
	userHome := os.Getenv("HOME") + "/"
	configPath := userHome + "/.config/viper4linux-configloader"
	_, err := os.Stat(configPath)

	// create config directory when not exists.
	if os.IsNotExist(err) {
		os.MkdirAll(configPath, 0755)
	}

	// change work directory.
	os.Chdir(configPath)

	// Check is the viper installed...
	cmd := exec.Command("which", "viper")
	isins, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	// not found viper
	if len(isins) == 0 {
		fmt.Println("[ERR] viper not found in environment.")
		os.Exit(0)
	}
	// found viper
	fmt.Println("Viper4Linux found at:", string(isins))

	// What to do if config of loader is not find 👇

	if cxzconfutils.ConfIsNotExist(loaderConfigPath) {
		// First launch and initialization.
		fmt.Println("It seems that this is the first time you launch this application\nyou need to do some settings fitst.")
		settingInput("Where you store audio configs?", &configSaveDir)
		settingInput("Where is the audio.conf?", &mainConfPath)
		cxzconfutils.WriteConfFile(loaderConfigPath, map[string]string{
			"configSaveDir": configSaveDir,
			"mainConfPath":  mainConfPath,
		})

	} else {
		// Read config from file.
		configSaveDir = cxzconfutils.GetValueFromFile(loaderConfigPath, "configSaveDir")
		mainConfPath = cxzconfutils.GetValueFromFile(loaderConfigPath, "mainConfPath")
		if mainConfPath == "" {
			fmt.Println("Config is not setted correctly,exit.")
			os.Remove(loaderConfigPath)
		}
		if configSaveDir == "" {
			// Default config storage directory.
			configSaveDir = "/usr/local/viper4linux-configloader/configs"
		}
	}

}

// Program Entry
func main() {
	
	// BANNER display
	fmt.Printf("%s\n", banner)
	fmt.Printf("Please enter number of -1 to quit, quit with 'Ctrl+C' will cause viper stop.\n\n")

	// get config informations
	infos, err := loadConfigList()
	if err != nil {
		fmt.Println("[ERR]", err.Error())
		return
	}

	// When the config folder is empty.
	if len(infos) == 0 {
		fmt.Println("It does not seems that you have any config.")
		fmt.Println("You need to add configs to config folder:", configSaveDir, "and run this program.")
		return
	}

	// User Input Loop
	for {
		for i, v := range infos {
			// Print info of configs
			fmt.Printf("%d.", i)
			v.Print()
		}

		fmt.Printf("\nLoad config of number? ")

		// Get user input
		var num int
		fmt.Scanf("%d", &num)

		// Incorrect input number
		if num > len(infos)-1 || num < -1 {
			fmt.Printf("Wrong number, please try again.")
			continue
		} else {
			// Replace config files and restart viper4linux
			runReplacement(infos[num].Path, mainConfPath, infos[num].WithIRS)
		}

	}
}

// copyConf copy config file form src to dst,and backup the previous file,
func copyConf(src, dst string) error {
	// replacement
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	data, _ := ioutil.ReadAll(srcFile)
	_, err = dstFile.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// NotExist check file existance
func NotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// runReplacement replace the audio.conf file and restart viper
func runReplacement(src, dst string, hasIRS bool) error {
	fmt.Println("\nLoad config file:", src)
	//fmt.Println("IRS:", hasIRS)

	// copy audio.conf
	err := copyConf(src, dst)
	if err != nil {
		fmt.Println(err.Error())
	}

	// For config use IRS file.
	// copy irs file.

	if hasIRS {
		//srcWithNoExt := src[:strings.Index(src, path.Ext(src))]
		convIrPath := cxzconfutils.GetValueFromFile(src, "conv_ir_path")
		convIrPath = strings.TrimSuffix(path.Dir(src)+"/"+path.Base(strings.Split(convIrPath, "/")[1]), "\"")
		//fmt.Println("copy:", convIrPath)
		if NotExist("./" + convIrPath) {
			fmt.Println("not exist")
		}
		to := path.Dir(dst) + "/" + path.Base(convIrPath)
		copyConf(convIrPath, to)
		//fmt.Println("to:", to)
	}

	// reload viper

	// stop first and restart

	cmd := exec.Command("viper", "restart")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Create a new pid for viper process.
		Setpgid: true,
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Reloaded")

	return nil
}

// SelfConfig myself config
type SelfConfig struct {
	configSaveDir string
}

// ViperConfigInfo Config Object InfoStruct
type ViperConfigInfo struct {
	Path    string
	Name    string
	WithIRS bool
	Size    int
}

// Print Print info of a config
func (v *ViperConfigInfo) Print() {
	if v.WithIRS {
		fmt.Printf("[%s] [With IRS]\n", v.Name)
	} else {
		fmt.Printf("[%s]\n", v.Name)
	}

}

// GetConfigs
func loadConfigList() ([]ViperConfigInfo, error) {

	fds, err := ioutil.ReadDir(configSaveDir)
	if err != nil {
		return nil, err
	}

	var configs []ViperConfigInfo

	for _, v := range fds {
		if v.IsDir() {
			continue
		}

		if path.Ext(v.Name()) != ".conf" {
			continue
		}
		var sigleConfig = ViperConfigInfo{
			Path: configSaveDir + "/" + v.Name(),
			Name: v.Name(),
			Size: int(v.Size()),
		}

		f, err := os.Open(sigleConfig.Path)
		if err != nil {
			continue
		}
		defer f.Close()
		if cxzconfutils.CheckFileKey(sigleConfig.Path, "conv_ir_path") {
			sigleConfig.WithIRS = true
		} else {
			sigleConfig.WithIRS = false
		}
		// add a config
		configs = append(configs, sigleConfig)
	}
	return configs, nil
}
