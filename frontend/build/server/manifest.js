//#region .svelte-kit/adapter-bun/manifest.js
const manifest = (() => {
	function __memo(fn) {
		let value;
		return () => value ??= value = fn();
	}
	return {
		appDir: "_app",
		appPath: "_app",
		assets: new Set(["robots.txt"]),
		mimeTypes: { ".txt": "text/plain" },
		_: {
			client: {
				start: "_app/immutable/entry/start.BYwTlBXL.js",
				app: "_app/immutable/entry/app.D9wqgGwD.js",
				imports: [
					"_app/immutable/entry/start.BYwTlBXL.js",
					"_app/immutable/chunks/P8b-s9jE.js",
					"_app/immutable/chunks/1-JAo-Vw.js",
					"_app/immutable/chunks/DFFLsmBn.js",
					"_app/immutable/chunks/CVaLoUw_.js",
					"_app/immutable/chunks/U_jKa9eL.js",
					"_app/immutable/chunks/DvDWJgkm.js",
					"_app/immutable/chunks/CdP2dVa8.js",
					"_app/immutable/chunks/smeha_4o.js",
					"_app/immutable/entry/app.D9wqgGwD.js",
					"_app/immutable/chunks/DFFLsmBn.js",
					"_app/immutable/chunks/CVaLoUw_.js",
					"_app/immutable/chunks/DsnmJJEf.js",
					"_app/immutable/chunks/1-JAo-Vw.js",
					"_app/immutable/chunks/U_jKa9eL.js",
					"_app/immutable/chunks/DIYmMOxU.js",
					"_app/immutable/chunks/NJrgjD4c.js",
					"_app/immutable/chunks/BOruTPHP.js",
					"_app/immutable/chunks/CdP2dVa8.js"
				],
				stylesheets: [],
				fonts: [],
				uses_env_dynamic_public: true
			},
			nodes: [
				__memo(() => import("./chunks/0-DO7Isd8t.js")),
				__memo(() => import("./chunks/1-_t5sTLIb.js")),
				__memo(() => import("./chunks/2-CVrLTvCx.js")),
				__memo(() => import("./chunks/3-BaEn5F0m.js")),
				__memo(() => import("./chunks/4-CB62pYM_.js")),
				__memo(() => import("./chunks/5-83KAhEkj.js")),
				__memo(() => import("./chunks/6-DpJPQhUK.js")),
				__memo(() => import("./chunks/7-7OYvtrxP.js")),
				__memo(() => import("./chunks/8-CmtPhWYh.js")),
				__memo(() => import("./chunks/9-BfurDKHj.js")),
				__memo(() => import("./chunks/10-CbsUhYAp.js")),
				__memo(() => import("./chunks/11-B0mxSb3q.js")),
				__memo(() => import("./chunks/12-vIHjVtt1.js")),
				__memo(() => import("./chunks/13-BRlC-5AN.js")),
				__memo(() => import("./chunks/14-BUCb1yt7.js")),
				__memo(() => import("./chunks/15-BZsFK7P-.js"))
			],
			remotes: {},
			routes: [
				{
					id: "/",
					pattern: /^\/$/,
					params: [],
					page: {
						layouts: [0],
						errors: [1],
						leaf: 4
					},
					endpoint: null
				},
				{
					id: "/app",
					pattern: /^\/app\/?$/,
					params: [],
					page: {
						layouts: [0, 2],
						errors: [1, ,],
						leaf: 6
					},
					endpoint: __memo(() => import("./chunks/_server-DAkJzVGy.js"))
				},
				{
					id: "/app/calendar",
					pattern: /^\/app\/calendar\/?$/,
					params: [],
					page: {
						layouts: [0, 2],
						errors: [1, ,],
						leaf: 9
					},
					endpoint: null
				},
				{
					id: "/app/documentation",
					pattern: /^\/app\/documentation\/?$/,
					params: [],
					page: {
						layouts: [0, 2],
						errors: [1, ,],
						leaf: 7
					},
					endpoint: null
				},
				{
					id: "/app/me",
					pattern: /^\/app\/me\/?$/,
					params: [],
					page: {
						layouts: [0, 2],
						errors: [1, ,],
						leaf: 8
					},
					endpoint: null
				},
				{
					id: "/app/notes",
					pattern: /^\/app\/notes\/?$/,
					params: [],
					page: {
						layouts: [0, 2],
						errors: [1, ,],
						leaf: 11
					},
					endpoint: null
				},
				{
					id: "/app/progress",
					pattern: /^\/app\/progress\/?$/,
					params: [],
					page: {
						layouts: [0, 2],
						errors: [1, ,],
						leaf: 12
					},
					endpoint: null
				},
				{
					id: "/app/user-management",
					pattern: /^\/app\/user-management\/?$/,
					params: [],
					page: {
						layouts: [0, 2],
						errors: [1, ,],
						leaf: 10
					},
					endpoint: null
				},
				{
					id: "/app/workouts",
					pattern: /^\/app\/workouts\/?$/,
					params: [],
					page: {
						layouts: [0, 2],
						errors: [1, ,],
						leaf: 13
					},
					endpoint: null
				},
				{
					id: "/login",
					pattern: /^\/login\/?$/,
					params: [],
					page: {
						layouts: [0, 3],
						errors: [1, ,],
						leaf: 14
					},
					endpoint: null
				},
				{
					id: "/logout",
					pattern: /^\/logout\/?$/,
					params: [],
					page: {
						layouts: [0],
						errors: [1],
						leaf: 15
					},
					endpoint: null
				},
				{
					id: "/privacy",
					pattern: /^\/privacy\/?$/,
					params: [],
					page: {
						layouts: [0],
						errors: [1],
						leaf: 5
					},
					endpoint: null
				}
			],
			prerendered_routes: new Set([]),
			matchers: async () => {
				return {};
			},
			server_assets: {}
		}
	};
})();
const prerendered = new Set([]);
const base = "";

//#endregion
export { base, manifest, prerendered };
//# sourceMappingURL=manifest.js.map