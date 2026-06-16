# app

Template for a single-page web app. Pre-wired, opinionated, strict.

## Stack

- **[Bun](https://bun.sh)** — runtime, package manager, script runner
- **[Vite](https://vite.dev)** — dev server and bundler
- **[React 19](https://react.dev)** — UI library
- **[TypeScript](https://www.typescriptlang.org)** — strictest practical config (`strict`, `exactOptionalPropertyTypes`, `noUncheckedIndexedAccess`, `noImplicitOverride`, `noPropertyAccessFromIndexSignature`, …). `@/*` is aliased to `src/*`.
- **[Tailwind CSS v4](https://tailwindcss.com)** — via `@tailwindcss/vite` (no PostCSS config). Theme tokens live in `src/index.css`.
- **[shadcn/ui](https://ui.shadcn.com)** — copy-in components under `src/components/ui` (style: `new-york`, base color: `neutral`). Add components with `bunx shadcn@latest add <name>`.
- **[lucide-react](https://lucide.dev)** — generic icon set (shadcn default).
- **[@icons-pack/react-simple-icons](https://github.com/icons-pack/react-simple-icons)** — typed React wrappers around [simple-icons](https://simpleicons.org) (~3000 brand icons: Google, Apple, X, GitHub, Discord, Slack, …).
- **[Biome](https://biomejs.dev)** — lint + format + import organizing (replaces ESLint + Prettier).

## Layout

```
src/
  components/
    ui/              # shadcn components (excluded from Biome)
    login-card.tsx   # bundled login card (see below)
  lib/utils.ts       # cn() helper
  App.tsx
  main.tsx
  index.css          # Tailwind + theme tokens
```

## `LoginCard`

The template ships with `src/components/login-card.tsx` — a configurable login card derived from shadcn's `login-03` block. It's the main pre-built piece in this template.

### Props

```ts
type LoginCardProps = {
  title: ReactNode;          // CardTitle content
  description?: ReactNode;   // CardDescription (hidden if omitted)
  magicLink?: boolean;       // default true — show magic-link email form
  social?: SocialOption[];   // default [] — list of social buttons, in order
  onMagicLink?: (email: string) => void;
};

type SocialOption = {
  label: string;       // shown as "Continue with {label}"
  redirect: string;    // URL the button navigates to
  icon?: ReactNode;    // any icon element (lucide-react, simple-icons, …)
};
```

### Design choices

- **Social = "Continue with", not "Login with".** Social auth is meant to be seamless — if the user doesn't have an account yet, the OAuth callback handler should create one transparently. The button copy reflects that: one path, no separate "Sign up with Google" flow.
- **Email = magic link only, no password.** The email field is wired for passwordless ("Send me a magic link") — no password, no "Forgot password?", no confirm-password. The UI is bundled but **not wired to a backend**; you provide `onMagicLink(email)` and call your own send-link endpoint. Same seamless idea as social: if the email doesn't have an account yet, the link-handler creates one.
- **No email/password signup component, by design.** Password-based auth (signup or login) is intentionally not part of this template. It implies email verification, password rules, confirm-password fields, reset flows — all of which belong in their own routes. The expected default here is: social or magic link, both of which collapse login and signup into a single path.
- **Social = links, not callbacks.** Each social button renders as an `<a href={redirect}>` (via shadcn's `asChild`). That gives you real link semantics — middle-click / cmd-click open in a new tab, hover shows the URL, right-click → Copy link works, and OAuth still works with JS disabled. OAuth needs a full-page redirect anyway, so a callback would just get in the way.
- **Magic link = handler.** The email form calls `onMagicLink(email)`. `event.preventDefault()` is already handled — you do the API call.
- **Icons are arbitrary.** `icon` is a `ReactNode`, so anything renders. Sizing is inherited from the Button's `[&_svg:not([class*='size-'])]:size-4` rule.
- **Separator auto-hides.** "Or continue with" only appears when both `magicLink` is true and `social` has at least one entry.

### Example

```tsx
import { SiApple, SiGithub, SiGoogle, SiX } from "@icons-pack/react-simple-icons";
import { LoginCard } from "@/components/login-card";

<LoginCard
  title="Welcome back"
  description="Continue with your social account or email"
  magicLink
  social={[
    { label: "Google", redirect: "/auth/google", icon: <SiGoogle /> },
    { label: "Apple",  redirect: "/auth/apple",  icon: <SiApple /> },
    { label: "X",      redirect: "/auth/x",      icon: <SiX /> },
    { label: "GitHub", redirect: "/auth/github", icon: <SiGithub /> },
  ]}
  onMagicLink={async (email) => {
    await api.sendMagicLink(email);
  }}
/>
```

### Picking icons

- **Brand icons** (Google, Discord, Spotify, …) → `@icons-pack/react-simple-icons`:
  ```tsx
  import { SiDiscord } from "@icons-pack/react-simple-icons";
  ```
- **Generic icons** (Mail, Key, Lock, …) → `lucide-react`:
  ```tsx
  import { Mail } from "lucide-react";
  ```
- **Custom SVG** — pass any `<svg>` JSX with `fill="currentColor"` and a `viewBox`.

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
