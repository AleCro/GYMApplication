import { API_URL } from '$lib/config.js';


/** @type {import('@sveltejs/kit').RequestHandler} */
export async function GET({ url }) {
	const muscle = url.searchParams.get('muscle');

	if (!muscle) {
		return new Response(JSON.stringify({ error: 'Muscle parameter is required' }), {
			status: 400,
			headers: { 'Content-Type': 'application/json' }
		});
	}

	try {
		const res = await fetch(`${API_URL}/exercise?muscle=${muscle}`);
		if (!res.ok) throw new Error('Failed to fetch exercises');
		const data = await res.json();

		return new Response(JSON.stringify(data), {
			status: 200,
			headers: { 'Content-Type': 'application/json' }
		});
	} catch (err) {
		return new Response(JSON.stringify({ error: err.message }), {
			status: 500,
			headers: { 'Content-Type': 'application/json' }
		});
	}
}

export async function POST({ request }) {
  const body = await request.text();
  const res = await fetch(`${API_URL}/exercise`, {
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