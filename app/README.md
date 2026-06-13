# app

Template for a single-page web app. Pre-wired, opinionated, strict.

## Stack

- **[Bun](https://bun.sh)** — runtime, package manager, script runner
- **[Vite](https://vite.dev)** — dev server and bundler
- **[React 19](https://react.dev)** — UI library
- **[TypeScript](https://www.typescriptlang.org)** — strictest practical config (`strict`, `exactOptionalPropertyTypes`, `noUncheckedIndexedAccess`, `noImplicitOverride`, `noPropertyAccessFromIndexSignature`, …). `@/*` is aliased to `src/*`.
- **[Tailwind CSS v4](https://tailwindcss.com)** — via `@tailwindcss/vite` (no PostCSS config). Theme tokens live in `src/index.css`.
- **[shadcn/ui](https://ui.shadcn.com)** — copy-in components under `src/components/ui` (style: `new-york`, base color: `neutral`, icons: `lucide-react`). Add components with `bunx shadcn@latest add <name>`.
- **[Biome](https://biomejs.dev)** — lint + format + import organizing (replaces ESLint + Prettier).

## Layout

```
src/
  components/ui/   # shadcn components (excluded from Biome)
  lib/utils.ts     # cn() helper
  App.tsx
  main.tsx
  index.css        # Tailwind + theme tokens
```

## Scripts

| Script | What it does |
| --- | --- |
| `bun dev` / `bun start` | Start Vite dev server |
| `bun run build` | Type-check (`tsc -b`) then build with Vite |
| `bun run preview` | Preview the production build |
| `bun run typecheck` | `tsc -b --noEmit` |
| `bun run lint` | Biome lint (read-only) |
| `bun run lint:fix` | Biome lint with safe fixes |
| `bun run format` | Biome format (writes) |
| `bun run format:check` | Biome format (check only) |
| `bun run check` | Biome lint + format + organize imports (read-only) |
| `bun run check:fix` | Same, but writes fixes |
| `bun run ci` | `biome ci .` — read-only, for CI pipelines |

## Getting started

```sh
bun install
bun dev
```

Add a shadcn component:

```sh
bunx shadcn@latest add card dialog input
```
