import { Route, Routes } from "react-router-dom";
import ConsoleLayout from "./components/ConsoleLayout";
import Overview from "./pages/Overview";
import SessionDetail from "./pages/SessionDetail";

export default function App() {
  return (
    <div className="app-shell">
      <header className="app-header">
        <h1>promptgate on-call console</h1>
        <span className="subtitle">agent session health for the on-call rotation</span>
      </header>
      <Routes>
        <Route path="/" element={<ConsoleLayout />}>
          <Route index element={<Overview />} />
          <Route path="sessions/:id" element={<SessionDetail />} />
        </Route>
      </Routes>
    </div>
  );
}
