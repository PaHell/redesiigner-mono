const production = process.env.NODE_ENV !== 'debug';

const config = {
	plugins: [
		require('postcss-import'),
		//Some plugins, like tailwindcss/nesting, need to run before Tailwind,
		require('tailwindcss/nesting')(require('postcss-nesting')),
		require('tailwindcss'),
		//But others, like autoprefixer, need to run after,
		require('autoprefixer'),
		production &&
			require('cssnano')({
				preset: 'default'
			})
	]
};

module.exports = config;
