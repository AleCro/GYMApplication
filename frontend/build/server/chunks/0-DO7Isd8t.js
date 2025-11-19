import { __export } from "./exports-BwVzmOlf.js";
import "./internal-CyqLiTQC.js";

//#region .svelte-kit/adapter-bun/entries/pages/_layout.server.js
var _layout_server_exports = {};
__export(_layout_server_exports, { load: () => load });
const load = async ({ locals }) => {
	return { user: locals.user };
};

//#endregion
//#region .svelte-kit/adapter-bun/nodes/0.js
const index = 0;
let component_cache;
const component = async () => component_cache ??= (await import("./layout.svelte-DoRVC-M7.js")).default;
const server_id = "src/routes/+layout.server.js";
const imports = [
	"_app/immutable/nodes/0.CPihVZeV.js",
	"_app/immutable/chunks/DsnmJJEf.js",
	"_app/immutable/chunks/DFFLsmBn.js",
	"_app/immutable/chunks/U_jKa9eL.js"
];
const stylesheets = [];
const fonts = [];

//#endregion
export { component, fonts, imports, index, _layout_server_exports as server, server_id, stylesheets };
//# sourceMappingURL=0-DO7Isd8t.js.map