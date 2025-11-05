import { API_URL } from '$lib/config.js';

export async function POST({ request }) {
	try {
		const body = await request.json();

		if (!body.endpoint) {
			return new Response(
				JSON.stringify({ error: 'Missing endpoint field' }),
				{ status: 400, headers: { 'Content-Type': 'application/json' } }
			);
		}

		// Accept both { endpoint, data: {...} } and { endpoint, ...rest }
		const { endpoint, data, ...rest } = body;
		const payload = data || rest;

		// Forward request to backend
		
		const res = await fetch(`${API_URL}/${endpoint}`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(payload)
		});

		const headers = new Headers(res.headers);
		headers.delete('content-encoding');
		headers.delete('content-length');
		headers.set('access-control-allow-origin', '*');

		const responseBody = await res.arrayBuffer();

		return new Response(responseBody, {
			status: res.status,
			headers
		});
	} catch (err) {
		console.error('Goal proxy error:', err);
		return new Response(
			JSON.stringify({ error: 'Failed to contact backend', details: err.message }),
			{ status: 500, headers: { 'Content-Type': 'application/json' } }
		);
	}
}
