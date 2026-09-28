import { Route, Routes } from "react-router-dom";
import ConsoleLayout from "./components/ConsoleLayout";
import SessionDetail from "./pages/SessionDetail";

export default function App() {
  return (
    <div className="app-shell">
      <header className="app-header">
        <h1>promptgate console</h1>
        <span className="subtitle">agent session observability</span>
      </header>
      <Routes>
        <Route path="/" element={<ConsoleLayout />}>
          <Route index element={<p className="loading">Select a session.</p>} />
          <Route path="sessions/:id" element={<SessionDetail />} />
        </Route>
      </Routes>
    </div>
  );
}
