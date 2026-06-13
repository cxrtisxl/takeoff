import { useState } from "react";
import { Button } from "@/components/ui/button";

function App() {
  const [count, setCount] = useState(0);

  return (
    <main className="min-h-svh flex flex-col items-center justify-center gap-6 p-8">
      <h1 className="text-4xl font-semibold tracking-tight">Vite + React</h1>
      <Button onClick={() => setCount((c) => c + 1)}>Count is {count}</Button>
      <p className="text-muted-foreground text-sm">
        Edit <code className="font-mono">src/App.tsx</code> and save to test HMR
      </p>
    </main>
  );
}

export default App;
