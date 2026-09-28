// Payloads are static files, so responses land instantly. Scale a delay with
// payload size to keep local behaviour close to the deployed gateway.
async function delayFor(bytes: number): Promise<void> {
  const ms = Math.min(1500, Math.max(100, Math.round(bytes / 9)));
  await new Promise((resolve) => setTimeout(resolve, ms));
}

export async function fetchJson<T>(path: string): Promise<T> {
  const res = await fetch(path);
  if (!res.ok) throw new Error(`${res.status} ${res.statusText} for ${path}`);
  const text = await res.text();
  await delayFor(text.length);
  return JSON.parse(text) as T;
}
