import {
  SiApple,
  SiGithub,
  SiGoogle,
  SiX,
} from "@icons-pack/react-simple-icons";

import { LoginCard } from "@/components/login-card";

function App() {
  return (
    <main className="min-h-svh flex items-center justify-center p-6">
      <LoginCard
        title="Welcome to Takeoff!"
        description="Continue with your social account or email"
        magicLink={true}
        social={[
          { label: "Google", redirect: "/auth/google", icon: <SiGoogle /> },
          { label: "Apple", redirect: "/auth/apple", icon: <SiApple /> },
          { label: "X", redirect: "/auth/x", icon: <SiX /> },
          { label: "GitHub", redirect: "/auth/github", icon: <SiGithub /> },
        ]}
      />
    </main>
  );
}

export default App;
