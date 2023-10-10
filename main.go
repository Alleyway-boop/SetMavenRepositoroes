package main

import (
	"embed"
	"fmt"
	"os"
)

//go:embed settings.xml
var dir embed.FS

func main() {
	go CheckMavenFolderExit()
	MavenOnMacPath := ""
	args := os.Args
	MavenOnMacPath = "/Applications/IntelliJ IDEA.app/Contents/plugins/maven/lib/maven3/conf/"
	MavenOnMacPath = "./bin/"
	if len(args) >= 2 {
		MavenOnMacPath = args[1]
	}
	DirEntry, err := os.ReadDir(MavenOnMacPath)
	if err != nil {
		panic(err)
	}
	file, err := dir.ReadFile("settings.xml")
	if err != nil {
		panic(err)
	}
	for _, v := range DirEntry {
		TargetFileName := "settings.xml"
		if v.Name() == TargetFileName {
			err := os.WriteFile(MavenOnMacPath+v.Name(), file, 0666)
			AlertMessage(MavenOnMacPath, v)
			if err != nil {
				panic(err)
			}
		}
	}
}

/*
CheckMavenFolderExit
检查~/.m2/文件夹是否存在，如果存在则将settings.xml文件复制到~/.m2/文件夹下 让maven生效
*/
func CheckMavenFolderExit() {
	DefaultPath := "~/.m2/"
	DefaultPathDirs, err := os.ReadDir(DefaultPath)
	if err != nil {
		fmt.Println(".m2文件夹不存在，故忽略")
		return
	}
	file, _ := dir.ReadFile("settings.xml")
	for _, v := range DefaultPathDirs {
		if v.Name() == "settings.xml" {
			err := os.WriteFile(DefaultPath+v.Name(), file, 0666)
			if err != nil {
				panic(err)
			}
		}
	}
}

func AlertMessage(MavenOnMacPath string, v os.DirEntry) {
	fmt.Println("已设置Maven仓库为阿里云镜像")
	fmt.Println("文件位于：" + MavenOnMacPath + v.Name())
	fmt.Println("如未生效：可以尝试在pom.xml中添加")
	message := `<repositories>
			<repository>  
				<id>alimaven</id>  
				<name>aliyun maven</name>  
				<url>http://maven.aliyun.com/nexus/content/groups/public/</url>  
				<releases>  
					<enabled>true</enabled>  
				</releases>  
				<snapshots>  
					<enabled>false</enabled>  
				</snapshots>  
			    </repository>  
			</repositories>`
	fmt.Println(message)
}
func init() {
	fmt.Println("Maven仓库设置工具")
	fmt.Println("作者：yuanfang")
	fmt.Println("邮箱：yuanfangwa.gmail.com")
	/*
		todo 未完成 解析命令行参数
	*/
	//path := flag.String("path", "", "Maven安装路径")
	//flag.Parse()
}
