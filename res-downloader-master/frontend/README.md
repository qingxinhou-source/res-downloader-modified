# Frontend

Vue 3 + TypeScript + Vite frontend for the Wails app.

## Development

```bash
npm install
npm run dev
```

## Build

```bash
npm run build
```

## Theme System

- Theme presets live in `src/theme/presets.ts`
- Molly transparent PNG assets live in `src/assets/image/molly`
- Custom themes are persisted through the app config as `CustomThemes`
- Light/dark mode is global: switching themes keeps the current mode

The settings page provides the built-in blind-box theme cards and the custom theme editor.
