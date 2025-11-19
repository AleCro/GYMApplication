import { __export, redirect } from "./exports-BwVzmOlf.js";
import "./internal-CyqLiTQC.js";

//#region .svelte-kit/adapter-bun/entries/pages/app/_layout.server.js
var _layout_server_exports = {};
__export(_layout_server_exports, { load: () => load });
const load = async ({ locals }) => {
	if (!locals.user) throw redirect(302, "/login");
	return { user: locals.user };
};

//#endregion
//#region .svelte-kit/adapter-bun/nodes/2.js
const index = 2;
let component_cache;
const component = async () => component_cache ??= (await import("./_layout.svelte-VfyOVXJq.js")).default;
const server_id = "src/routes/app/+layout.server.js";
const imports = [
	"_app/immutable/nodes/2.DOnDJr5_.js",
	"_app/immutable/chunks/DsnmJJEf.js",
	"_app/immutable/chunks/hNqiKWnM.js",
	"_app/immutable/chunks/DFFLsmBn.js",
	"_app/immutable/chunks/1-JAo-Vw.js",
	"_app/immutable/chunks/CVaLoUw_.js",
	"_app/immutable/chunks/U_jKa9eL.js",
	"_app/immutable/chunks/DIYmMOxU.js",
	"_app/immutable/chunks/wBDlzmEs.js",
	"_app/immutable/chunks/CxnIUwRS.js",
	"_app/immutable/chunks/BPl-Gnor.js",
	"_app/immutable/chunks/iA-kc8Oh.js",
	"_app/immutable/chunks/BUx8HaOQ.js",
	"_app/immutable/chunks/Co19WjoH.js",
	"_app/immutable/chunks/NJrgjD4c.js",
	"_app/immutable/chunks/BOruTPHP.js",
	"_app/immutable/chunks/CdP2dVa8.js",
	"_app/immutable/chunks/ha0gbev_.js",
	"_app/immutable/chunks/CwswN3iS.js",
	"_app/immutable/chunks/smeha_4o.js",
	"_app/immutable/chunks/CITems8I.js",
	"_app/immutable/chunks/iVSWiVfi.js"
];
const stylesheets = ["_app/immutable/assets/app.rXozgjFU.css"];
const fonts = [];

//#endregion
export { component, fonts, imports, index, _layout_server_exports as server, server_id, stylesheets };
//# sourceMappingURL=2-CVrLTvCx.js.map