import { private_env } from "./shared-server-aoX5LlPz.js";
import { jwtVerify } from "jose";

//#region .svelte-kit/adapter-bun/chunks/hooks.server.js
const SECRET = new TextEncoder().encode(private_env.JWT_SECRET.trim());
const handle = async ({ event, resolve }) => {
	const token = event.cookies.get("token");
	if (token) try {
		const { payload } = await jwtVerify(token, SECRET);
		event.locals.user = payload;
	} catch (err) {
		event.locals.user = null;
		console.warn("JWT validation failed:", err.message);
		event.cookies.delete("session", { path: "/" });
	}
	else event.locals.user = null;
	return await resolve(event);
};

//#endregion
export { handle };
//# sourceMappingURL=hooks.server-TpyQUb9V.js.map