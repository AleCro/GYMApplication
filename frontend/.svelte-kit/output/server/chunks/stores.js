import { a8 as ssr_context } from "./index2.js";
import { w as writable } from "./index.js";
function onDestroy(fn) {
  /** @type {SSRContext} */
  ssr_context.r.on_destroy(fn);
}
const userStore = writable(null);
export {
  onDestroy as o,
  userStore as u
};
