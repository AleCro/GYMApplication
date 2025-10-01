const API_URL = 'https://api-gym.alecro.click';

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
