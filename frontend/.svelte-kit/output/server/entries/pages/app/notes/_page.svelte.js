import { a0 as escape_html, a2 as attr, a4 as ensure_array_like } from "../../../../chunks/index2.js";
import "../../../../chunks/cookie.js";
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let notes = [];
    let title = "";
    let content = "";
    $$renderer2.push(`<div class="space-y-8"><div class="flex items-center justify-between"><h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Notes</h1></div> <div class="glass-card p-6"><h2 class="text-xl font-semibold mb-4 text-primary">${escape_html("New Note")}</h2> <div class="space-y-4"><input${attr("value", title)} placeholder="Title" class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition"/> <textarea placeholder="Content" rows="4" class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition resize-none">`);
    const $$body = escape_html(content);
    if ($$body) {
      $$renderer2.push(`${$$body}`);
    }
    $$renderer2.push(`</textarea> <div class="flex gap-3"><button class="glass-button px-6 py-2 rounded-lg hover:shadow-lg hover:shadow-primary/20 transition transform hover:-translate-y-0.5">${escape_html("Add")} Note</button> `);
    {
      $$renderer2.push("<!--[!-->");
    }
    $$renderer2.push(`<!--]--></div></div></div> <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"><!--[-->`);
    const each_array = ensure_array_like(notes);
    for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
      let note = each_array[$$index];
      $$renderer2.push(`<div class="glass-card p-6 flex flex-col relative group hover:border-primary/50 transition duration-300"><h3 class="text-xl font-bold text-white mb-2">${escape_html(note.title)}</h3> <p class="text-gray-300 whitespace-pre-wrap flex-grow leading-relaxed">${escape_html(note.content)}</p> <div class="mt-6 flex justify-end gap-3 pt-4 border-t border-white/5"><button class="text-primary hover:text-green-300 font-medium text-sm">Edit</button> <button class="text-accent hover:text-red-300 font-medium text-sm">Delete</button></div></div>`);
    }
    $$renderer2.push(`<!--]--></div></div>`);
  });
}
export {
  _page as default
};
