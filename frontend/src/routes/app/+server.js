export async function POST({ locals }) {
	if (!locals.user) {
		return Response.json({});
	}
	return Response.json(locals.user);
}