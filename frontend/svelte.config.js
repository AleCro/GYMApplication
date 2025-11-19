// import adapter from "svelte-adapter-bun";
// import adapter from '@sveltejs/adapter-node';
import adapter from '@sveltejs/adapter-vercel';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		// adapter: adapter()
		adapter: adapter({
  			runtime: 'nodejs20.x'
		})
	}
};

export default config;
