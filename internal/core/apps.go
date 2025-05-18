package core

import (
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	expiryDate := time.Now().AddDate(10, 0, 0).Format("2006-01-02")

	apps := []Plugin{
		{Name: "IntelliJ IDEA", Code: "II,PCWMP,PSI", Icon: path.Join("static", "icons", "IntelliJ_IDEA_icon.svg"), IsFree: false, Describe: "IDE for Java and Kotlin developers", Tags: []string{"Java", "Kotlin", "Spring"}},
		{Name: "PyCharm", Code: "PC,PSI,PCWMP", Icon: path.Join("static", "icons", "PyCharm_icon.svg"), IsFree: false, Describe: "IDE for Python developers and data scientists", Tags: []string{"Python", "Django", "Jupyter"}},
		{Name: "PhpStorm", Code: "PS,PCWMP,PSI", Icon: path.Join("static", "icons", "PhpStorm_icon.svg"), IsFree: false, Describe: "IDE for PHP developers", Tags: []string{"PHP", "Laravel", "Symfony"}},
		{Name: "GoLand", Code: "GO,PSI,PCWMP", Icon: path.Join("static", "icons", "GoLand_icon.svg"), IsFree: false, Describe: "IDE for Go developers", Tags: []string{"Go (Golang)", "JavaScript", "TypeScript"}},
		{Name: "Rider", Code: "RD,PDB,PSI,PCWMP", Icon: path.Join("static", "icons", "Rider_icon.svg"), IsFree: true, Describe: "IDE for .NET and game developers", Tags: []string{"C#", ".NET", "ASP.NET"}},
		{Name: "CLion", Code: "CLCL,PSI,PCWMP", Icon: path.Join("static", "icons", "CLion_icon.svg"), IsFree: true, Describe: "IDE for C and C++ developers", Tags: []string{"C", "C++", "CMake"}},
		{Name: "RustRover", Code: "RR,PRR,PSI", Icon: path.Join("static", "icons", "RustRover_icon.svg"), IsFree: true, Describe: "IDE for Rust developers", Tags: []string{"Rust", "SQL", "JavaScript"}},
		{Name: "WebStorm", Code: "WS,PCWMP,PSI", Icon: path.Join("static", "icons", "WebStorm_icon.svg"), IsFree: true, Describe: "IDE for JavaScript and TypeScript developers", Tags: []string{"JavaScript", "TypeScript", "React"}},
		{Name: "RubyMine", Code: "RM,PCWMP,PSI", Icon: path.Join("static", "icons", "RubyMine_icon.svg"), IsFree: false, Describe: "IDE for Ruby and Rails developers", Tags: []string{"Ruby on Rails (RoR)", "Hotwire", "RuboCop"}},
		{Name: "DataGrip", Code: "DB,PSI,PDB", Icon: path.Join("static", "icons", "DataGrip_icon.svg"), IsFree: false, Describe: "Tool for multiple databases", Tags: []string{"Databases", "SQL", "NoSQL"}},
		{Name: "DataSpell", Code: "DS,PSI,PDB,PCWMP", Icon: path.Join("static", "icons", "DataSpell_icon.svg"), IsFree: false, Describe: "Tool for data analysis", Tags: []string{"Databases", "SQL", "NoSQL"}},
		{Name: "dotCover", Code: "DC", Icon: path.Join("static", "icons", "dotCover_icon.svg"), IsFree: false, Describe: ".NET Unit Test Runner and Code Coverage Tool", Tags: []string{"C#", ".NET", "ASP.NET"}},
		{Name: "dotTrace", Code: "DPN,DP", Icon: path.Join("static", "icons", "dotTrace_icon.svg"), IsFree: false, Describe: ".NET Performance Profiler", Tags: []string{"C#", ".NET", "ASP.NET"}},
		{Name: "dotMemory", Code: "DM", Icon: path.Join("static", "icons", "dotMemory_icon.svg"), IsFree: false, Describe: ".NET Memory Profiler", Tags: []string{"C#", ".NET", "ASP.NET"}},
		{Name: "Aqua", Code: "AQ", Icon: path.Join("static", "icons", "Aaqua_icon.svg"), IsFree: false, Describe: "IDE for test automation", Tags: []string{"AQ"}},
		{Name: "AppCode", Code: "AC,PCWMP,PSI", Icon: path.Join("static", "icons", "AppCode_icon.svg"), IsFree: false, Describe: "AppCode is no longer available as a commercial product as of December 14, 2022."},
	}

	c.HTML(http.StatusOK, "/index.html", gin.H{
		"title":        "License Generator",
		"licenseeName": "Evaluator",
		"assigneeName": "Evaluator",
		"expiryDate":   expiryDate,
		"apps":         apps,
		"plugins":      AllPluginList,
		"jaNetfilter":  jaNetfilter,
	})
}
