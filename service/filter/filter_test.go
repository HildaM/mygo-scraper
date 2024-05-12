package filter

import (
	"regexp"
	"strings"
	"testing"
)

var (
	content = `ElevenLabs 推出音乐生成模型 ElevenLabs Music 可直接通过文本提示生成完整音乐 – XiaoHu.AI学院
	ElevenLabs 推出音乐生成模型 ElevenLabs Music 可直接通过文本提示生成完整音乐 – XiaoHu.AI学院
	
	### Menu
	
	- [Home](https://xiaohu.ai/)
	- [Trending](https://xiaohu.ai/trending)
	- [Recommended](https://xiaohu.ai/home-2)
	- [Latest](https://xiaohu.ai/latest)
	
	### 分类目录
	
	- [AI 工具](https://xiaohu.ai/c/aitools)
	- [XiaoHu.AI 学院](https://xiaohu.ai/c/aischool)
	  - [AI 教程](https://xiaohu.ai/c/aischool/aitutorials)
	  - [AI 课程](https://xiaohu.ai/c/aischool/aiclasses)
	- [XiaoHu.AI日报](https://xiaohu.ai/c/ainews)
	- [开源项目案例库](https://xiaohu.ai/c/developer)
	  - [AI 论文](https://xiaohu.ai/c/developer/paper)
	  - [AI 资源](https://xiaohu.ai/c/developer/airesources)
	  - [AI 项目](https://xiaohu.ai/c/developer/airoject)
	
	[![XiaoHu.AI学院](https://xiaohu.ai/wp-content/uploads/2024/02/350_100.png)](https://xiaohu.ai/)
	
	- [Home](https://xiaohu.ai/)
	- [AI 工具](https://xiaohu.ai/c/aitools)
	- [XiaoHu.AI 学院](https://xiaohu.ai/c/aischool)
	  - [AI 教程](https://xiaohu.ai/c/aischool/aitutorials)
	  - [AI 课程](https://xiaohu.ai/c/aischool/aiclasses)
	- [XiaoHu.AI日报](https://xiaohu.ai/c/ainews)
	- [开源项目案例库](https://xiaohu.ai/c/developer)
	  - [AI 论文](https://xiaohu.ai/c/developer/paper)
	  - [AI 资源](https://xiaohu.ai/c/developer/airesources)
	  - [AI 项目](https://xiaohu.ai/c/developer/airoject)
	- [加入会员](https://xiaohu.ai/vip)
	
	No Result
	
	View All Result
	
	- [Login](#jeg_loginform)
	
	[![XiaoHu.AI学院](https://xiaohu.ai/wp-content/uploads/2024/02/350_100.png)](https://xiaohu.ai/)
	
	No Result
	
	View All Result
	
	[Home](https://xiaohu.ai)[XiaoHu.AI日报](https://xiaohu.ai/c/ainews)
	
	# ElevenLabs 推出音乐生成模型 ElevenLabs Music 可直接通过文本提示生成完整音乐
	
	by [小互](https://xiaohu.ai/p/author/xiaohu)
	
	[2024年5月10日](https://xiaohu.ai/p/7687)
	
	in [XiaoHu.AI日报](https://xiaohu.ai/c/ainews)
	
	00
	
	[0](https://xiaohu.ai/p/7687)
	
	[![ElevenLabs 推出音乐生成模型 ElevenLabs Music 可直接通过文本提示生成完整音乐](https://xiaohu.ai/wp-content/themes/jnews/assets/img/jeg-empty.png)](https://img.xiaohu.ai/2024/05/Jietu20240510-100906@2x.jpg)
	
	0
	
	SHARES
	
	212
	
	VIEWS
	
	[Share on Facebook](https://www.facebook.com/sharer.php?u=https%3A%2F%2Fxiaohu.ai%2Fp%2F7687) [Share on Twitter](https://twitter.com/intent/tweet?text=ElevenLabs%20%E6%8E%A8%E5%87%BA%E9%9F%B3%E4%B9%90%E7%94%9F%E6%88%90%E6%A8%A1%E5%9E%8B%20ElevenLabs%20Music%20%E5%8F%AF%E7%9B%B4%E6%8E%A5%E9%80%9A%E8%BF%87%E6%96%87%E6%9C%AC%E6%8F%90%E7%A4%BA%E7%94%9F%E6%88%90%E5%AE%8C%E6%95%B4%E9%9F%B3%E4%B9%90&url=https%3A%2F%2Fxiaohu.ai%2Fp%2F7687)
	
	ElevenLabs 推出其自己的音乐生成模型 ElevenLabs Music，并展示了早期预览版生成的歌曲，该模型可直接通过文本提示生成完整带歌词的音乐。
	
	- ElevenLabs Music的早期预览展示了多首使用单一文本提示生成的歌曲。`
)

var blockString = []string{
	"Hello", "World", "Email", "Password", "Login", "Logout", "Profile", "Username", "Administrator", "Server",
	"Database", "Network", "API", "Cloud", "Storage", "File", "Edit", "Delete", "Create", "View",
	"Download", "Upload", "Refresh", "Search", "Help", "Settings", "Privacy", "Terms", "Contact", "About",
	"Home", "User", "Page", "Next", "Previous", "Submit", "Reset", "Error", "Success", "Warning",
	"Info", "Notification", "Message", "Chat", "Forum", "Comment", "Post", "Thread", "Category", "Tag",
	"公告", "登录", "注册", "用户名", "密码", "邮箱", "手机号", "验证码", "提交", "取消",
	"@", "#", "$", "%", "^", "&", "*", "(", ")", "-",
	"+", "=", "[", "]", "{", "}", "|", "\\", ":", ";",
	"\"", "'", "<", ">", ",", ".", "/", "?", "~", "`",
}

// ContentFilter2方法：使用正则表达式一次性替换所有违禁词
func (f *Filter) ContentFilter2(content string) string {
	content = strings.ToLower(content)
	var pattern strings.Builder
	pattern.WriteString("(")
	for i, b := range f.blockString {
		if i > 0 {
			pattern.WriteString("|")
		}
		pattern.WriteString(regexp.QuoteMeta(strings.ToLower(b)))
	}
	pattern.WriteString(")")
	re := regexp.MustCompile(pattern.String())
	content = re.ReplaceAllString(content, "")
	return content
}

// BenchmarkContentFilter测试ContentFilter的性能
func BenchmarkContentFilter(b *testing.B) {
	// 创建一个Filter实例，包含大量违禁词
	filter := &Filter{
		blockString: blockString,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = filter.ContentFilter(content)
	}
}

// BenchmarkContentFilter2测试ContentFilter2的性能
func BenchmarkContentFilter2(b *testing.B) {
	// 创建一个Filter实例，包含大量违禁词
	filter := &Filter{
		blockString: blockString,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = filter.ContentFilter2(content)
	}
}
