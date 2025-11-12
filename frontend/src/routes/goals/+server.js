import { API_URL } from '$lib/config.js';

export async function POST({ request }) {
	try {
		const contentType = request.headers.get('content-type') || '';

		// -------------------------------
		// Handle file uploads (photo upload)
		// -------------------------------
		if (contentType.includes('multipart/form-data')) {
			const res = await fetch(`${API_URL}/upload`, {
				method: 'POST',
				body: request.body,
				duplex: 'half'
			});
			const data = await res.json();
			return new Response(JSON.stringify(data), {
				status: res.status,
				headers: { 'Content-Type': 'application/json' }
			});
		}

		// -------------------------------
		// Handle JSON requests
		// -------------------------------
		const body = await request.json();
		let { endpoint, data } = body || {};

		if (!endpoint) {
			return new Response(
				JSON.stringify({ error: 'Missing endpoint name' }),
				{ status: 400, headers: { 'Content-Type': 'application/json' } }
			);
		}

		// ✅ Special handling: Add Step → forward to /updategoal with flag
		if (endpoint === 'addstep') {
			endpoint = 'updategoal';
			data.newStep = true;
		}

		// Forward request to Go backend
		const res = await fetch(`${API_URL}/${endpoint}`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(data)
		});

		// Preserve backend response exactly
		const text = await res.text();
		return new Response(text, {
			status: res.status,
			headers: {
				'Content-Type': res.headers.get('content-type') || 'application/json'
			}
		});
	} catch (err) {
		console.error('[Progress Proxy Error]', err);
		return new Response(
			JSON.stringify({
				error: 'Failed to contact backend',
				details: err.message
			}),
			{ status: 502, headers: { 'Content-Type': 'application/json' } }
		);
	}
}
