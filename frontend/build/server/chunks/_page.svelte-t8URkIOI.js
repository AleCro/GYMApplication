import { attr, ensure_array_like, escape_html } from "./index2-BLSeo-Vt.js";
import "./clsx-cC83_lR5.js";
import "./cookie-toVUY8nI.js";

//#region .svelte-kit/adapter-bun/entries/pages/app/progress/_page.svelte.js
function _page($$renderer, $$props) {
	$$renderer.component(($$renderer2) => {
		let progressEntries = [];
		let title = "";
		let description = "";
		let uploading = false;
		$$renderer2.push(`<div class="space-y-8"><h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Progress Tracker</h1> <div class="glass-card p-8 rounded-2xl"><h2 class="text-2xl font-bold text-white mb-6 flex items-center gap-2"><span class="text-primary">${escape_html("+")}</span> ${escape_html("Add New Entry")}</h2> <div class="space-y-6"><div class="grid grid-cols-1 md:grid-cols-2 gap-6"><div class="space-y-2"><label class="text-xs font-bold text-gray-400 uppercase tracking-wider ml-1">Title</label> <input${attr("value", title)} placeholder="e.g., Week 1 Transformation" class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition"/></div> <div class="space-y-2"><label class="text-xs font-bold text-gray-400 uppercase tracking-wider ml-1">Photo</label> <input type="file" accept="image/*" id="fileInput"${attr("disabled", false, true)} class="w-full glass-input px-4 py-2.5 rounded-lg focus:ring-2 focus:ring-primary transition file:mr-4 file:py-1 file:px-3 file:rounded-md file:border-0 file:text-xs file:font-semibold file:bg-primary file:text-black hover:file:bg-green-400 cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"/> `);
		$$renderer2.push("<!--[!-->");
		$$renderer2.push(`<!--]--></div></div> <div class="space-y-2"><label class="text-xs font-bold text-gray-400 uppercase tracking-wider ml-1">Description</label> <textarea placeholder="How do you feel? Current weight? Personal records?" rows="3" class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition resize-none">`);
		const $$body = escape_html(description);
		if ($$body) $$renderer2.push(`${$$body}`);
		$$renderer2.push(`</textarea></div> <div class="flex justify-end gap-3">`);
		$$renderer2.push("<!--[!-->");
		$$renderer2.push(`<!--]--> <button${attr("disabled", uploading, true)} class="glass-button px-8 py-3 rounded-lg hover:shadow-lg hover:shadow-primary/20 transition transform hover:-translate-y-0.5 disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none font-bold text-black">${escape_html("Save Progress")}</button></div></div></div> <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8"><!--[-->`);
		const each_array = ensure_array_like(progressEntries);
		for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
			let entry = each_array[$$index];
			$$renderer2.push(`<div class="glass-card rounded-2xl overflow-hidden flex flex-col hover:shadow-2xl hover:shadow-primary/10 transition duration-300 group border border-white/5 relative"><div class="aspect-[4/3] bg-black/40 relative overflow-hidden"><img${attr("src", entry.imageData)}${attr("alt", entry.title)} class="w-full h-full object-cover transition duration-700 group-hover:scale-110"/> <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition duration-300 flex items-end p-6"><span class="text-white font-medium bg-primary/20 backdrop-blur-md px-3 py-1 rounded-full border border-primary/30 text-sm">${escape_html(new Date(entry.createdAt).toLocaleDateString(void 0, {
				year: "numeric",
				month: "long",
				day: "numeric"
			}))}</span></div></div> <div class="p-6 flex-grow flex flex-col gap-2"><div class="flex justify-between items-start"><h3 class="text-2xl font-bold text-white group-hover:text-primary transition">${escape_html(entry.title)}</h3></div> <p class="text-gray-400 text-sm leading-relaxed">${escape_html(entry.description)}</p> <div class="mt-4 flex justify-end gap-3 pt-4 border-t border-white/5"><button class="text-sm font-medium text-primary hover:text-green-300 transition">Edit</button> <button class="text-sm font-medium text-accent hover:text-red-300 transition">Delete</button></div></div></div>`);
		}
		$$renderer2.push(`<!--]--></div></div>`);
	});
}

//#endregion
export { _page as default };
//# sourceMappingURL=_page.svelte-t8URkIOI.js.map