import { escape_html, getContext, head, store_get, unsubscribe_stores } from "./index2-BLSeo-Vt.js";
import "./internal-CyqLiTQC.js";
import "./server-CzoYAPU4.js";
import "./utils-DKQ8rzhJ.js";
import "./clsx-cC83_lR5.js";
import "./state.svelte-kTmDZfx6.js";

//#region .svelte-kit/adapter-bun/entries/pages/_error.svelte.js
const getStores = () => {
	const stores = getContext("__svelte__");
	return {
		page: { subscribe: stores.page.subscribe },
		navigating: { subscribe: stores.navigating.subscribe },
		updated: stores.updated
	};
};
const page = { subscribe(fn) {
	const store = getStores().page;
	return store.subscribe(fn);
} };
function _error($$renderer, $$props) {
	$$renderer.component(($$renderer2) => {
		var $$store_subs;
		let status, message, title, path;
		status = store_get($$store_subs ??= {}, "$page", page).status || 500;
		message = store_get($$store_subs ??= {}, "$page", page).error?.message || "An unexpected error occurred.";
		title = status === 404 ? "Page Not Found" : "Server Error";
		path = store_get($$store_subs ??= {}, "$page", page).url.pathname;
		head($$renderer2, ($$renderer3) => {
			$$renderer3.title(($$renderer4) => {
				$$renderer4.push(`<title>${escape_html(status)} ${escape_html(title)} | Starter API</title>`);
			});
		});
		$$renderer2.push(`<div class="flex flex-col items-center justify-center text-center p-8 space-y-6 min-h-screen"><h1 class="text-7xl font-extrabold text-indigo-400 tracking-tight">${escape_html(status)}</h1> <h2 class="text-4xl font-bold text-white">${escape_html(title)}</h2> <p class="text-xl text-gray-400 max-w-lg">${escape_html(message)}</p> <p class="text-sm text-gray-500">Attempted path: <code class="font-mono">${escape_html(path)}</code></p> <div class="flex flex-col sm:flex-row justify-center gap-4 pt-4"><a href="/app" class="px-6 py-3 bg-indigo-600 text-white font-semibold rounded-lg shadow-lg hover:bg-indigo-700 transition duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-indigo-500 focus-visible:ring-offset-2 focus-visible:ring-offset-gray-900">Go to App</a> <a href="/" class="px-6 py-3 border border-gray-600 text-gray-300 font-semibold rounded-lg hover:bg-gray-800 transition duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gray-400 focus-visible:ring-offset-2 focus-visible:ring-offset-gray-900">Return Home</a></div></div>`);
		if ($$store_subs) unsubscribe_stores($$store_subs);
	});
}

//#endregion
export { _error as default };
//# sourceMappingURL=_error.svelte-HbWhFoiv.js.map