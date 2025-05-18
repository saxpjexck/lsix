package core

import (
	"encoding/json"
	"fmt"
	"jetbra-free/internal/util"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	pluginBaseUrl = "https://plugins.jetbrains.com"
)

var (
	binDir         = util.GetBinDir()
	certPath       = filepath.Join(binDir, "jetbra.pem")
	keyPath        = filepath.Join(binDir, "jetbra.key")
	powerPath      = filepath.Join(binDir, "power.txt")
	jaNetfilter    = filepath.Join(binDir, "static", "ja-netfilter", "ja-netfilter.jar")
	pluginJsonFile = filepath.Join(binDir, "plugins.json")
	client         = http.Client{Timeout: 60 * time.Second}
	AllPluginList  []*Plugin
)

type ListPluginResponse struct {
	Plugins        []*Plugin `json:"plugins,omitempty"`
	CorrectedQuery string    `json:"correctedQuery,omitempty"`
	Total          int       `json:"total,omitempty"`
}

type Plugin struct {
	Code         string   `json:"code,omitempty"`
	Name         string   `json:"name"`
	PricingModel string   `json:"pricingModel"`
	Icon         string   `json:"icon"`
	Id           int      `json:"id"`
	IsFree       bool     `json:"isFree"`
	Describe     string   `json:"describe"`
	Tags         []string `json:"tags"`
	LicenseKey   string   `json:"licenseKey"`
}

type PluginDetail struct {
	PurchaseInfo struct {
		BuyUrl        any    `json:"buyUrl"`
		PurchaseTerms any    `json:"purchaseTerms"`
		ProductCode   string `json:"productCode"`
		TrialPeriod   int    `json:"trialPeriod"`
		Optional      bool   `json:"optional"`
	} `json:"purchaseInfo"`
	Id int `json:"id"`
}

func PluginsInit() {
	fmt.Println("PluginsInit...")

	var skipFetch bool
	info, err := os.Stat(pluginJsonFile)
	if err == nil {
		modTime := info.ModTime()
		if time.Since(modTime) < 10*time.Minute {
			fmt.Println("Skipping remote fetch because the plugin file was updated within the last 10 minutes.")
			skipFetch = true
		}
		pluginFile, err := os.OpenFile(pluginJsonFile, os.O_RDONLY, 0644)
		if err == nil {
			defer pluginFile.Close()
			err = json.NewDecoder(pluginFile).Decode(&AllPluginList)
			if err != nil {
				panic(err)
			}
		}
	}

	if !skipFetch {
		loadAllPlugin()
		savePlugin()
	}
}

func loadAllPlugin() {
	pluginIdCodeMap := make(map[int]string, len(AllPluginList))
	for _, plugin := range AllPluginList {
		pluginIdCodeMap[plugin.Id] = plugin.Code
	}

	pluginList, err := client.Get(pluginBaseUrl + "/api/searchPlugins?max=10000&offset=0")
	if err != nil {
		panic(err)
	}
	defer pluginList.Body.Close()

	var listPluginResponse ListPluginResponse
	err = json.NewDecoder(pluginList.Body).Decode(&listPluginResponse)
	if err != nil {
		panic(err)
	}

	for i, plugin := range listPluginResponse.Plugins {
		if plugin.PricingModel == "FREE" {
			continue
		}
		if pluginIdCodeMap[plugin.Id] != "" {
			continue
		}
		fmt.Println("found new plugin ", plugin.Name, plugin.PricingModel)
		if plugin.Icon == "" || plugin.Icon == "https://plugins.jetbrains.com" {
			listPluginResponse.Plugins[i].Icon = filepath.Join("static", "icons", "Plugin_icon.svg")

		} else {
			listPluginResponse.Plugins[i].Icon = pluginBaseUrl + listPluginResponse.Plugins[i].Icon
		}
		AllPluginList = append(AllPluginList, listPluginResponse.Plugins[i])
	}

	for _, plugin := range AllPluginList {
		if plugin.Code == "" {
			plugin.Code = getCodeByPluginID(plugin.Id)
			fmt.Println("new plugin code ", plugin.Name, plugin.Code)
		}
	}
}

func getCodeByPluginID(id int) string {
	pluginDetailResp, err := client.Get(pluginBaseUrl + "/api/plugins/" + strconv.Itoa(id))
	if err != nil {
		panic(err)
	}
	defer pluginDetailResp.Body.Close()

	var pluginDetail PluginDetail
	err = json.NewDecoder(pluginDetailResp.Body).Decode(&pluginDetail)
	if err != nil {
		panic(err)
	}

	return pluginDetail.PurchaseInfo.ProductCode
}

func savePlugin() {
	f, err := os.Create(pluginJsonFile)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(f).Encode(AllPluginList)
	if err != nil {
		panic(err)
	}
}
