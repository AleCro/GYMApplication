// src/routes/login/+server.js
import { API_URL } from '$lib/config.js'; // <-- Add this

export async function POST({ request, cookies }) {
  const { username, password } = await request.json();

  const res = await fetch(`${API_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password })
  });

  const data = await res.json();

  if (res.ok) {
    // Save session in cookie
    cookies.set('session', data.session, {
      path: '/',
      httpOnly: true,
      sameSite: 'strict'
    });
  }

  return new Response(JSON.stringify(data), {
    headers: { "Content-Type": "application/json" }
  });
}
