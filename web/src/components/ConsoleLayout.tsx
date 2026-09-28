import { Outlet } from "react-router-dom";
import SessionsList from "../pages/SessionsList";

// The list stays mounted across every route so an operator can click
// straight from one session into another without returning to "/" first.
export default function ConsoleLayout() {
  return (
    <div className="console-layout">
      <div className="console-sidebar">
        <SessionsList />
      </div>
      <div className="console-detail">
        <Outlet />
      </div>
    </div>
  );
}
