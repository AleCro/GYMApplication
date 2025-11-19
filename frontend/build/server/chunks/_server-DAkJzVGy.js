//#region .svelte-kit/adapter-bun/entries/endpoints/app/_server.js
async function POST({ locals }) {
	if (!locals.user) return Response.json({});
	return Response.json(locals.user);
}

//#endregion
export { POST };
//# sourceMappingURL=_server-DAkJzVGy.js.map