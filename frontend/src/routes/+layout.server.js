import { redirect } from '@sveltejs/kit';

export async function load({ locals, url }) {
  console.log(locals.user);
  if (locals.user == undefined && url.pathname != "/login") {
    throw redirect(302, "/login");
  }
  if (url.pathname == "/login" && locals.user != undefined) {
      throw redirect(302, "/");
  }
  return {
    user: locals.user
  };
}
