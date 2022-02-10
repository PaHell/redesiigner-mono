<script lang="ts" context="module">
	const alias = {
		loading: 'loader-4',
		home: 'home-gear',
		menu: 'menu-2',
		select: 'arrow-down-s',
		language: 'global',
		notification: 'notification',
		device: 'computer',
		add: 'add',
		check: 'check',
		cross: 'close',
		link: 'global',
		privacy: 'account-box',
		delete: 'delete-bin',
		warning: 'alarm-warning'
	};
</script>

<script lang="ts">
	export let name: string = 'home';
	export let css: string = '';
	export let large: boolean = false;
	export let circle: boolean = false;
</script>

<template>
	<div
		class="app-icon {css}"
		class:app-icon-large={large}
		class:app-icon-circle={circle}
	>
		<i class="ri-{alias[name]}-line" />
	</div>
</template>

<style global lang="postcss">
	.app-icon {
		@apply flex justify-center items-center
    		text-ic-base;

		&.app-icon-large {
			@apply text-ic-lg sm:text-ic-base;
		}

		&.app-icon-circle {
			@apply flex-shrink-0 w-16 h-16 sm:w-12 sm:h-12 rounded-full;
		}
	}

	html.dark .app-icon {
		&.app-icon-circle {
			@apply border-2 border-gray-700 bg-transparent;
		}
	}
</style>
