package cmd

import (
	"log"
	"github.com/go-programming-tour-book/tour/internal/word"
	"strings"
)


const (
	ModeUpper                      = iota + 1 // 转换为大写
	ModeLower                                 // 转换为小写
	ModeUnderscoreToUpperCamelCase            // 下划线转换为大驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转换为小驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转换为下划线
)


var desc = strings.Join([] string{
	"该子命令支持各种单词格式转换，模式如下："，
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 下划线单词转为大驼峰单词",
	"4: 下划线单词转为小驼峰单词",
	"5: 驼峰单词转为下划线单词"
}, "\n")

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func (cmd *conbra.Command, args []string)  {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持改转换模式，请执行help word查看帮助文档")
		}
		log.Printf("输出结果： %s", content)
	}

	func init()  {
		
	}
}