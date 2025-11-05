// src/routes/api/progress/+server.js
import { API_URL } from '$lib/config.js';

// This central POST handler receives all API requests (add/get/remove progress)
export async function POST({ request }) {
	let endpoint = '';
	let dataPayload = {};

	try {
		// Read the entire body once
		const bodyText = await request.text();
		const clientBody = JSON.parse(bodyText);

		// 1. Extract the target endpoint
		endpoint = clientBody.endpoint;
		if (!endpoint) {
			return new Response(JSON.stringify({ error: 'Endpoint not specified' }), { status: 400 });
		}
		
		// 2. The rest of the data is the payload for the Go backend
		dataPayload = clientBody.data || {};

	} catch (e) {
		console.error("Proxy JSON parsing error:", e);
		return new Response(JSON.stringify({ error: 'Invalid JSON format sent from client.' }), { status: 400 });
	}

	// 3. Forward request to the specific Go backend endpoint
	const targetURL = `${API_URL}/${endpoint}`;
	
	try {
		const res = await fetch(targetURL, {
			method: 'POST',
			headers: { 
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(dataPayload), // Send payload directly
		});
		
		// 4. **CRITICAL FIX:** Read the response body *once* and return it with the correct status/headers.
        // This ensures the Svelte client gets exactly what the Go server sent, whether it's valid JSON or an error message.
		const responseBody = await res.arrayBuffer();
        const headers = new Headers(res.headers);
        headers.delete('content-encoding');
        headers.delete('content-length');
        headers.set('access-control-allow-origin', '*');

		// Return the raw response from the Go server back to Svelte
		return new Response(responseBody, { status: res.status, headers });

	} catch (err) {
		console.error(`Error forwarding request to ${targetURL}:`, err);
		return new Response(JSON.stringify({ error: 'Failed to reach backend API.' }), { status: 503 });
	}
}