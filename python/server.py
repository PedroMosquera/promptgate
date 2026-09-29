"""A small HTTP service over the agent session traces in fixtures/sessions/.

Standard library only. Run with `python3 server.py`.
"""
import json
import os
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer

from cost import session_cost
from sessions import detail, list_session_ids, summarize


class Handler(BaseHTTPRequestHandler):
    def _write_json(self, status, body):
        payload = json.dumps(body).encode("utf-8")
        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(payload)))
        self.end_headers()
        self.wfile.write(payload)

    def do_GET(self):
        if self.path == "/api/sessions":
            sessions = []
            for session_id in list_session_ids():
                summary = summarize(session_id)
                summary["cost_usd"] = round(session_cost(summary), 4)
                sessions.append(summary)
            self._write_json(200, sessions)
            return

        if self.path.startswith("/api/sessions/"):
            session_id = self.path[len("/api/sessions/"):]
            if session_id not in list_session_ids():
                self._write_json(404, {"error": f"no such session: {session_id}"})
                return
            payload = detail(session_id)
            payload["summary"]["cost_usd"] = round(session_cost(payload["summary"]), 4)
            self._write_json(200, payload)
            return

        self._write_json(404, {"error": "not found"})

    def log_message(self, fmt, *args):
        pass  # keep test and manual-run output quiet


def main():
    port = int(os.environ.get("PORT", "8081"))
    server = ThreadingHTTPServer(("", port), Handler)
    print(f"listening on :{port}")
    server.serve_forever()


if __name__ == "__main__":
    main()
