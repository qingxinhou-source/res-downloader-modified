<div align="center">

<img src="build/appicon.png" width="120"/>

# res-downloader-modified

<h4>📖 English | <a href="./README.md">中文</a></h4>

[![License](https://img.shields.io/github/license/putyy/res-downloader)](./LICENSE)
![Version](https://img.shields.io/badge/version-3.1.3-blue)

</div>

---

## 🌟 About This Modified Version

This project is an **unofficial modified version** based on [putyy/res-downloader](https://github.com/putyy/res-downloader), developed and published in accordance with the MIT License.

### 🙏 Acknowledgments to the Original Project

This modified version is based on **res-downloader** developed by [putyy](https://github.com/putyy). The original project is an open-source resource downloader maintained by an excellent independent developer. Sincere thanks to the original author.

> **Original Repository**: [https://github.com/putyy/res-downloader](https://github.com/putyy/res-downloader)  
> **License**: MIT License  
> **Original Author**: putyy

---

### 💝 Why Does This Modified Version Exist?

> *"Pure love never fails."*

This modified version was made for a girl I secretly loved.

She liked Pop Mart's Molly, especially the "Scenery Along The Way" 20th Anniversary series. I thought, if there was a tool that showed those gentle colors and Molly's smile every time she opened it, maybe she'd be a little happier using it.

So I added to the original version:
- 🎨 Complete Molly "Scenery Along The Way" theme system (11 themes corresponding to 11 blind box figures)
- 🌗 Light/dark mode switching
- 🧪 Custom theme lab
- ✨ More delicate interaction effects

But in the end, I didn't get the chance to send it to her.

So now, I'm open-sourcing this modified version. I hope these colors and themes can make more people happy using it. And... this way, at least this intention won't go to waste 😊

*— Dedicated to an unrequited secret love. Pure love never fails.*

---

### 🎉 Aixiang Resource Downloader (Modified)

> A cross-platform resource downloader built with Go + [Wails](https://github.com/wailsapp/wails).  
> Clean UI, easy to use, and supports a wide range of resource sniffing and downloading.

> This modified build adds the Molly "Scenery Along The Way" blind-box theme system, light/dark modes, a custom theme lab, and richer UI interactions.

## ✨ Features

- 🚀 **User-Friendly**: Simple operation with an intuitive and beautiful UI
- 🖥️ **Cross-Platform**: Available on Windows / macOS / Linux
- 🌐 **Supports Multiple Resource Types**: Video / Audio / Images / m3u8 / Live streams, and more
- 📱 **Wide Platform Compatibility**: Works with WeChat Channels, Mini Programs, Douyin, Kuaishou, Xiaohongshu, KuGou Music, QQ Music, and more
- 🌍 **Proxy Capture**: Built-in proxy allows fetching resources behind network restrictions
- 🎨 **Molly Blind-Box Themes**: 11 built-in "Scenery Along The Way" character-card themes with automatic light/dark variants
- 🌗 **Global Light/Dark Mode**: Switch with the sun/moon button on the homepage; theme changes keep the current mode
- 🧪 **Custom Theme Lab**: Customize theme name, description, character image, primary/secondary/accent/warning colors, and Molly detail colors
- ✨ **Theme Interactions**: Character-card hover feedback, Molly companion widget, theme-colored burst effects, and subtle global motion

## 🎨 Themes & Customization

- Settings → Basic → **Blind-box Themes**: choose a built-in character-card theme
- Settings → Basic → **Custom Theme Lab**: create, edit, and delete personal themes
- Custom themes are persisted locally and survive app restarts
- Custom themes also generate both light and dark global color schemes
- Theme character images live in `frontend/src/assets/image/molly`

## 🛠️ Modified Build

This repository preserves historical build artifacts. Recommended build command:

```powershell
.\build-preserve.ps1
```

After building, the following files will be generated:

- `build/bin/res-downloader-modified.exe`
- `build/bin/res-downloader-modified_yyyyMMdd_HHmmss.exe`

Timestamped historical versions will not be deleted.

## 📚 Docs & Versions

- 📘 [Online Docs](https://res.putyy.com/)
- 💬 [Join Community](https://www.putyy.com/app/admin/upload/img/20250418/6801d9554dc7.webp)
- 🧩 [Latest Release](https://github.com/putyy/res-downloader/releases) ｜ [Mini Version (UI in default browser)](https://github.com/putyy/resd-mini) ｜ [Electron Old Version (supports Win7)](https://github.com/putyy/res-downloader/tree/old)
  > *If the group is full, add WeChat `AmorousWorld` with note "github"*

## 🧩 Download

- 🆕 [GitHub Download](https://github.com/putyy/res-downloader/releases)
- 🆕 [Lanzou Cloud Download (password: 9vs5)](https://wwjv.lanzoum.com/b04wgtfyb)
- ⚠️ *Win7 users please download version `2.3.0`*

## 🖼️ Preview

![Preview](docs/images/show.webp)

---

## 🚀 Usage

> Please follow these steps to use the software correctly:

1. During installation, make sure to **allow certificate installation** and **allow network access**
2. Open the app → Click **"Start Proxy"** on the top left of the home page
3. Select the resource type to capture (default: all)
4. Open the resource page externally (e.g., WeChat Channels, Mini Programs, web pages, etc.)
5. Return to the app home page to see the resource list

## ❓ FAQ

### 📺 m3u8 Video Resources

- Online preview: [m3u8play](https://m3u8play.com/)
- Video download: [m3u8-down](https://m3u8-down.gowas.cn/)

### 📡 Live Stream Resources

- Recommended: [OBS](https://obsproject.com/) for recording (tutorial available on Baidu)

### 🐢 Slow Download / Large File Fails?

- Recommended tools:
  - [Neat Download Manager](https://www.neatdownloadmanager.com/index.php/en/)
  - [Motrix](https://motrix.app/download)
- For WeChat Channels resources, click `Video Decrypt (WeChat Channels)` in the action menu after downloading

### 🧩 Software Can't Capture Resources?

- Check if the system proxy is set correctly:  
  Address: 127.0.0.1  
  Port: 8899

### 🌐 Can't Access Internet After Closing the App?

- Manually turn off system proxy settings

### 💡 More Questions

- [GitHub Issues](https://github.com/putyy/res-downloader/issues)
- [Aixiang Forum Discussion](https://s.gowas.cn/d/4089)

## 💡 Implementation & Original Intent

This tool implements network packet capture through a proxy, then filters available resources. The principle is similar to Fiddler, Charles, and browser DevTools, but with more user-friendly resource filtering, display, and processing, significantly lowering the usage threshold and making it more suitable for general users.

---

## ⚠️ Disclaimer

> This software is for learning and research purposes only. Any commercial or illegal use is prohibited.  
> The author is not responsible for any legal liabilities arising from its use!

## 📄 Open Source License

This modified version is open-sourced under the MIT License. The original project [putyy/res-downloader](https://github.com/putyy/res-downloader) is also licensed under MIT.

Please check the [LICENSE](./LICENSE) file for the full license text.
