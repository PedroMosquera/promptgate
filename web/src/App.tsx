import { Route, Routes } from "react-router-dom";
import SessionsList from "./pages/SessionsList";
import SessionDetail from "./pages/SessionDetail";

export default function App() {
  return (
    <div className="app-shell">
      <header className="app-header">
        <h1>promptgate console</h1>
        <span className="subtitle">agent session observability</span>
      </header>
      <Routes>
        <Route path="/" element={<SessionsList />} />
        <Route path="/sessions/:id" element={<SessionDetail />} />
      </Routes>
    </div>
  );
}
