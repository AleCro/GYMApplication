import { $ as head, a0 as escape_html, a2 as attr, a4 as ensure_array_like, a7 as attr_class, a3 as stringify, a6 as bind_props } from "../../../../chunks/index2.js";
import { p as public_env } from "../../../../chunks/shared-server.js";
import { c as cookies } from "../../../../chunks/cookie.js";
import { u as userStore, o as onDestroy } from "../../../../chunks/stores.js";
let request = (path = "/", method = "GET", body = null, headers = null) => {
  return new Promise((resolve, reject) => {
    let req = { method };
    if (headers) {
      req.headers = headers;
    }
    if (body) {
      req.body = body;
    }
    fetch(`${public_env.PUBLIC_API_URL}${path}`, req).then((res) => res.json()).then(resolve).catch(reject);
  });
};
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let data = $$props["data"];
    userStore.set(data?.user ?? null);
    let unsubscribe = userStore.subscribe((value) => {
    });
    onDestroy(() => {
      unsubscribe();
    });
    let users = data.users || [];
    let total = data.total || 0;
    let limit = data.limit || 10;
    let currentPage = data.page || 1;
    let isLoading = false;
    let error = null;
    let searchTerm = "";
    let isSearching = false;
    const ROLE_MAP = { 0: "User", 255: "Administrator" };
    const getRoleName = (role) => {
      return ROLE_MAP[role] || `Custom (${role})`;
    };
    async function fetchUsers(page, search = "") {
      isLoading = true;
      error = null;
      currentPage = page;
      try {
        let path = `/users?limit=${limit}&page=${page}`;
        if (search.trim()) {
          path += `&search=${encodeURIComponent(search.trim())}`;
        }
        const response = await request(path, "GET", null, { Authorization: `Bearer ${cookies.token}` });
        users = response.users || [];
        total = response.total || 0;
        limit = response.limit || 10;
        currentPage = response.page || 1;
      } catch (err) {
        error = "Failed to load users: " + (err.message || "Network error");
        console.error(err);
      } finally {
        isLoading = false;
        isSearching = false;
      }
    }
    data.user;
    {
      if (users.length === 0 && !isLoading) {
        fetchUsers(currentPage);
      }
    }
    head($$renderer2, ($$renderer3) => {
      $$renderer3.title(($$renderer4) => {
        $$renderer4.push(`<title>User Management</title>`);
      });
    });
    $$renderer2.push(`<div class="space-y-6"><header class="pb-4 border-b border-gray-700"><h1 class="text-3xl font-bold text-white">👤 User Management</h1> <p class="mt-1 text-gray-400">Manage user accounts, roles, and details. Total Users: ${escape_html(total)}</p></header> <div class="bg-gray-800 p-4 rounded-lg border border-gray-700"><div class="flex gap-4 items-center"><div class="flex-1 relative"><div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"><svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg></div> <input type="text"${attr("value", searchTerm)} placeholder="Search by username or ID..." class="w-full pl-10 pr-4 py-2 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"/> `);
    {
      $$renderer2.push("<!--[!-->");
    }
    $$renderer2.push(`<!--]--></div> `);
    {
      $$renderer2.push("<!--[!-->");
    }
    $$renderer2.push(`<!--]--></div> `);
    {
      $$renderer2.push("<!--[!-->");
    }
    $$renderer2.push(`<!--]--></div> `);
    if (error) {
      $$renderer2.push("<!--[-->");
      $$renderer2.push(`<div class="p-4 bg-red-900/30 text-red-300 rounded-lg border border-red-700 font-medium" role="alert">${escape_html(error)}</div>`);
    } else {
      $$renderer2.push("<!--[!-->");
    }
    $$renderer2.push(`<!--]--> <div class="overflow-x-auto rounded-lg border border-gray-700 shadow-xl bg-gray-800"><table class="min-w-full divide-y divide-gray-700"><thead class="bg-gray-700/50"><tr><th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">ID</th><th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">Username</th><th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">Role</th><th scope="col" class="relative px-6 py-3"><span class="sr-only">Edit</span></th></tr></thead><tbody class="divide-y divide-gray-800">`);
    if (isLoading) {
      $$renderer2.push("<!--[-->");
      $$renderer2.push(`<tr><td colspan="5" class="px-6 py-4 text-center text-gray-500 italic">`);
      if (isSearching) {
        $$renderer2.push("<!--[-->");
        $$renderer2.push(`Searching users...`);
      } else {
        $$renderer2.push("<!--[!-->");
        $$renderer2.push(`Loading users...`);
      }
      $$renderer2.push(`<!--]--></td></tr>`);
    } else {
      $$renderer2.push("<!--[!-->");
      if (users.length === 0) {
        $$renderer2.push("<!--[-->");
        $$renderer2.push(`<tr><td colspan="5" class="px-6 py-4 text-center text-gray-500 italic">`);
        {
          $$renderer2.push("<!--[!-->");
          $$renderer2.push(`No users found.`);
        }
        $$renderer2.push(`<!--]--></td></tr>`);
      } else {
        $$renderer2.push("<!--[!-->");
        $$renderer2.push(`<!--[-->`);
        const each_array = ensure_array_like(users);
        for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
          let user = each_array[$$index];
          $$renderer2.push(`<tr${attr_class(`transition duration-150 ${stringify(user.id == data?.user.userID ? "bg-gray-700 hover:bg-gray-900" : "hover:bg-gray-700")}`)}><td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-300">${escape_html(user?.id)}</td><td class="px-6 py-4 whitespace-nowrap text-sm text-white font-medium">${escape_html(user.username)}</td><td class="px-6 py-4 whitespace-nowrap text-sm text-gray-400"><span${attr_class("px-2 inline-flex text-xs leading-5 font-semibold rounded-full", void 0, {
            "bg-indigo-600": user.group >= 100,
            "bg-blue-600": user.group < 100
          })}>${escape_html(getRoleName(user.group))}</span></td><td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium"><button class="text-indigo-400 hover:text-indigo-300">Edit</button></td></tr>`);
        }
        $$renderer2.push(`<!--]-->`);
      }
      $$renderer2.push(`<!--]-->`);
    }
    $$renderer2.push(`<!--]--></tbody></table></div> <div class="flex justify-between items-center pt-4"><p class="text-sm text-gray-400">Showing ${escape_html(Math.min(limit * (currentPage - 1) + 1, total))} to ${escape_html(Math.min(limit * currentPage, total))} of ${escape_html(total)} results. `);
    {
      $$renderer2.push("<!--[!-->");
    }
    $$renderer2.push(`<!--]--></p> <div class="flex gap-2"><button${attr("disabled", currentPage === 1 || isLoading, true)}${attr_class("px-4 py-2 text-sm font-medium rounded-md transition disabled:opacity-50 disabled:cursor-not-allowed", void 0, {
      "bg-gray-700": currentPage !== 1,
      "hover:bg-gray-600": currentPage !== 1,
      "text-white": currentPage !== 1
    })}>Previous</button> <button${attr("disabled", currentPage * limit >= total || isLoading, true)}${attr_class("px-4 py-2 text-sm font-medium rounded-md transition disabled:opacity-50 disabled:cursor-not-allowed", void 0, {
      "bg-gray-700": currentPage * limit < total,
      "hover:bg-gray-600": currentPage * limit < total,
      "text-white": currentPage * limit < total
    })}>Next</button></div></div></div> `);
    {
      $$renderer2.push("<!--[!-->");
    }
    $$renderer2.push(`<!--]-->`);
    bind_props($$props, { data });
  });
}
export {
  _page as default
};
