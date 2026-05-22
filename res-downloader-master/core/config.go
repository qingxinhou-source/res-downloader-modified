package core

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type MimeInfo struct {
	Type   string `json:"Type"`
	Suffix string `json:"Suffix"`
}

type CustomTheme struct {
	Key        string `json:"key"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Family     string `json:"family"`
	Primary    string `json:"primary"`
	Secondary  string `json:"secondary"`
	Accent     string `json:"accent"`
	Warning    string `json:"warning"`
	MollyHair  string `json:"mollyHair"`
	MollyMouth string `json:"mollyMouth"`
	Image      string `json:"image"`
}

type Personalization struct {
	ShowMetrics         bool           `json:"ShowMetrics"`
	ShowFooterLinks     bool           `json:"ShowFooterLinks"`
	ShowMollyCompanion  bool           `json:"ShowMollyCompanion"`
	ShowMollyBackground bool           `json:"ShowMollyBackground"`
	VisualFeedback      bool           `json:"VisualFeedback"`
	ShowClock           bool           `json:"ShowClock"`
	CustomWidgets       []CustomWidget `json:"CustomWidgets"`
}

type CustomWidget struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Action  string  `json:"action"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Visible bool    `json:"visible"`
}

// Config struct
type Config struct {
	storage         *Storage
	Theme           string                 `json:"Theme"`
	ThemeMode       string                 `json:"ThemeMode"`
	CustomThemes    map[string]CustomTheme `json:"CustomThemes"`
	Personalization Personalization        `json:"Personalization"`
	Locale          string                 `json:"Locale"`
	Host            string                 `json:"Host"`
	Port            string                 `json:"Port"`
	Quality         int                    `json:"Quality"`
	SaveDirectory   string                 `json:"SaveDirectory"`
	FilenameLen     int                    `json:"FilenameLen"`
	FilenameTime    bool                   `json:"FilenameTime"`
	UpstreamProxy   string                 `json:"UpstreamProxy"`
	OpenProxy       bool                   `json:"OpenProxy"`
	DownloadProxy   bool                   `json:"DownloadProxy"`
	AutoProxy       bool                   `json:"AutoProxy"`
	WxAction        bool                   `json:"WxAction"`
	TaskNumber      int                    `json:"TaskNumber"`
	DownNumber      int                    `json:"DownNumber"`
	UserAgent       string                 `json:"UserAgent"`
	UseHeaders      string                 `json:"UseHeaders"`
	InsertTail      bool                   `json:"InsertTail"`
	MimeMap         map[string]MimeInfo    `json:"MimeMap"`
	Rule            string                 `json:"Rule"`
}

var (
	mimeMux sync.RWMutex
)

func initConfig() *Config {
	if globalConfig != nil {
		return globalConfig
	}

	defaultConfig := &Config{
		Theme:        "mollyGazingTheme",
		ThemeMode:    "light",
		CustomThemes: map[string]CustomTheme{},
		Personalization: Personalization{
			ShowMetrics:         true,
			ShowFooterLinks:     true,
			ShowMollyCompanion:  true,
			ShowMollyBackground: true,
			VisualFeedback:      true,
			ShowClock:           true,
			CustomWidgets:       []CustomWidget{},
		},
		Locale:        "zh",
		Host:          "127.0.0.1",
		Port:          "8899",
		Quality:       1,
		SaveDirectory: getDefaultDownloadDir(),
		FilenameLen:   0,
		FilenameTime:  true,
		UpstreamProxy: "",
		OpenProxy:     false,
		DownloadProxy: false,
		AutoProxy:     false,
		WxAction:      true,
		TaskNumber:    runtime.NumCPU() * 2,
		DownNumber:    3,
		UserAgent:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36",
		UseHeaders:    "default",
		InsertTail:    true,
		MimeMap:       getDefaultMimeMap(),
		Rule:          "*",
	}

	rawDefaults, err := json.Marshal(defaultConfig)
	if err != nil {
		return globalConfig
	}

	storage := NewStorage("config.json", rawDefaults)
	defaultConfig.storage = storage
	globalConfig = defaultConfig

	data, err := storage.Load()
	if err != nil {
		globalLogger.Esg(err, "load config failed, using defaults")
		return globalConfig
	}

	var cacheMap map[string]interface{}
	if err := json.Unmarshal(data, &cacheMap); err != nil {
		globalLogger.Esg(err, "parse cached config failed, using defaults")
		return globalConfig
	}

	var defaultMap map[string]interface{}
	defaultBytes, _ := json.Marshal(defaultConfig)
	_ = json.Unmarshal(defaultBytes, &defaultMap)

	for k, v := range cacheMap {
		if _, ok := defaultMap[k]; ok {
			defaultMap[k] = v
		}
	}

	finalBytes, err := json.Marshal(defaultMap)
	if err != nil {
		globalLogger.Esg(err, "marshal merged config failed")
		return globalConfig
	}

	if err := json.Unmarshal(finalBytes, globalConfig); err != nil {
		globalLogger.Esg(err, "unmarshal merged config to struct failed")
	}
	globalConfig.Personalization = normalizePersonalization(globalConfig.Personalization)

	return globalConfig
}

func normalizePersonalization(value Personalization) Personalization {
	if value.CustomWidgets == nil {
		value.CustomWidgets = []CustomWidget{}
	}
	if !value.ShowMetrics &&
		!value.ShowFooterLinks &&
		!value.ShowMollyCompanion &&
		!value.ShowMollyBackground &&
		!value.VisualFeedback &&
		!value.ShowClock &&
		len(value.CustomWidgets) == 0 {
		value.ShowMetrics = true
		value.ShowFooterLinks = true
		value.ShowMollyCompanion = true
		value.ShowMollyBackground = true
		value.VisualFeedback = true
		value.ShowClock = true
	}
	return value
}

func getDefaultMimeMap() map[string]MimeInfo {
	return map[string]MimeInfo{
		"image/png":                     {Type: "image", Suffix: ".png"},
		"image/webp":                    {Type: "image", Suffix: ".webp"},
		"image/jpeg":                    {Type: "image", Suffix: ".jpeg"},
		"image/jpg":                     {Type: "image", Suffix: ".jpg"},
		"image/gif":                     {Type: "image", Suffix: ".gif"},
		"image/avif":                    {Type: "image", Suffix: ".avif"},
		"image/bmp":                     {Type: "image", Suffix: ".bmp"},
		"image/tiff":                    {Type: "image", Suffix: ".tiff"},
		"image/heic":                    {Type: "image", Suffix: ".heic"},
		"image/x-icon":                  {Type: "image", Suffix: ".ico"},
		"image/svg+xml":                 {Type: "image", Suffix: ".svg"},
		"image/vnd.adobe.photoshop":     {Type: "image", Suffix: ".psd"},
		"image/jp2":                     {Type: "image", Suffix: ".jp2"},
		"image/jpeg2000":                {Type: "image", Suffix: ".jp2"},
		"image/apng":                    {Type: "image", Suffix: ".apng"},
		"audio/mpeg":                    {Type: "audio", Suffix: ".mp3"},
		"audio/mp3":                     {Type: "audio", Suffix: ".mp3"},
		"audio/wav":                     {Type: "audio", Suffix: ".wav"},
		"audio/aiff":                    {Type: "audio", Suffix: ".aiff"},
		"audio/x-aiff":                  {Type: "audio", Suffix: ".aiff"},
		"audio/aac":                     {Type: "audio", Suffix: ".aac"},
		"audio/ogg":                     {Type: "audio", Suffix: ".ogg"},
		"audio/flac":                    {Type: "audio", Suffix: ".flac"},
		"audio/midi":                    {Type: "audio", Suffix: ".mid"},
		"audio/x-midi":                  {Type: "audio", Suffix: ".mid"},
		"audio/x-ms-wma":                {Type: "audio", Suffix: ".wma"},
		"audio/opus":                    {Type: "audio", Suffix: ".opus"},
		"audio/webm":                    {Type: "audio", Suffix: ".webm"},
		"audio/mp4":                     {Type: "audio", Suffix: ".m4a"},
		"audio/amr":                     {Type: "audio", Suffix: ".amr"},
		"video/mp4":                     {Type: "video", Suffix: ".mp4"},
		"video/webm":                    {Type: "video", Suffix: ".webm"},
		"video/ogg":                     {Type: "video", Suffix: ".ogv"},
		"video/x-msvideo":               {Type: "video", Suffix: ".avi"},
		"video/mpeg":                    {Type: "video", Suffix: ".mpeg"},
		"video/quicktime":               {Type: "video", Suffix: ".mov"},
		"video/x-ms-wmv":                {Type: "video", Suffix: ".wmv"},
		"video/3gpp":                    {Type: "video", Suffix: ".3gp"},
		"video/x-matroska":              {Type: "video", Suffix: ".mkv"},
		"audio/video":                   {Type: "live", Suffix: ".flv"},
		"video/x-flv":                   {Type: "live", Suffix: ".flv"},
		"application/dash+xml":          {Type: "live", Suffix: ".mpd"},
		"application/vnd.apple.mpegurl": {Type: "m3u8", Suffix: ".m3u8"},
		"application/x-mpegurl":         {Type: "m3u8", Suffix: ".m3u8"},
		"application/x-mpeg":            {Type: "m3u8", Suffix: ".m3u8"},
		"audio/x-mpegurl":               {Type: "m3u8", Suffix: ".m3u8"},
		"application/pdf":               {Type: "pdf", Suffix: ".pdf"},
		"application/vnd.ms-powerpoint": {Type: "ppt", Suffix: ".ppt"},
		"application/vnd.openxmlformats-officedocument.presentationml.presentation": {Type: "ppt", Suffix: ".pptx"},
		"application/vnd.ms-excel": {Type: "xls", Suffix: ".xls"},
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": {Type: "xls", Suffix: ".xlsx"},
		"text/csv":           {Type: "xls", Suffix: ".csv"},
		"application/msword": {Type: "doc", Suffix: ".doc"},
		"application/rtf":    {Type: "doc", Suffix: ".rtf"},
		"text/rtf":           {Type: "doc", Suffix: ".rtf"},
		"application/vnd.oasis.opendocument.text":                                 {Type: "doc", Suffix: ".odt"},
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": {Type: "doc", Suffix: ".docx"},
		"font/woff":                {Type: "font", Suffix: ".woff"},
		"application/octet-stream": {Type: "stream", Suffix: "default"},
	}
}

func getDefaultDownloadDir() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}

	homeDir := usr.HomeDir
	var downloadDir string

	switch runtime.GOOS {
	case "windows", "darwin":
		downloadDir = filepath.Join(homeDir, "Downloads")
	case "linux":
		downloadDir = filepath.Join(homeDir, "Downloads")
		if xdgDir := os.Getenv("XDG_DOWNLOAD_DIR"); xdgDir != "" {
			downloadDir = xdgDir
		}
	}

	if stat, err := os.Stat(downloadDir); err == nil && stat.IsDir() {
		return downloadDir
	}

	return ""
}

func (c *Config) setConfig(config Config) {
	oldProxy := c.UpstreamProxy
	openProxy := c.OpenProxy
	oldRule := c.Rule
	c.Host = config.Host
	c.Port = config.Port
	c.Theme = config.Theme
	if config.ThemeMode == "dark" {
		c.ThemeMode = "dark"
	} else {
		c.ThemeMode = "light"
	}
	if config.CustomThemes == nil {
		c.CustomThemes = map[string]CustomTheme{}
	} else {
		c.CustomThemes = config.CustomThemes
	}
	c.Personalization = normalizePersonalization(config.Personalization)
	c.Locale = config.Locale
	c.Quality = config.Quality
	c.SaveDirectory = config.SaveDirectory
	c.FilenameLen = config.FilenameLen
	c.FilenameTime = config.FilenameTime
	c.UpstreamProxy = config.UpstreamProxy
	c.UserAgent = config.UserAgent
	c.OpenProxy = config.OpenProxy
	c.DownloadProxy = config.DownloadProxy
	c.AutoProxy = config.AutoProxy
	c.TaskNumber = config.TaskNumber
	c.DownNumber = config.DownNumber
	c.WxAction = config.WxAction
	c.UseHeaders = config.UseHeaders
	c.InsertTail = config.InsertTail
	c.Rule = config.Rule
	if oldProxy != c.UpstreamProxy || openProxy != c.OpenProxy {
		proxyOnce.setTransport()
	}

	if oldRule != c.Rule {
		err := ruleOnce.Load(c.Rule)
		if err != nil {
			globalLogger.Esg(err, "set rule failed")
		}
	}

	mimeMux.Lock()
	c.MimeMap = config.MimeMap
	mimeMux.Unlock()

	jsonData, err := json.Marshal(c)
	if err == nil {
		_ = globalConfig.storage.Store(jsonData)
	}
}

func (c *Config) getConfig(key string) interface{} {
	switch key {
	case "Host":
		return c.Host
	case "Port":
		return c.Port
	case "Theme":
		return c.Theme
	case "ThemeMode":
		return c.ThemeMode
	case "CustomThemes":
		return c.CustomThemes
	case "Personalization":
		return c.Personalization
	case "Locale":
		return c.Locale
	case "Quality":
		return c.Quality
	case "SaveDirectory":
		return c.SaveDirectory
	case "FilenameLen":
		return c.FilenameLen
	case "FilenameTime":
		return c.FilenameTime
	case "UpstreamProxy":
		return c.UpstreamProxy
	case "UserAgent":
		return c.UserAgent
	case "OpenProxy":
		return c.OpenProxy
	case "DownloadProxy":
		return c.DownloadProxy
	case "AutoProxy":
		return c.AutoProxy
	case "TaskNumber":
		return c.TaskNumber
	case "DownNumber":
		return c.DownNumber
	case "WxAction":
		return c.WxAction
	case "UseHeaders":
		return c.UseHeaders
	case "InsertTail":
		return c.InsertTail
	case "MimeMap":
		mimeMux.RLock()
		defer mimeMux.RUnlock()
		return c.MimeMap
	case "Rule":
		return c.Rule
	default:
		return nil
	}
}

func (c *Config) typeSuffix(mime string) (string, string) {
	mimeMux.RLock()
	defer mimeMux.RUnlock()
	mime = strings.ToLower(strings.Split(mime, ";")[0])
	if v, ok := c.MimeMap[mime]; ok {
		return v.Type, v.Suffix
	}
	return "", ""
}
