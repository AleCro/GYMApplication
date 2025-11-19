import "@sveltejs/kit";
const load = async ({ locals }) => {
  return {
    user: locals.user
  };
};
export {
  load
};
