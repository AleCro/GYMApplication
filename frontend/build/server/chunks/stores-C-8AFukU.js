import { ssr_context } from "./index2-BLSeo-Vt.js";
import { writable } from "./chunks-DNg0X1yh.js";

//#region .svelte-kit/adapter-bun/chunks/stores.js
function onDestroy(fn) {
	/** @type {SSRContext} */
	ssr_context.r.on_destroy(fn);
}
const userStore = writable(null);

//#endregion
export { onDestroy, userStore };
//# sourceMappingURL=stores-C-8AFukU.js.map