import adapter from '@sveltejs/adapter-auto';
import sveltePreprocess from 'svelte-preprocess';
import tailwindCss from 'tailwindcss'; 
import postCssNesting from 'postcss-nesting';
import autoprefixer from 'autoprefixer';

const production = process.env.NODE_ENV !== 'debug';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: sveltePreprocess({
		sourceMap: !production,
		postcss: {
			plugins: [
				postCssNesting,
				tailwindCss,
				autoprefixer
			],
		},
	}),

	kit: {
		adapter: adapter(),

		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte'
	}
};

export default config;
