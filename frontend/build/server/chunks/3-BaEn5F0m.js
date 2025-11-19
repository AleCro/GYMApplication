import { __export, redirect } from "./exports-BwVzmOlf.js";
import "./internal-CyqLiTQC.js";

//#region .svelte-kit/adapter-bun/entries/pages/login/_layout.server.js
var _layout_server_exports = {};
__export(_layout_server_exports, { load: () => load });
const load = async ({ locals }) => {
	if (locals?.user) throw redirect(302, "/app");
	return {};
};

//#endregion
//#region .svelte-kit/adapter-bun/nodes/3.js
const index = 3;
let component_cache;
const component = async () => component_cache ??= (await import("./_layout.svelte-kwR-KxmS.js")).default;
const server_id = "src/routes/login/+layout.server.js";
const imports = [
	"_app/immutable/nodes/3.BYX7BXg3.js",
	"_app/immutable/chunks/DsnmJJEf.js",
	"_app/immutable/chunks/hNqiKWnM.js",
	"_app/immutable/chunks/DFFLsmBn.js",
	"_app/immutable/chunks/CxnIUwRS.js"
];
const stylesheets = [];
const fonts = [];

//#endregion
export { component, fonts, imports, index, _layout_server_exports as server, server_id, stylesheets };
//# sourceMappingURL=3-BaEn5F0m.js.map