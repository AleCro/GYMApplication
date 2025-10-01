import { API_URL } from '$lib/config.js';

export async function POST({ request }) {
  const { notes, session } = await request.json();

  try {
    const res = await fetch(`${API_URL}/notes`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ notes, session })
    });

    const data = await res.json();
    return new Response(JSON.stringify(data), {
      headers: { 'Content-Type': 'application/json' }
    });
  } catch (err) {
    return new Response(JSON.stringify({ error: err.message }), { status: 500 });
  }
}
