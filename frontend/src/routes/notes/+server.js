import { API_URL } from '$lib/config.js';

export async function POST({ request }) {
  const body = await request.text();

  const res = await fetch(`${API_URL}/addnote`, {
    method: "POST",
    headers: {
      "Content-Type": request.headers.get("Content-Type") || "application/json",
    },
    body,
  });

  const headers = new Headers(res.headers);
  headers.delete('content-encoding');
  headers.delete('content-length');
  headers.set('access-control-allow-origin', '*');

  const responseBody = await res.arrayBuffer();

  return new Response(responseBody, {
    status: res.status,
    headers,
  });
}

export async function PATCH({ request }) {
  const body = await request.text();

  const res = await fetch(`${API_URL}/updatenote`, {
    method: "POST",
    headers: {
      "Content-Type": request.headers.get("Content-Type") || "application/json",
    },
    body,
  });

  const headers = new Headers(res.headers);
  headers.delete('content-encoding');
  headers.delete('content-length');
  headers.set('access-control-allow-origin', '*');

  const responseBody = await res.arrayBuffer();

  return new Response(responseBody, {
    status: res.status,
    headers,
  });
}

export async function DELETE({ request }) {
  const body = await request.text();
  const res = await fetch(`${API_URL}/deletenote`, {
    method: "POST",
    headers: {
      "Content-Type": request.headers.get("Content-Type") || "application/json",
    },
    body,
  });

  const headers = new Headers(res.headers);
  headers.delete('content-encoding');
  headers.delete('content-length');
  headers.set('access-control-allow-origin', '*');

  const responseBody = await res.arrayBuffer();

  return new Response(responseBody, {
    status: res.status,
    headers,
  });
}