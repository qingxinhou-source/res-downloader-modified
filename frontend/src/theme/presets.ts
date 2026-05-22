export type ThemeMode = "light" | "dark"

export interface ThemeSeed {
  key: string
  name: string
  desc: string
  family: "molly" | "classic" | "custom"
  primary: string
  secondary: string
  accent: string
  warning: string
  mollyHair: string
  mollyMouth: string
  image: string
}

export interface ThemePreset extends ThemeSeed {
  mode: ThemeMode
  primaryHover: string
  primaryPressed: string
  error: string
  background: string
  surface: string
  surfaceContainer: string
  surfaceContainerHigh: string
  outline: string
  text: string
  textMuted: string
  shadow1: string
  shadow2: string
  glow: string
  molly: string
  mollyEye: string
  swatch: string
}

const makeShadow = (rgb: string) => ({
  shadow1: `0 12px 28px rgba(${rgb}, 0.13), 0 2px 6px rgba(32, 31, 28, 0.08)`,
  shadow2: `0 24px 58px rgba(${rgb}, 0.20), 0 8px 18px rgba(32, 31, 28, 0.12)`,
})

export const defaultThemeKey = "mollyGazingTheme"

export const themeSeeds: Record<string, ThemeSeed> = {
  mollyGazingTheme: {
    key: "mollyGazingTheme",
    name: "看着你",
    desc: "粉桃、湖绿和暖金的直视感",
    family: "molly",
    primary: "#e45d86",
    secondary: "#158f8b",
    accent: "#158f8b",
    warning: "#d99a26",
    mollyHair: "#e8af32",
    mollyMouth: "#c53d5d",
    image: "看着你.png",
  },
  mollyRealmTheme: {
    key: "mollyRealmTheme",
    name: "我的国度",
    desc: "深蓝舞台、皇冠金和冷月光",
    family: "molly",
    primary: "#7fa8ff",
    secondary: "#f2c75a",
    accent: "#f2c75a",
    warning: "#f2c75a",
    mollyHair: "#d6a43d",
    mollyMouth: "#fb91a5",
    image: "我的国度.png",
  },
  mollyBloomsTheme: {
    key: "mollyBloomsTheme",
    name: "花花世界",
    desc: "玫粉花瓣、嫩绿叶片和淡紫点缀",
    family: "molly",
    primary: "#df4f99",
    secondary: "#69a84c",
    accent: "#7c6bc8",
    warning: "#d68b20",
    mollyHair: "#e5aa38",
    mollyMouth: "#bc3c64",
    image: "花花世界.png",
  },
  mollyForestTheme: {
    key: "mollyForestTheme",
    name: "童年森林",
    desc: "苔绿、树皮棕和旧时光暖黄",
    family: "molly",
    primary: "#5d8840",
    secondary: "#b87138",
    accent: "#b87138",
    warning: "#c8902d",
    mollyHair: "#d4a23b",
    mollyMouth: "#b44a55",
    image: "童年森林.png",
  },
  mollyCastleTheme: {
    key: "mollyCastleTheme",
    name: "城堡外的星星",
    desc: "夜紫、星光金和薄荷蓝",
    family: "molly",
    primary: "#a997ff",
    secondary: "#f4c653",
    accent: "#72d7e6",
    warning: "#f4c653",
    mollyHair: "#d5a13a",
    mollyMouth: "#f28da7",
    image: "城堡外的星星.png",
  },
  mollyCottageTheme: {
    key: "mollyCottageTheme",
    name: "红顶小屋",
    desc: "砖红屋顶、亚麻墙面和庭院绿",
    family: "molly",
    primary: "#c64e3f",
    secondary: "#6b9970",
    accent: "#6b9970",
    warning: "#c7892c",
    mollyHair: "#dfa63c",
    mollyMouth: "#b8404d",
    image: "红顶小屋.png",
  },
  mollyThunderTheme: {
    key: "mollyThunderTheme",
    name: "一起听打雷",
    desc: "雨夜蓝、闪电黄和安静灰",
    family: "molly",
    primary: "#73aee8",
    secondary: "#f1c84c",
    accent: "#f1c84c",
    warning: "#f1c84c",
    mollyHair: "#d69f35",
    mollyMouth: "#f08ca0",
    image: "一起听打雷.png",
  },
  mollySleepTheme: {
    key: "mollySleepTheme",
    name: "抱你入睡",
    desc: "睡衣蓝、月光奶油和浅紫梦境",
    family: "molly",
    primary: "#688bd7",
    secondary: "#bd92d2",
    accent: "#bd92d2",
    warning: "#d5a33c",
    mollyHair: "#d7a33f",
    mollyMouth: "#bd4b68",
    image: "抱你入睡.png",
  },
  mollyBeeTheme: {
    key: "mollyBeeTheme",
    name: "我是蜂子",
    desc: "蜂蜜黄、草地绿和黑色条纹",
    family: "molly",
    primary: "#d99b12",
    secondary: "#6d8f3c",
    accent: "#2f2b24",
    warning: "#d99b12",
    mollyHair: "#d99b12",
    mollyMouth: "#b63e52",
    image: "我是蜂子.png",
  },
  mollyBrownTheme: {
    key: "mollyBrownTheme",
    name: "第一只茉莉棕色",
    desc: "复古棕、旧玩具金和奶咖底色",
    family: "molly",
    primary: "#805238",
    secondary: "#c4a05e",
    accent: "#c4a05e",
    warning: "#bd8132",
    mollyHair: "#b48742",
    mollyMouth: "#9f4047",
    image: "第一只molly棕色.png",
  },
  mollyGreenTheme: {
    key: "mollyGreenTheme",
    name: "第一只MOLLY绿色",
    desc: "初代绿、奶白和微暖金",
    family: "molly",
    primary: "#4c9466",
    secondary: "#e6b35a",
    accent: "#e6b35a",
    warning: "#dfa43a",
    mollyHair: "#d8a947",
    mollyMouth: "#c24c65",
    image: "第一只molly绿色.png",
  },
}

let customThemeSeeds: Record<string, ThemeSeed> = {}

const themeCache = new Map<string, ThemePreset>()

const hexToRgb = (hex: string) => {
  const value = hex.replace("#", "")
  if (value.length !== 6) return [120, 120, 140]
  return [
    Number.parseInt(value.slice(0, 2), 16),
    Number.parseInt(value.slice(2, 4), 16),
    Number.parseInt(value.slice(4, 6), 16),
  ]
}

const toHex = (value: number) => value.toString(16).padStart(2, "0")

const mixHex = (from: string, to: string, amount: number) => {
  const a = hexToRgb(from)
  const b = hexToRgb(to)
  return `#${a.map((channel, index) => toHex(Math.round(channel * amount + b[index] * (1 - amount)))).join("")}`
}

const isHexColor = (value: string) => /^#[0-9a-fA-F]{6}$/.test(value)

const normalizeThemeSeed = (seed: ThemeSeed): ThemeSeed | null => {
  if (!seed.key || !seed.name || !seed.image) return null

  const fallback = themeSeeds[defaultThemeKey]

  return {
    key: seed.key,
    name: seed.name.trim() || "自定义主题",
    desc: seed.desc?.trim() || "高度客制化主题",
    family: seed.family === "custom" ? "custom" : "molly",
    primary: isHexColor(seed.primary) ? seed.primary : fallback.primary,
    secondary: isHexColor(seed.secondary) ? seed.secondary : fallback.secondary,
    accent: isHexColor(seed.accent) ? seed.accent : fallback.accent,
    warning: isHexColor(seed.warning) ? seed.warning : fallback.warning,
    mollyHair: isHexColor(seed.mollyHair) ? seed.mollyHair : fallback.mollyHair,
    mollyMouth: isHexColor(seed.mollyMouth) ? seed.mollyMouth : fallback.mollyMouth,
    image: seed.image,
  }
}

export const setCustomThemeSeeds = (themes?: Record<string, ThemeSeed>) => {
  customThemeSeeds = {}

  Object.values(themes ?? {}).forEach((theme) => {
    const seed = normalizeThemeSeed({...theme, family: "custom"})
    if (seed && !themeSeeds[seed.key]) {
      customThemeSeeds[seed.key] = seed
    }
  })

  themeCache.clear()
}

export const getThemeSeeds = () => ({
  ...themeSeeds,
  ...customThemeSeeds,
})

export const getThemeOptions = () => Object.values(getThemeSeeds())

export const hasTheme = (themeKey?: string) => Boolean(themeKey && getThemeSeeds()[themeKey])

export const themeImageUrl = (image: string, baseUrl = "") => {
  if (/^[a-zA-Z]:[\\/]/.test(image) || image.startsWith("\\\\") || image.startsWith("/")) {
    const prefix = baseUrl || `${window.location.protocol}//${window.location.host}`
    return `${prefix}/api/local-image?path=${encodeURIComponent(image)}`
  }

  return new URL(`../assets/image/molly/${image}`, import.meta.url).href
}

const makeLightPreset = (base: ThemeSeed): ThemePreset => {
  const primary = mixHex(base.primary, "#6f7686", 0.72)
  const primaryHover = mixHex(base.primary, "#40495d", 0.66)
  const primaryPressed = mixHex(base.primary, "#1f2937", 0.70)
  const secondary = mixHex(base.secondary, "#7b8174", 0.62)
  const accent = mixHex(base.accent, "#69727e", 0.68)
  const warning = mixHex(base.warning, "#b8842c", 0.62)
  const rgb = hexToRgb(primary).join(", ")

  return {
    ...base,
    mode: "light",
    primary,
    primaryHover,
    primaryPressed,
    secondary,
    error: "#c94b5f",
    warning,
    background: "#f7f4ef",
    surface: "#fffdf9",
    surfaceContainer: "#ece8df",
    surfaceContainerHigh: "#ddd8cf",
    outline: "#c9c1b6",
    text: "#2e2b27",
    textMuted: "#716b63",
    ...makeShadow(rgb),
    glow: `rgba(${rgb}, 0.24)`,
    accent,
    molly: "#ead4c8",
    mollyHair: mixHex(base.mollyHair, "#caa45c", 0.50),
    mollyEye: accent,
    mollyMouth: mixHex(base.mollyMouth, "#a84f63", 0.56),
    swatch: `linear-gradient(135deg, #fffdf9 0%, #ece8df 22%, ${primary} 58%, ${accent} 100%)`,
  }
}

const makeDarkPreset = (base: ThemeSeed): ThemePreset => {
  const primary = mixHex(base.primary, "#cdd4e5", 0.58)
  const primaryHover = mixHex(base.primary, "#eef2ff", 0.48)
  const primaryPressed = mixHex(base.primary, "#68718a", 0.70)
  const secondary = mixHex(base.secondary, "#d5d2bd", 0.56)
  const accent = mixHex(base.accent, "#d4e0e4", 0.58)
  const warning = mixHex(base.warning, "#e2c16a", 0.56)
  const rgb = hexToRgb(primary).join(", ")

  return {
    ...base,
    mode: "dark",
    primary,
    primaryHover,
    primaryPressed,
    secondary,
    error: "#ff8a9d",
    warning,
    background: "#101217",
    surface: "#171a22",
    surfaceContainer: "#20242f",
    surfaceContainerHigh: "#2b3040",
    outline: "#4b5363",
    text: "#f4f2ef",
    textMuted: "#c7c1b8",
    ...makeShadow(rgb),
    glow: `rgba(${rgb}, 0.30)`,
    accent,
    molly: "#3a3130",
    mollyHair: mixHex(base.mollyHair, "#dfc27a", 0.44),
    mollyEye: accent,
    mollyMouth: mixHex(base.mollyMouth, "#f1a4b3", 0.52),
    swatch: `linear-gradient(135deg, #101217 0%, #2b3040 28%, ${primary} 66%, ${accent} 100%)`,
  }
}

export const resolveTheme = (themeKey?: string, mode: ThemeMode = "light") => {
  const safeMode = mode === "dark" ? "dark" : "light"
  const base = getThemeSeeds()[themeKey || ""] ?? themeSeeds[defaultThemeKey]
  const cacheKey = `${base.key}:${safeMode}`
  const cached = themeCache.get(cacheKey)

  if (cached) return cached

  const preset = safeMode === "dark" ? makeDarkPreset(base) : makeLightPreset(base)
  themeCache.set(cacheKey, preset)
  return preset
}

export const resolveThemeSeed = (seed: ThemeSeed, mode: ThemeMode = "light") =>
  mode === "dark" ? makeDarkPreset(seed) : makeLightPreset(seed)

export const resolveThemeOptions = (mode: ThemeMode = "light") =>
  getThemeOptions().map(item => resolveTheme(item.key, mode))
