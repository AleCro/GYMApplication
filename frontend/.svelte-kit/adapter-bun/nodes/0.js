import * as server from '../entries/pages/_layout.server.js';

export const index = 0;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/layout.svelte.js')).default;
export { server };
export const server_id = "src/routes/+layout.server.js";
export const imports = ["_app/immutable/nodes/0.CPihVZeV.js","_app/immutable/chunks/DsnmJJEf.js","_app/immutable/chunks/DFFLsmBn.js","_app/immutable/chunks/U_jKa9eL.js"];
export const stylesheets = [];
export const fonts = [];
