//#region node_modules/@sveltejs/kit/src/runtime/server/constants.js
const IN_WEBCONTAINER = !!globalThis.process?.versions?.webcontainer;

//#endregion
//#region node_modules/@sveltejs/kit/src/exports/internal/event.js
/** @type {RequestStore | null} */
let sync_store = null;
/** @type {AsyncLocalStorage<RequestStore | null> | null} */
let als;
import("node:async_hooks").then((hooks) => als = new hooks.AsyncLocalStorage()).catch(() => {});
/**
* @template T
* @param {RequestStore | null} store
* @param {() => T} fn
*/
function with_request_store(store, fn) {
	try {
		sync_store = store;
		return als ? als.run(store, fn) : fn();
	} finally {
		if (!IN_WEBCONTAINER) sync_store = null;
	}
}

//#endregion
//#region node_modules/@sveltejs/kit/src/exports/internal/server.js
/**
* @template {{ tracing: { enabled: boolean, root: import('@opentelemetry/api').Span, current: import('@opentelemetry/api').Span } }} T
* @param {T} event_like
* @param {import('@opentelemetry/api').Span} current
* @returns {T}
*/
function merge_tracing(event_like, current) {
	return {
		...event_like,
		tracing: {
			...event_like.tracing,
			current
		}
	};
}

//#endregion
export { merge_tracing, with_request_store };
//# sourceMappingURL=server-CzoYAPU4.js.map