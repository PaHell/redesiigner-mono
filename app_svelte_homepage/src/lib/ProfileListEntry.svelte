<script lang="ts" context="module">
	import Icon from '$lib/Icon.svelte';
	import Btn from '$lib/Button.svelte';

	const colors = [
		'bg-red-100 dark:bg-red-500 text-red-500',
		'bg-orange-100 dark:bg-orange-500 text-orange-500',
		'bg-amber-100 dark:bg-amber-500 text-amber-500',
		'bg-lime-100 dark:bg-lime-500 text-lime-500',
		'bg-green-100 dark:bg-green-500 text-green-500',
		'bg-emerald-100 dark:bg-emerald-500 text-emerald-500',
		'bg-teal-100 dark:bg-teal-500 text-teal-500',
		'bg-sky-100 dark:bg-sky-500 text-sky-500',
		'bg-blue-100 dark:bg-blue-500 text-blue-500',
		'bg-indigo-100 dark:bg-indigo-500 text-indigo-500',
		'bg-purple-100 dark:bg-purple-500 text-purple-500'
	];
</script>

<script lang="ts">
	export let img = '';
	export let name = 'Name';
	export let sub = 'Sub';
	export let status = 'offline';

	function getInitials() {
		const split = name.split(' ');
		const first = split[0];
		const last = split[split.length - 1];
		if (split.length === 1) return first[0].toUpperCase() + first[1];
		return first[0].toUpperCase() + last[0].toUpperCase();
	}

	function getColorFromName() {
		return colors[name.charCodeAt(0) % colors.length];
	}
</script>

<template>
	<div class="app-list-entry-profile">
		<div class="app-img">
			{#if img}
				<img src={img} />
				<div class="app-img-shadow" />
			{:else}
				<div class="app-img-placeholder {getColorFromName()} text-white dark:text-gray-900">
					{getInitials()}
				</div>
			{/if}
			{#if status}
				<div class="app-status app-status-{status}" />
			{/if}
		</div>
		<div class="flex-1">
			<h2 class="font-semibold dark:text-white text-gray-800">
				{name}
			</h2>
			<p class="text-sm mt-[-2px] mb-[2px] dark:text-gray-400 text-gray-500">
				{sub}
			</p>
		</div>
		<slot />
	</div>
</template>

<style global lang="postcss">
	.app-list-entry-profile {
		@apply flex items-center justify-center
        w-full p-4;

		& > .app-img {
			@apply flex flex-wrap justify-center items-center
            w-11 h-11 mr-4
            rounded-full overflow-hidden;

			& > .app-status {
				content: '';
				@apply block absolute;
			}

			& > .app-img-placeholder {
				@apply w-full h-full
                leading-[2.8rem] text-center
                font-normal font-mono text-lg tracking-wide
                rounded-full;
			}

			& > img {
				@apply min-w-full min-h-full;
			}

			& > .app-img-shadow {
				@apply absolute
                w-11 h-11 rounded-full;
				box-shadow: inset 0 0 4px rgba(0, 0, 0, 0.3);
			}
		}
	}

	html.dark .app-list-entry-profile {
		& .app-img-shadow {
			@apply hidden;
		}
		& .app-img-placeholder {
			@apply font-semibold text-white;
		}
	}

	.app-status {
		@apply w-3 h-3 mt-8 ml-8
        rounded-full ring-[2px] ring-gray-100;
		&.app-status-offline {
			@apply bg-gray-400;
		}
		&.app-status-online {
			@apply bg-green-500;
		}
		&.app-status-away {
			@apply bg-orange-400;
		}
		&.app-status-busy {
			@apply bg-red-500;
		}
	}

	html.dark .app-status {
		@apply ring-gray-800;
	}
</style>
