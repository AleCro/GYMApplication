//#region .svelte-kit/adapter-bun/chunks/shared-server.js
let private_env = {};
let public_env = {};
function set_private_env(environment) {
	private_env = environment;
}
function set_public_env(environment) {
	public_env = environment;
}

//#endregion
export { private_env, public_env, set_private_env, set_public_env };
//# sourceMappingURL=shared-server-aoX5LlPz.js.map