export namespace appType {
    interface App {
        AppName: string
        Version: string
        Description: string
        Copyright: string
    }

    interface MimeMap {
        Type: string
        Suffix: string
    }

    interface Config {
        Theme: string
        ThemeMode: "light" | "dark"
        CustomThemes: { [key: string]: CustomTheme }
        Personalization: Personalization
        Locale: string
        Host: string
        Port: string
        Quality: number
        SaveDirectory: string
        FilenameLen: number
        FilenameTime: boolean
        UpstreamProxy: string
        OpenProxy: boolean
        DownloadProxy: boolean
        AutoProxy: boolean
        WxAction: boolean
        TaskNumber: number
        DownNumber: number
        UserAgent: string
        UseHeaders: string
        InsertTail: boolean
        MimeMap: { [key: string]: MimeMap }
        Rule: string
    }

    interface CustomTheme {
        key: string
        name: string
        desc: string
        family: "custom"
        primary: string
        secondary: string
        accent: string
        warning: string
        mollyHair: string
        mollyMouth: string
        image: string
    }

    interface Personalization {
        ShowMetrics: boolean
        ShowFooterLinks: boolean
        ShowMollyCompanion: boolean
        ShowMollyBackground: boolean
        VisualFeedback: boolean
        ShowClock: boolean
        CustomWidgets: CustomWidget[]
    }

    interface CustomWidget {
        id: string
        type: "note" | "clock" | "action"
        title: string
        content: string
        action: "none" | "toggleProxy" | "clearList" | "toggleThemeMode"
        x: number
        y: number
        visible: boolean
    }

    interface MediaInfo {
        Id: string
        Url: string
        UrlSign: string
        CoverUrl: string
        Size: number
        Domain: string
        Classify: string
        Suffix: string
        SavePath: string
        Status: string
        DecodeKey: string
        Description: string
        ContentType: string
        OtherData: { [key: string]: string }
    }

    interface DownloadProgress {
        Id: string
        SavePath: string
        Status: string
        Message: string
        Size?: number
    }

    interface Message {
        code: number
        message: string
    }

    interface Handle {
        type: string
        event: any
    }

    interface Res<T = any> {
        code: number;
        message: string;
        data: T;  // T will be the specific type of your data
    }
}
