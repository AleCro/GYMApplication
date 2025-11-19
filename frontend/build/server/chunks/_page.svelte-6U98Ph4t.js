import { attr, bind_props, ensure_array_like, escape_html, head } from "./index2-BLSeo-Vt.js";
import "./clsx-cC83_lR5.js";
import "./chunks-DNg0X1yh.js";
import { onDestroy, userStore } from "./stores-C-8AFukU.js";
import "./cookie-toVUY8nI.js";
import { isPasswordSecure } from "./passwordCheck-cXi-893K.js";

//#region .svelte-kit/adapter-bun/entries/pages/app/me/_page.svelte.js
function _page($$renderer, $$props) {
	$$renderer.component(($$renderer2) => {
		let user, passwordCheck, passwordsMatch, isFormValid;
		let data = $$props["data"];
		userStore.set(data?.user ?? null);
		let unsubscribe = userStore.subscribe((value) => {
			user = value;
		});
		const userInitial = (user.username ?? "G").charAt(0).toUpperCase();
		let oldPassword = "";
		let newPassword = "";
		let confirmNewPassword = "";
		let loading = false;
		onDestroy(() => {
			unsubscribe();
		});
		user = data.user;
		passwordCheck = isPasswordSecure(newPassword);
		passwordsMatch = newPassword.length > 0 && newPassword === confirmNewPassword;
		isFormValid = passwordCheck.secure && passwordsMatch && oldPassword.length > 0 && !loading;
		head($$renderer2, ($$renderer3) => {
			$$renderer3.title(($$renderer4) => {
				$$renderer4.push(`<title>My Profile</title>`);
			});
		});
		$$renderer2.push(`<div class="space-y-8 max-w-2xl mx-auto"><section><h1 class="text-2xl font-semibold text-white mb-6">My Profile</h1> <div class="flex items-center gap-6 p-6 bg-gray-800 rounded-lg"><div class="w-24 h-24 rounded-full bg-indigo-600 text-white flex items-center justify-center text-5xl font-medium shrink-0">${escape_html(userInitial)}</div> <div><div class="text-xs text-gray-400 uppercase tracking-wider">Username</div> <div class="text-3xl font-bold text-gray-100">${escape_html(user.username)}</div> `);
		if (user.email) {
			$$renderer2.push("<!--[-->");
			$$renderer2.push(`<div class="text-sm text-gray-300 mt-1">${escape_html(user.email)}</div>`);
		} else $$renderer2.push("<!--[!-->");
		$$renderer2.push(`<!--]--></div></div></section> <section><h2 class="text-xl font-semibold text-white mb-5">Change Password</h2> <form class="p-6 bg-gray-800 rounded-lg space-y-6">`);
		$$renderer2.push("<!--[!-->");
		$$renderer2.push(`<!--]--> `);
		$$renderer2.push("<!--[!-->");
		$$renderer2.push(`<!--]--> <div><label for="old-password" class="block text-sm font-medium text-gray-300 mb-2">Old Password</label> <input id="old-password" type="password"${attr("value", oldPassword)} required class="w-full px-3 py-2 bg-gray-900 border border-gray-700 rounded-md placeholder-gray-500 text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"/></div> <div><label for="new-password" class="block text-sm font-medium text-gray-300 mb-2">New Password</label> <input id="new-password" type="password"${attr("value", newPassword)} required${attr("aria-invalid", newPassword.length > 0 && !passwordCheck.secure)} class="w-full px-3 py-2 bg-gray-900 border border-gray-700 rounded-md placeholder-gray-500 text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"/></div> `);
		if (newPassword.length > 0 && !passwordCheck.secure) {
			$$renderer2.push("<!--[-->");
			$$renderer2.push(`<div class="p-4 bg-gray-900 border border-gray-700 rounded-md"><ul class="space-y-1 text-sm text-gray-400"><!--[-->`);
			const each_array = ensure_array_like(passwordCheck.issues);
			for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
				let issue = each_array[$$index];
				$$renderer2.push(`<li class="flex items-center"><svg class="w-4 h-4 mr-2 text-red-400 shrink-0" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor"><path d="M6 18L18 6M6 6l12 12"></path></svg> <span>${escape_html(issue)}</span></li>`);
			}
			$$renderer2.push(`<!--]--></ul></div>`);
		} else $$renderer2.push("<!--[!-->");
		$$renderer2.push(`<!--]--> <div><label for="confirm-password" class="block text-sm font-medium text-gray-300 mb-2">Confirm New Password</label> <input id="confirm-password" type="password"${attr("value", confirmNewPassword)} required${attr("aria-invalid", confirmNewPassword.length > 0 && !passwordsMatch)} class="w-full px-3 py-2 bg-gray-900 border border-gray-700 rounded-md placeholder-gray-500 text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"/> `);
		if (confirmNewPassword.length > 0 && !passwordsMatch) {
			$$renderer2.push("<!--[-->");
			$$renderer2.push(`<p class="mt-2 text-sm text-red-400">Passwords do not match.</p>`);
		} else $$renderer2.push("<!--[!-->");
		$$renderer2.push(`<!--]--></div> <div class="pt-2"><button type="submit"${attr("disabled", !isFormValid || loading, true)} class="w-full sm:w-auto px-6 py-2.5 bg-indigo-600 text-white font-semibold rounded-md shadow-md hover:bg-indigo-500 transition duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed">`);
		{
			$$renderer2.push("<!--[!-->");
			$$renderer2.push(`Change Password`);
		}
		$$renderer2.push(`<!--]--></button></div></form></section></div>`);
		bind_props($$props, { data });
	});
}

//#endregion
export { _page as default };
//# sourceMappingURL=_page.svelte-6U98Ph4t.js.map